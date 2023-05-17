package bots

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gracig/mstreamer"
)

func NewDiscordBot(token string) (mstreamer.Input, mstreamer.Output, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating Discord session - %v", err)
	}
	input, err := mstreamer.NewInputFromProducer(func(f mstreamer.Feedback, w mstreamer.MeasureWriter) {
		dg.Identify.Intents = discordgo.IntentMessageContent
		dg.Identify.Intents |= discordgo.IntentGuildMessages
		dg.Identify.Intents |= discordgo.IntentsMessageContent
		dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
			if m.Author.ID == s.State.User.ID {
				return
			}
			if err := w.Write(mstreamer.Measure{
				Name: DiscordName,
				Time: time.Now().Unix(),
				Tags: []mstreamer.Tag{
					mstreamer.MakeTag(MessageID, m.ChannelID),
					mstreamer.MakeTag(MessageIn, m.Content),
				},
			}); err != nil {
				f("error while writing message %v", err)
			}
		})
		err = dg.Open()
		if err != nil {
			fmt.Println("error opening connection,", err)
			return
		}
		fmt.Println("Bot is now running.  Press CTRL-C to exit.")
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-sc
		dg.Close()
	})
	if err != nil {
		return nil, nil, fmt.Errorf("error creating Discord input listener - %v", err)
	}
	output, err := mstreamer.NewOutput(func(m mstreamer.Measure) error {
		if m.Name != DiscordName {
			return fmt.Errorf("ignoring non-discord message %v", m)
		}
		cid, err := m.TagValue(MessageID)
		if err != nil {
			return fmt.Errorf("channel id not found on message %v", m)
		}
		out, err := m.TagValue(MessageOut)
		if err != nil {
			return fmt.Errorf("message out is empty %v", m)
		}
		_, err = dg.ChannelMessageSend(cid, out)
		return err
	})
	return input, output, nil
}
