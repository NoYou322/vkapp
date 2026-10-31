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
	"github.com/syumai/workers-go/cloudflare"
)

func requiredEnv(key string) string {
	value := cloudflare.Getenv(key)

	if value == "" {
		panic("missing environment variable: " + key)
	}

	return value
}

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
	vk := api.NewVK(requiredEnv("TOKEN"))

	cb := callback.NewCallback()
	cb.ConfirmationKey = requiredEnv("CONFIRMATION_KEY")
	cb.SecretKey = cloudflare.Getenv("VK_SECRET")	

	cb.MessageNew(messagenew.New(log, vk))
	cb.OnEvent(EventMessageReaction, messagereaction.New(log, vk))

	mux := http.NewServeMux()

	mux.HandleFunc("POST /callback", cb.HandleFunc)

	log.Info("Server starting")

	workers.Serve(mux)


}


