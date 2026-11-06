package main

import (
	"log/slog"
	"net/http"
	"os"
	messagenew "vkbot/internal/callback/handlers/messages/message-new"
	messagereaction "vkbot/internal/callback/handlers/messages/message-reaction"

	"github.com/SevereCloud/vksdk/v3/api"
	"github.com/SevereCloud/vksdk/v3/callback"
	"github.com/syumai/workers-go"

)


const EventMessageReaction = "callback_message_reaction_event"

func main() {
	// TODO config

	// TODO logger
	log := slog.New(slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))

	// TODO init vk
	vk := api.NewVK(os.Getenv("TOKEN"))

cb := callback.NewCallback()
cb.ConfirmationKey = os.Getenv("CONFIRMATION_KEY")
cb.SecretKey = os.Getenv("VK_SECRET")

	cb.MessageNew(messagenew.New(log, vk))
	cb.OnEvent(EventMessageReaction, messagereaction.New(log, vk))

	mux := http.NewServeMux()

	mux.HandleFunc("POST /callback", cb.HandleFunc)

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))

	})

	log.Info("Server starting")

	workers.Serve(mux)


}


