package repos

import (
	"testing"
	"time"

	ical "github.com/arran4/golang-ical"
	"github.com/stretchr/testify/assert"
)

func TestNewStream(t *testing.T) {
	row := "Math Lecture,2024-09-01,10:00,2024-09-01,11:00,Room 101,Introduction to Algebra"
	stream := NewStream(row)

	expectedStart, _ := time.Parse("2006-01-02 15:04", "2024-09-01 10:00")
	expectedEnd, _ := time.Parse("2006-01-02 15:04", "2024-09-01 11:00")

	assert.Equal(t, "Math Lecture", stream.Subject)
	assert.Equal(t, expectedStart, stream.Start)
	assert.Equal(t, expectedEnd, stream.End)
	assert.Equal(t, "Room 101", stream.Location)
	assert.Equal(t, "Introduction to Algebra", stream.Description)
}

func TestNewStream_InvalidRow(t *testing.T) {
	row := "Incomplete Data,2024-09-01,10:00"
	stream := NewStream(row)

	assert.Equal(t, Stream{}, stream)
}

func TestNewStreams(t *testing.T) {
	rows := `
Math Lecture,2024-09-01,10:00,2024-09-01,11:00,Room 101,Introduction to Algebra
Physics Lecture,2024-09-01,12:00,2024-09-01,13:00,Room 102,Introduction to Mechanics
`
	streams := NewStreams(rows)

	assert.Len(t, streams, 2)

	expectedStart1, _ := time.Parse("2006-01-02 15:04", "2024-09-01 10:00")
	expectedEnd1, _ := time.Parse("2006-01-02 15:04", "2024-09-01 11:00")
	expectedStart2, _ := time.Parse("2006-01-02 15:04", "2024-09-01 12:00")
	expectedEnd2, _ := time.Parse("2006-01-02 15:04", "2024-09-01 13:00")

	assert.Equal(t, "Math Lecture", streams[0].Subject)
	assert.Equal(t, expectedStart1, streams[0].Start)
	assert.Equal(t, expectedEnd1, streams[0].End)
	assert.Equal(t, "Room 101", streams[0].Location)
	assert.Equal(t, "Introduction to Algebra", streams[0].Description)

	assert.Equal(t, "Physics Lecture", streams[1].Subject)
	assert.Equal(t, expectedStart2, streams[1].Start)
	assert.Equal(t, expectedEnd2, streams[1].End)
	assert.Equal(t, "Room 102", streams[1].Location)
	assert.Equal(t, "Introduction to Mechanics", streams[1].Description)
}

func TestNewStreams_WithInvalidRow(t *testing.T) {
	rows := `
Math Lecture,2024-09-01,10:00,2024-09-01,11:00,Room 101,Introduction to Algebra
Incomplete Data,2024-09-01,10:00
Physics Lecture,2024-09-01,12:00,2024-09-01,13:00,Room 102,Introduction to Mechanics
`
	streams := NewStreams(rows)

	assert.Len(t, streams, 2)

	expectedStart1, _ := time.Parse("2006-01-02 15:04", "2024-09-01 10:00")
	expectedEnd1, _ := time.Parse("2006-01-02 15:04", "2024-09-01 11:00")
	expectedStart2, _ := time.Parse("2006-01-02 15:04", "2024-09-01 12:00")
	expectedEnd2, _ := time.Parse("2006-01-02 15:04", "2024-09-01 13:00")

	assert.Equal(t, "Math Lecture", streams[0].Subject)
	assert.Equal(t, expectedStart1, streams[0].Start)
	assert.Equal(t, expectedEnd1, streams[0].End)
	assert.Equal(t, "Room 101", streams[0].Location)
	assert.Equal(t, "Introduction to Algebra", streams[0].Description)

	assert.Equal(t, "Physics Lecture", streams[1].Subject)
	assert.Equal(t, expectedStart2, streams[1].Start)
	assert.Equal(t, expectedEnd2, streams[1].End)
	assert.Equal(t, "Room 102", streams[1].Location)
	assert.Equal(t, "Introduction to Mechanics", streams[1].Description)
}

func TestStreams_Ics(t *testing.T) {
	rows := `
Math Lecture,2024-09-01,10:00,2024-09-01,11:00,Room 101,Introduction to Algebra
Physics Lecture,2024-09-01,12:00,2024-09-01,13:00,Room 102,Introduction to Mechanics
`
	streams := NewStreams(rows)
	cal := streams.Ics()

	events := cal.Events()
	assert.Len(t, events, 2)

	assert.Equal(t, "Math Lecture", events[0].GetProperty(ical.ComponentPropertySummary).Value)
	assert.Equal(t, "Room 101", events[0].GetProperty(ical.ComponentPropertyLocation).Value)
	assert.Equal(t, "Introduction to Algebra", events[0].GetProperty(ical.ComponentPropertyDescription).Value)

	assert.Equal(t, "Physics Lecture", events[1].GetProperty(ical.ComponentPropertySummary).Value)
	assert.Equal(t, "Room 102", events[1].GetProperty(ical.ComponentPropertyLocation).Value)
	assert.Equal(t, "Introduction to Mechanics", events[1].GetProperty(ical.ComponentPropertyDescription).Value)
}
