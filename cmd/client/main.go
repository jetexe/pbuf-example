package main

import (
	"encoding/hex"
	"log"
	"math/rand"
	"time"

	messages "github.com/jetexe/pbuf-example/api/messages"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var texts = [...]string{
	"Hello there. Thanks for the follow. Did you notice, that I am an egg? A talking egg? Damn!",
	"Thanks mate! Feel way better now",
	"Yeah that is crazy",
	"Hi",
	"Thanks",
	"Okay",
}

func main() {
	rand.Seed(time.Now().Unix())
	for _, account := range []string{"Alice", "Bob", "Charlie", "Daemon"} {
		text := texts[rand.Intn(len(texts))]
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
