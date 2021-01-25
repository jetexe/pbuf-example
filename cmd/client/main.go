package main

import (
	"encoding/hex"
	"log"
	"math/rand"
	"time"

	messages "github.com/jetexe/pbuf-example/internal/pkg/grpc/messages/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	rand.Seed(time.Now().Unix())

	texts := [...]string{
		"Hello there. Thanks for the follow. Did you notice, that I am an egg? A talking egg? Damn!",
		"Thanks mate! Feel way better now",
		"Yeah that is crazy",
		"Hi",
		"Thanks",
		"Okay",
	}

	for _, account := range []string{"Alice", "Bob", "Charlie", "Daemon"} {
		text := texts[rand.Intn(len(texts))] //nolint:gosec // no need additional crypto here
		m := messages.Direct{
			Account: account,
			Text:    text,
		}

		bytes, err := proto.Marshal(&m)
		if err != nil {
			log.Fatalf("[FATAL]  %v", err)
		}

		log.Printf("[INFO] Message(hex )=%s", hex.EncodeToString(bytes))

		bytes, err = protojson.Marshal(&m)
		if err != nil {
			log.Fatalf("[FATAL]  %v", err)
		}

		log.Printf("[INFO] Message(json)=%s", bytes)
	}
}
