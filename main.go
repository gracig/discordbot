package main

import (
	"discordbot/pkg/bots"
	"log"
	"os"

	"github.com/gracig/mstreamer"
)

func main() {

	input, output, err := bots.NewDiscordBot(os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatalf("error creating discord bot")
	}

	pingpong, err := bots.NewPingPongFilter()
	if err != nil {
		log.Fatalf("error creating ping pong filter")
	}

	greeting, err := bots.NewGreetingFilter()
	if err != nil {
		log.Fatalf("error creating greeting filter")
	}

	filter, err := mstreamer.NewComposedFilter(pingpong, greeting)
	if err != nil {
		log.Fatalf("error creating composed filter")
	}

	pipeline, err := mstreamer.NewIOPipeline(input, filter, output)

	if err != nil {
		log.Fatalf("Could not create bot pipeline")
	}

	if pipeline(log.Printf) != nil {
		log.Fatalf("Pipeline has failed")
	}
}
