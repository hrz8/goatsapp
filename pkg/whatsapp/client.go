package whatsapp

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

const (
	LogLevel       = "INFO"
	LogLevelDB     = "INFO"
	LogLevelDevice = "INFO"
)

type EventHandler func(cli *whatsmeow.Client, clientDeviceID string) whatsmeow.EventHandler
type Option func(c *Client)

var (
	ErrClientAlreadyExist = errors.New("whatsapp client with specific id already exist")
	ErrClientNotExist     = errors.New("whatsapp client is not exits")
	ErrQRAlreadyExist     = errors.New("qrcode with specific id already exist")
	ErrQRNotExist         = errors.New("qrcode is not exits")
)

type Client struct {
	mut sync.RWMutex
	WA  map[string]*whatsmeow.Client
	QR  map[string]string

	// customable
	osName     string
	osVersion  [3]uint32
	evtHandler EventHandler

	// default
	container *sqlstore.Container
	mig       *Migration
	repo      *DeviceRepo
	log       waLog.Logger
}

func WithDB(db *sql.DB) Option {
	return func(c *Client) {
		log := waLog.Stdout("Database", LogLevelDB, true)
		c.container = sqlstore.NewWithDB(db, "postgres", log)
		c.mig = &Migration{db, log}
		c.repo = &DeviceRepo{db}
	}
}

func WithOsInfo(name string, version [3]uint32) Option {
	return func(c *Client) {
		c.osName = fmt.Sprintf("%s@%s", name, fmt.Sprintf("v%d.%d.%d", version[0], version[1], version[2]))
		c.osVersion = version
	}
}

func WithEventHandler(handler EventHandler) Option {
	return func(c *Client) {
		c.evtHandler = handler
	}
}

func NewClient(opts ...Option) *Client {
	wa := make(map[string]*whatsmeow.Client)
	qr := make(map[string]string)

	log := waLog.Stdout("Client", LogLevel, true)

	waCli := &Client{
		// core
		WA: wa,
		QR: qr,

		// customable
		osName:     "Whatsapp",
		osVersion:  [3]uint32{0, 1, 0},
		evtHandler: nil,
		container:  nil,
		mig:        nil,
		repo:       nil,

		// default
		log: log,
	}

	for _, opt := range opts {
		opt(waCli)
	}

	store.SetOSInfo(waCli.osName, waCli.osVersion)
	return waCli
}

func (c *Client) NewMeow(clientDeviceID string) *whatsmeow.Client {
	meowDevice := c.container.NewDevice()
	cli := c.initMeow(meowDevice, clientDeviceID)
	return cli
}

func (c *Client) initMeow(device *store.Device, clientDeviceID string) *whatsmeow.Client {
	cliLog := waLog.Stdout("Device-"+clientDeviceID, LogLevelDevice, true)
	cli := whatsmeow.NewClient(device, cliLog)
	cli.AddEventHandler(c.defaultEventHandler(clientDeviceID))
	if c.evtHandler != nil {
		cli.AddEventHandler(c.evtHandler(cli, clientDeviceID))
	}
	return cli
}

func (c *Client) Upgrade() error {
	var err error
	err = c.container.Upgrade()
	if err != nil {
		panic(err)
	}
	err = c.mig.Upgrade()
	if err != nil {
		panic(err)
	}

	return err
}

func (c *Client) GetQR(clientDeviceID string) string {
	c.mut.RLock()
	qr := c.QR[clientDeviceID]
	c.mut.RUnlock()
	if qr == "" {
		c.log.Warnf("cannot find qrcode for id: %v", clientDeviceID)
		return ""
	}
	return qr
}

func (c *Client) SetQR(clientDeviceID string, qr string) error {
	curr := c.GetQR(clientDeviceID)
	if curr != "" {
		c.log.Errorf("cannot reassign qrcode for id: %v", clientDeviceID)
		return ErrQRAlreadyExist
	}
	c.mut.Lock()
	c.QR[clientDeviceID] = qr
	c.mut.Unlock()
	return nil
}

func (c *Client) ResetQR(clientDeviceID string) error {
	curr := c.GetQR(clientDeviceID)
	if curr == "" {
		return ErrQRNotExist
	}
	c.mut.Lock()
	c.QR[clientDeviceID] = ""
	c.mut.Unlock()
	return nil
}

func (c *Client) Get(clientDeviceID string) *whatsmeow.Client {
	c.mut.RLock()
	cli := c.WA[clientDeviceID]
	c.mut.RUnlock()
	if cli == nil {
		c.log.Warnf("cannot find whatsapp client with id: %v", clientDeviceID)
		return nil
	}
	return cli
}

func (c *Client) Set(clientDeviceID string, cli *whatsmeow.Client) error {
	curr := c.Get(clientDeviceID)
	if curr != nil {
		c.log.Errorf("cannot reassigned whatsapp client with id: %v", clientDeviceID)
		return ErrClientAlreadyExist
	}
	c.mut.Lock()
	c.WA[clientDeviceID] = cli
	c.mut.Unlock()
	return nil
}

func (c *Client) Reset(clientDeviceID string) error {
	curr := c.Get(clientDeviceID)
	if curr == nil {
		return ErrClientNotExist
	}
	c.mut.Lock()
	c.WA[clientDeviceID] = nil
	c.mut.Unlock()
	return nil
}

func (c *Client) Backup() {
	c.log.Infof("attempting back up whatsapp clients...")
	var wg sync.WaitGroup
	for clientDeviceID, cli := range c.WA {
		wg.Add(1)
		go func(clientDeviceID string, cli *whatsmeow.Client) {
			defer wg.Done()
			if cli == nil {
				return
			}
			if cli.Store.ID != nil {
				c.log.Debugf("backing up whatsapp client for id: %v", clientDeviceID)
				_, _ = c.repo.SetJID(clientDeviceID, cli.Store.ID.String())
			}
		}(clientDeviceID, cli)
	}
	wg.Wait()
	c.log.Infof("backup done!")
}

func (c *Client) defaultEventHandler(clientDeviceID string) whatsmeow.EventHandler {
	return func(evt interface{}) {
		switch evt.(type) { //nolint
		case *events.PairSuccess:
			err := c.ResetQR(clientDeviceID)
			if err != nil {
				c.log.Errorf("failed when delete qr code after scanned, err: %v", err)
			}
		}
	}
}

func (c *Client) Restore() {
	c.log.Infof("attempting to restoring whatsapp clients connections...")
	meowDevices, err := c.container.GetAllDevices()
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	for _, meowDevice := range meowDevices {
		wg.Add(1)
		go func(meowDevice *store.Device) {
			var extendedDevice *Device

			defer wg.Done()

			jid := meowDevice.ID.String()
			c.log.Debugf("restoring backup for client jid: %v", jid)
			extendedDevice, err = c.repo.GetDeviceByJID(jid)
			if err != nil {
				c.log.Warnf("error getting device from db, jid: %v", jid)
				return
			}
			cli := c.initMeow(meowDevice, extendedDevice.ClientDeviceID)
			err = cli.Connect()
			if err != nil {
				c.log.Errorf("failed when connecting client to web socket, jid: %v", jid)
			}
			_ = c.Set(extendedDevice.ClientDeviceID, cli)
			c.log.Debugf("reset backup data for client id: %v", jid)
			_ = c.repo.Reset(extendedDevice.ClientDeviceID)
		}(meowDevice)
	}
	wg.Wait()
	c.log.Infof("restore done!")
}
