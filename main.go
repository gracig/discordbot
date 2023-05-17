package main

import (
	"discordbot/pkg/bots"
	"flag"
	"log"

	"github.com/gracig/mstreamer"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	input, output, err := bots.NewDiscordBot(Token)
	if err != nil {
		log.Fatalf("error creating discord bot")
	}
	pingpong, err := bots.NewPingPongFilter()
	if err != nil {
		log.Fatalf("error creating ping pong filter")
	}

	pipeline, err := mstreamer.NewIOPipeline(input, pingpong, output)

	if err != nil {
		log.Fatalf("Could not create bot pipeline")
	}

	if pipeline(log.Printf) != nil {
		log.Fatalf("Pipeline has failed")
	}
}
