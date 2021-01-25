package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/jetexe/pbuf-example/internal/pkg/grpc/messages/v1"
	"google.golang.org/protobuf/proto"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var m messages.Direct

		hexMsg := scanner.Bytes()
		msg := make([]byte, hex.DecodedLen(len(hexMsg)))

		if _, err := hex.Decode(msg, hexMsg); err != nil {
			log.Fatalf("[FATAL]  %v", err)
		}

		if err := proto.Unmarshal(msg, &m); err != nil {
			log.Fatalf("[FATAL]  %v", err)
		}

		fmt.Printf("Account=%s, Text=%s\n", m.Account, m.Text)
	}

	if scanner.Err() != nil {
		log.Fatalf("[FATAL]  %v", scanner.Err())
	}
}
