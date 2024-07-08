package whatsapp

import (
	"errors"
	"strings"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
)

var (
	ErrServerUnexpected  = errors.New("server error occurred")
	ErrAlreadyConnected  = errors.New("device already connected")
	ErrNotLogin          = errors.New("device not login yet")
	ErrRecipientNotFound = errors.New("recipient number not found")
)

func IsGroupJid(jid string) bool {
	return strings.Contains(jid, "@g.us")
}

func IsOnWhatsapp(waCli *whatsmeow.Client, jid string) bool {
	if strings.Contains(jid, "@s.whatsapp.net") {
		data, err := waCli.IsOnWhatsApp([]string{jid})
		if err != nil {
			return false
		}

		for _, v := range data {
			if !v.IsIn {
				return false
			}
		}
	}

	return true
}

func ParseJID(arg string) (types.JID, error) {
	if arg[0] == '+' {
		arg = arg[1:]
	}
	if !strings.ContainsRune(arg, '@') {
		return types.NewJID(arg, types.DefaultUserServer), nil
	}

	recipient, err := types.ParseJID(arg)
	if err != nil {
		return recipient, ErrRecipientNotFound
	} else if recipient.User == "" {
		return recipient, ErrRecipientNotFound
	}
	return recipient, nil
}
