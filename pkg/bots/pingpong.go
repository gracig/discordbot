package bots

import (
	"fmt"

	"github.com/gracig/mstreamer"
)

func NewPingPongFilter() (mstreamer.Filter, error) {
	return mstreamer.NewFilter(
		func(f mstreamer.Feedback, m *mstreamer.Measure, mw mstreamer.MeasureWriter) {
			msgIn, err := m.TagValue(MessageIn)
			if err != nil {
				f("%v", err)
				return
			}
			if msgIn != "ping" && msgIn != "pong" {
				return
			}
			for i := 0; i < 50; i++ {
				var msgOut string
				switch msgIn {
				case "ping":
					msgOut = "pong"
				case "pong":
					msgOut = "ping"
				}
				m.InsertOrUpdateTag(MessageOut, fmt.Sprintf("%v-%v", msgOut, i))
				mw.Write(*m)
			}

		}, nil)
}
