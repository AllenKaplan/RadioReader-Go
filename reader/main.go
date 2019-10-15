package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	api "github.com/RadioReader/api"
	"google.golang.org/grpc"
)

const(
	PORT = 14586
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf(":%d", PORT), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %s", err)
	}

	defer conn.Close()

	c := api.NewRadioReaderServiceClient(conn)

	var result map[string]string
	var lastSong string

	nextTime := time.Now().Add(time.Minute*2)

	//TODO make get last song function in api

	for {
		musicResponse, err := http.Get("https://cob.leanplayer.com/CINDFM/nowplaying")
		if err != nil {
			log.Fatalf("Failed to get current song from web player: %s", err)
		}

		if err = json.NewDecoder(musicResponse.Body).Decode(&result); err != nil {
			log.Printf("Unable to decode web player current song HTTP request: %s", err)
		}
		log.Printf("Decoded | Title: %s | Artist: %s", result["title"], result["artist"])

		if lastSong == result["title"] {
			log.Printf("Song not changed | %s", result)
			continue
		}

		lastSong = result["title"]


		response, err := c.AddSong(context.Background(), &api.Song{
			Title: result["title"],
			Artist: result["artist"],
		})

		if err != nil {
			log.Fatalf("Error when calling AddSong: %s", err)
		}
		log.Printf("Song added: %t", response.Created)

		time.Sleep(time.Until(nextTime))
		nextTime = time.Now().Add(time.Minute*2)
	}

}
