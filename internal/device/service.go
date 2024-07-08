package device

import (
	"context"
	"encoding/base64"
	"log/slog"
	"net/http"
	"time"

	"github.com/hrz8/goatsapp/internal/dbrepo"
	"github.com/hrz8/goatsapp/pkg/whatsapp"
	"github.com/hrz8/gofx"
	"github.com/skip2/go-qrcode"
	"go.mau.fi/whatsmeow"
)

type Service struct {
	log   *gofx.Logger
	repo  *dbrepo.Queries
	waCli *whatsapp.Client
}

func NewService(log *gofx.Logger, repo *dbrepo.Queries, waCli *whatsapp.Client) *Service {
	return &Service{log, repo, waCli}
}

func (s *Service) DecodeProjectID(c *gofx.Context, encodedID string) int32 {
	project, _ := s.repo.GetProjectByEncodedID(c.Request().Context(), encodedID)
	return project.ID
}

func (s *Service) CheckIfNotScanned(sessionID string) {
	cli := s.waCli.Get(sessionID)
	qr := s.waCli.GetQR(sessionID)

	if cli != nil && qr != "" && !cli.IsLoggedIn() {
		// initiated but qr not scanned yet
		_ = s.waCli.Reset(sessionID)
		_ = s.waCli.ResetQR(sessionID)
		_ = cli.Logout()
	}
}

func (s *Service) CreateDevice(c *gofx.Context, payload *CreateDeviceDto) error {
	evt := "CreateDevice"
	var err error
	var cookie *http.Cookie

	cookie, err = c.Cookie("project_id")
	if err != nil {
		s.log.JSON.Error(
			"failed to get project_id cookie",
			slog.String("event", evt), slog.String("err", err.Error()),
		)
		return ErrProjectIDNotFound
	}

	project, _ := s.repo.GetProjectByEncodedID(c.Request().Context(), cookie.Value)

	data := new(dbrepo.CreateNewDeviceParams)
	data.ProjectID = project.ID
	data.Name = payload.Name
	data.Description = &payload.Description
	data.ClientDeviceID = payload.ClientDeviceID
	data.CustomAttributes = []byte(payload.CustomAttributes)
	if _, err = s.repo.CreateNewDevice(c.Request().Context(), data); err != nil {
		s.log.JSON.Error(
			"failed when store to db",
			slog.String("event", evt), slog.String("err", err.Error()),
		)
		return ErrRepository
	}

	return nil
}

func (s *Service) GetDevicesByProjectID(
	c *gofx.Context, encodedID string,
) ([]*dbrepo.GetDevicesByProjectEncodedIDRow, error) {
	evt := "GetDevicesByProjectID"
	var err error
	var devices []*dbrepo.GetDevicesByProjectEncodedIDRow

	projectID := s.DecodeProjectID(c, encodedID)
	devices, err = s.repo.GetDevicesByProjectEncodedID(c.Request().Context(), projectID)
	if err != nil {
		s.log.JSON.Error(
			"failed when list to devices data",
			slog.String("event", evt), slog.String("err", err.Error()),
		)
		return devices, ErrRepository
	}

	return devices, nil
}

func (s *Service) RequestQRCode(_ *gofx.Context, sessionID string) (string, error) {
	evt := "RequestQRCode"
	s.CheckIfNotScanned(sessionID)

	var err error
	var cli *whatsmeow.Client

	cli = s.waCli.Get(sessionID)
	qr := s.waCli.GetQR(sessionID)

	if cli == nil {
		cli = s.waCli.NewMeow(sessionID)
		if err = s.waCli.Set(sessionID, cli); err != nil {
			s.log.JSON.Error(
				"failed when set new whatsapp client",
				slog.String("event", evt), slog.String("err", err.Error()),
			)
			return "", ErrWhatsappClient
		}
	}

	if qr != "" {
		return qr, nil
	}

	if cli.IsLoggedIn() {
		s.log.JSON.Error("client already logged in", slog.String("event", evt))
		return "", ErrAlreadyLoggedIn
	}

	ctx, cancel := context.WithCancel(context.Background())
	qrChan, errCh := cli.GetQRChannel(ctx)
	if errCh != nil {
		defer cancel()
		s.log.JSON.Error(
			"failed when perform whatsmeow.GetQRChannel()",
			slog.String("event", evt), slog.String("err", errCh.Error()),
		)
		return "", ErrWhatsappLibrary
	}

	chImg := make(chan string)
	go func() {
		evt := <-qrChan
		if evt.Event == "code" {
			go func() {
				time.Sleep(evt.Timeout - 10*time.Second)
				if !cli.IsLoggedIn() {
					s.log.Info("expiring qr code", slog.String("session_id", sessionID))
					_ = s.waCli.Reset(sessionID)
					_ = s.waCli.ResetQR(sessionID)
					cancel()
				}
			}()
			img, errQR := qrcode.Encode(evt.Code, qrcode.Medium, 256)
			if errQR != nil {
				cancel()
			}
			base64Img := base64.StdEncoding.EncodeToString(img)
			chImg <- "data:image/png;base64," + base64Img
		}
	}()

	errConn := cli.Connect()
	if errConn != nil {
		defer cancel()
		s.log.JSON.Error(
			"failed when perform whatsmeow.Connect()",
			slog.String("event", evt), slog.String("err", errConn.Error()),
		)
		return "", ErrWhatsappLibrary
	}

	qr = <-chImg
	_ = s.waCli.SetQR(sessionID, qr)

	return qr, nil
}
