package repos

import (
	"fmt"
	"strings"
	"time"

	ical "github.com/arran4/golang-ical"
)

type Streams []Stream

func NewStreams(rows string) Streams {
	var streams Streams
	for _, row := range strings.Split(strings.TrimSpace(rows), "\n") {
		stream := NewStream(row)
		if (stream != Stream{}) {
			streams = append(streams, stream)
		}
	}
	return streams
}

func (s Streams) Ics() *ical.Calendar {
	cal := ical.NewCalendar()
	cal.SetMethod(ical.MethodPublish)
	cal.SetProductId("-//WFICAL//EN")

	for _, stream := range s {
		event := cal.AddEvent(fmt.Sprintf("%s@wfical", stream.Start.Format("20060102T150405Z")))
		event.SetSummary(stream.Subject)
		event.SetStartAt(stream.Start)
		event.SetEndAt(stream.End)
		event.SetLocation(stream.Location)
		event.SetDescription(stream.Description)
	}

	return cal
}

// Stream represents a single stream event.
type Stream struct {
	Subject     string
	Start       time.Time
	End         time.Time
	Location    string
	Description string
}

func NewStream(row string) Stream {
	cells := strings.Split(row, ",")
	if len(cells) < 7 {
		return Stream{}
	}
	start, _ := time.Parse("2006-01-02 15:04", fmt.Sprintf("%s %s", cells[1], cells[2]))
	end, _ := time.Parse("2006-01-02 15:04", fmt.Sprintf("%s %s", cells[3], cells[4]))
	return Stream{
		Subject:     cells[0],
		Start:       start,
		End:         end,
		Location:    cells[5],
		Description: cells[6],
	}
}
