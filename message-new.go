package messagenew

import (
	"context"
	"log/slog"

	"github.com/SevereCloud/vksdk/v3/api"
	"github.com/SevereCloud/vksdk/v3/events"
)

func New(log *slog.Logger, vk *api.VK) func(context.Context, events.MessageNewObject) {
	return func(ctx context.Context, obj events.MessageNewObject) {
		const op = "handlers.messages.message-new.New"

		log = log.With(slog.String("op", op))

		_, err := vk.MessagesSend(api.Params{
			"peer_id": obj.Message.PeerID,
			"message": "hello, world",
		})
		if err != nil {
			log.Error("send message error", slog.String("error", err.Error()))

			return
		}

		log.Info("message sended")
	}
}
