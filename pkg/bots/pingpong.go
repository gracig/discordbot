package bots

import (
	"github.com/gracig/mstreamer"
)

func NewPingPongFilter() (mstreamer.Filter, error) {
	return mstreamer.NewFilter(
		func(f mstreamer.Feedback, m *mstreamer.Measure, mw mstreamer.MeasureWriter) {
			msgIn, err := m.TagValue(MessageIn)
			if err != nil {
				f("%v", err)
				mw.Write(*m)
				return
			}
			if msgIn != "ping" && msgIn != "pong" && msgIn != "ping ping" {
				mw.Write(*m)
				return
			}
			var msgOut string
			switch msgIn {
			case "ping":
				msgOut = "pong"
			case "pong":
				msgOut = "ping"
			}
			m.InsertOrUpdateTag(MessageOut, msgOut)
			mw.Write(*m)
		}, nil)
}
