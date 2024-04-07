package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type WebhookPayload struct {
	ObjectKind string `json:"object_kind"`
}

func main() {
	app := fiber.New()

	app.Post("/webhook", handleWebhook)

	fmt.Println("Starting server on :8080")
	log.Fatal(app.Listen(":8080"))
}

func handleWebhook(c *fiber.Ctx) error {
	var payload WebhookPayload
	if err := json.Unmarshal(c.Body(), &payload); err != nil {
		c.Status(http.StatusBadRequest)
		return c.SendString(fmt.Sprintf("Error decoding payload: %v", err))
	}

	switch payload.ObjectKind {
	case "push":
		fmt.Println("Received push event")
	case "merge_request":
		fmt.Println("Received merge request event")
	default:
		fmt.Println("Received unknown event type:", payload.ObjectKind)
	}

	return c.SendString("Webhook received and processed")
}
