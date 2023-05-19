package bots

import (
	"time"

	"github.com/gracig/mstreamer"
)

func NewGreetingFilter() (mstreamer.Filter, error) {
	return mstreamer.NewFilter(
		func(f mstreamer.Feedback, m *mstreamer.Measure, mw mstreamer.MeasureWriter) {
			msgIn, err := m.TagValue(MessageIn)
			if err != nil {
				f("%v", err)
				mw.Write(*m)
				return
			}
			if msgIn != "olá" {
				mw.Write(*m)
				return
			}
			var h, _, _ = time.Unix(m.Time, 0).Clock()
			var msgOut = "Good Morning"
			if h >= 12 && h < 18 {
				msgOut = "Good Afternoon"
			} else if h >= 18 {
				msgOut = "Good Evening"
			}
			m.InsertOrUpdateTag(MessageOut, msgOut)
			mw.Write(*m)
		}, nil)
}
