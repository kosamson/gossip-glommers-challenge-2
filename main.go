package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	node := maelstrom.NewNode()

	node.Handle("generate", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return fmt.Errorf("failed to unmarshal message body: %w", err)
		}

		body["type"] = "generate_ok"
		body["id"] = rand.Int()

		return node.Reply(msg, body)
	})

	if err := node.Run(); err != nil {
		log.Fatalf("failed to run node: %s", err)
	}
}
