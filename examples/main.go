package main

import (
	"github.com/azoff/hipchat"
	"log"
)

func main() {
	c := hipchat.Client{AuthToken: "<PUT YOUR AUTH TOKEN HERE>"}
	req := hipchat.MessageRequest{
		RoomId:        "Rat Man’s Den",
		From:          "GLaDOS",
		Message:       "Bad news: Combustible lemons failed.",
		Color:         hipchat.ColorPurple,
		MessageFormat: hipchat.FormatText,
		Notify:        true,
	}
	if err := c.PostMessage(req); err != nil {
		log.Printf("Expected no error, but got %q", err)
	}
}
