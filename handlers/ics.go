package handlers

import (
	"net/http"

	"wfical/repos"
)

func IcsHandler(w http.ResponseWriter, r *http.Request) {
	streamsData, err := repos.FetchStreamsData()
	if err != nil {
		http.Error(w, "Failed to fetch streams data", http.StatusInternalServerError)
		return
	}

	streams := repos.NewStreams(streamsData)
	icalCalendar := streams.Ics()

	w.Header().Set("Content-Type", "text/calendar")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(icalCalendar.Serialize()))
}
