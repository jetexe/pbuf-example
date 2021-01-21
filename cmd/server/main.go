package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	messages "github.com/korjavin/pbuf-example/messages"
	"google.golang.org/protobuf/proto"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var m messages.Direct
		hexMsg := scanner.Bytes()
		msg := make([]byte, hex.DecodedLen(len(hexMsg)))
		_, err := hex.Decode(msg, hexMsg)
		if err != nil {
			log.Fatalf("[FATAL]  %v", err)
		}

		err = proto.Unmarshal(msg, &m)
		if err != nil {
			log.Fatalf("[FATAL]  %v", err)
		}
		fmt.Printf("Account=%s, Text=%s\n", m.Account, m.Text)
	}

	if scanner.Err() != nil {
		log.Fatalf("[FATAL]  %v", scanner.Err())
	}
}
