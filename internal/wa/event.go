package wa //nolint

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/hrz8/goatsapp/pkg/whatsapp"
	"github.com/hrz8/gofx"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/appstate"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

type Event struct {
	log *gofx.Logger
}

func NewEvent(log *gofx.Logger) *Event {
	return &Event{log}
}

func (e *Event) OnPairSuccess(v *events.PairSuccess) {
	// websocket.Broadcast <- websocket.BroadcastMessage{
	// 		Code:    "LOGIN_SUCCESS",
	// 		Message: fmt.Sprintf("Successfully pair with %s", v.ID.String()),
	// 	}
}

func (e *Event) OnLoggedOut(v *events.LoggedOut) {
	// 	websocket.Broadcast <- websocket.BroadcastMessage{
	// 		Code:   "LIST_DEVICES",
	// 		Result: nil,
	// 	}
}

func (e *Event) OnDelete(v *events.DeleteForMe) {
	// statement here ...
}

func (e *Event) OnAppState(v *events.AppState) {
	// statement here ...
}

func (e *Event) OnReplaced(v *events.StreamReplaced) {
	// statement here ...
}

func (e *Event) OnAppSyncComplete(v *events.AppStateSyncComplete, cli *whatsmeow.Client, sessionID string) {
	if len(cli.Store.PushName) > 0 && v.Name == appstate.WAPatchCriticalBlock {
		err := cli.SendPresence(types.PresenceAvailable)
		if err != nil {
			e.log.Error("failed to send available presence", slog.String("err", err.Error()))
		} else {
			e.log.Info("marked self as available", slog.String("session_id", sessionID))
		}
	}
}

func (e *Event) OnConnected(cli *whatsmeow.Client, sessionID string) {
	if len(cli.Store.PushName) == 0 {
		return
	}
	err := cli.SendPresence(types.PresenceAvailable)
	if err != nil {
		e.log.Error("failed to send available presence", slog.String("err", err.Error()))
	} else {
		e.log.Info("marked self as available", slog.String("session_id", sessionID))
	}
}

func (e *Event) OnMessage(v *events.Message) {
	metaParts := []string{fmt.Sprintf("pushName: %s", v.Info.PushName)}
	if v.Info.Type != "" {
		metaParts = append(metaParts, fmt.Sprintf("type: %s", v.Info.Type))
	}
	if v.Info.Category != "" {
		metaParts = append(metaParts, fmt.Sprintf("category: %s", v.Info.Category))
	}
	if v.IsViewOnce {
		metaParts = append(metaParts, "view once")
	}

	e.log.Info(
		fmt.Sprintf(
			"received message %s from %s (%s)",
			v.Info.ID,
			v.Info.SourceString(),
			strings.Join(metaParts, ", "),
		),
		slog.Any("message", v.Message),
	)

	// img := v.Message.GetImageMessage()
	// if img != nil {
	// 	path, err := ExtractMedia(config.PathStorages, img)
	// 	if err != nil {
	// 		log.Errorf("Failed to download image: %v", err)
	// 	} else {
	// 		log.Infof("Image downloaded to %s", path)
	// 	}
	// }

	// if config.WhatsappAutoReplyMessage != "" &&
	// 	!isGroupJid(v.Info.Chat.String()) &&
	// 	!strings.Contains(v.Info.SourceString(), "broadcast") {
	// 	_, _ = cli.SendMessage(
	// 		context.Background(),
	// 		v.Info.Sender,
	// 		&waE2E.Message{Conversation: proto.String(config.WhatsappAutoReplyMessage)},
	// 	)
	// }

	// if config.WhatsappWebhook != "" &&
	// 	!strings.Contains(v.Info.SourceString(), "broadcast") &&
	// 	!isFromMySelf(v.Info.SourceString()) {
	// 	if err := forwardToWebhook(evt); err != nil {
	// 		logrus.Error("Failed forward to webhook", err)
	// 	}
	// }
}

func (e *Event) OnReceipt(v *events.Receipt) {
	if v.Type == types.ReceiptTypeRead || v.Type == types.ReceiptTypeReadSelf {
		e.log.Info(fmt.Sprintf("%v was read by %s at %s", v.MessageIDs, v.SourceString(), v.Timestamp))
	} else if v.Type == types.ReceiptTypeDelivered {
		e.log.Info(fmt.Sprintf("%s was delivered to %s at %s", v.MessageIDs[0], v.SourceString(), v.Timestamp))
	}
}

func (e *Event) OnPresence(v *events.Presence) {
	if v.Unavailable {
		if v.LastSeen.IsZero() {
			e.log.Info(fmt.Sprintf("%s is now offline", v.From))
		} else {
			e.log.Info(fmt.Sprintf("%s is now offline (last seen: %s)", v.From, v.LastSeen))
		}
	} else {
		e.log.Info(fmt.Sprintf("%s is now online", v.From))
	}
}

func (e *Event) OnHistorySync(v *events.HistorySync) {
	// 	id := atomic.AddInt32(&historySyncID, 1)
	// 	fileName := fmt.Sprintf("%s/history-%d-%s-%d-%s.json", config.PathStorages, startupTime, cli.Store.ID.String(), id, v.Data.SyncType.String())
	// 	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0600)
	// 	if err != nil {
	// 		log.Errorf("Failed to open file to write history sync: %v", err)
	// 		return
	// 	}
	// 	enc := json.NewEncoder(file)
	// 	enc.SetIndent("", "  ")
	// 	err = enc.Encode(v.Data)
	// 	if err != nil {
	// 		log.Errorf("Failed to write history sync: %v", err)
	// 		return
	// 	}
	// 	log.Infof("Wrote history sync to %s", fileName)
	// 	_ = file.Close()
}

func RegisterHandlers(log *gofx.Logger, e *Event) whatsapp.EventHandler { //nolint
	return func(cli *whatsmeow.Client, sessionID string) whatsmeow.EventHandler {
		return func(evt any) {
			switch v := evt.(type) {
			case *events.DeleteForMe:
				log.JSON.Info(
					fmt.Sprintf("deleted message %s for %s", v.MessageID, v.SenderJID.String()),
					slog.String("event", "whatsmeow.DeleteForMe"),
					slog.String("session_id", sessionID),
				)
				e.OnDelete(v)
			case *events.AppStateSyncComplete:
				log.JSON.Info(
					fmt.Sprintf("app sync complete for %s", v.Name),
					slog.String("event", "whatsmeow.AppStateSyncComplete"),
					slog.String("session_id", sessionID),
				)
				e.OnAppSyncComplete(v, cli, sessionID)
			case *events.PairSuccess:
				log.JSON.Info(
					fmt.Sprintf("success pair with %s", v.ID.String()),
					slog.String("event", "whatsmeow.PairSuccess"),
					slog.String("session_id", sessionID),
				)
				e.OnPairSuccess(v)
			case *events.LoggedOut:
				log.JSON.Info(
					fmt.Sprintf("logged out success with reason %s", v.Reason.String()),
					slog.String("event", "whatsmeow.LoggedOut"),
					slog.String("session_id", sessionID),
				)
				e.OnLoggedOut(v)
			case *events.Connected, *events.PushNameSetting:
				log.JSON.Info(
					fmt.Sprintf("updating push name %s", cli.Store.PushName),
					slog.String("event", "whatsmeow.Connected.PushNameSetting"),
					slog.String("session_id", sessionID),
				)
				e.OnConnected(cli, sessionID)
			case *events.StreamReplaced:
				log.JSON.Error(
					fmt.Sprintf("stream replaced: %s", v.PermanentDisconnectDescription()),
					slog.String("event", "whatsmeow.StreamReplaced"),
					slog.String("session_id", sessionID),
				)
				e.OnReplaced(v)
			case *events.Message:
				log.JSON.Info(
					fmt.Sprintf("message received %v, type: %v", v.Info.PushName, v.Info.Type),
					slog.String("event", "whatsmeow.Message"),
					slog.String("session_id", sessionID),
				)
				e.OnMessage(v)
			case *events.Receipt:
				log.JSON.Info(
					fmt.Sprintf("event receipt triggered for type %v, source: %v", v.Type, v.SourceString()),
					slog.String("event", "whatsmeow.Receipt"),
					slog.String("session_id", sessionID),
				)
				e.OnReceipt(v)
			case *events.Presence:
				log.JSON.Info(
					fmt.Sprintf("update presence from %s: %v", v.From, v.Unavailable),
					slog.String("event", "whatsmeow.Presence"),
					slog.String("session_id", sessionID),
				)
				e.OnPresence(v)
			case *events.HistorySync:
				log.JSON.Info(
					fmt.Sprintf("history synced (%s)", v.Data.String()),
					slog.String("event", "whatsmeow.HistorySync"),
					slog.String("session_id", sessionID),
				)
				e.OnHistorySync(v)
			case *events.AppState:
				log.JSON.Info(
					fmt.Sprintf("app state event: %+v / %+v", v.Index, v.SyncActionValue),
					slog.String("event", "whatsmeow.AppState"),
					slog.String("session_id", sessionID),
				)
				e.OnAppState(v)
			}
		}
	}
}
