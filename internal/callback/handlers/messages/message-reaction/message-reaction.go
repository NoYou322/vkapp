package messagereaction

import (
	"context"
	"encoding/json"
	"log/slog"
	"strconv"

	"github.com/SevereCloud/vksdk/v3/api"
	"github.com/SevereCloud/vksdk/v3/events"
)

type ReactionObject struct {
	ReactedID  int `json:"reacted_id"`
	PeerID     int `json:"peer_id"`
	CmID       int `json:"cmid"`
	ReactionID int `json:"reaction_id"`
}

func New(log *slog.Logger, vk *api.VK) func(context.Context, events.GroupEvent) {
	return func(ctx context.Context, e events.GroupEvent) {
		const op = "callback.handlers.messages.massage-reaction.New"

		log = log.With(
			slog.String("op", op),
		)

		var obj ReactionObject

		err := json.Unmarshal(e.Object, &obj)
		if err != nil {
			log.Error("json unmarshal", slog.String("error", err.Error()))

			return
		}

		log.Info("Reaction received", slog.Int("from", obj.ReactedID), slog.Int("peer_id", obj.PeerID))

		vk.MessagesSend(api.Params{
			"peer_id": obj.PeerID,
			"message": "Reaction ID = " + strconv.Itoa(obj.ReactionID),
		})
	}
}
