package main

import (
	"context"
	"fmt"
	"log"

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
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	c := api.NewRadioReaderServiceClient(conn)
/*
	response, err := c.AddSong(context.Background(), &api.Song{Title: "Song Test", Artist: "Yeezus"})

	if err != nil {
		log.Fatalf("Error when calling AddSong: %s", err)
	}

	log.Printf("Song added: %t", response.Created)

 */

	songsResposne, err := c.GetSongs(context.Background(), &api.GetRequest{})

	if err != nil {
		log.Fatalf("Error when getting songs: %s", err)
	}

	log.Printf("Songs are:\n%s", songsResposne.Songs)
}
