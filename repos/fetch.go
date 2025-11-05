package repos

import (
	"io"
	"net/http"

	"wfical/config"

	log "github.com/sirupsen/logrus"
)

func FetchStreamsData() (string, error) {
	resp, err := http.Get(config.GetEnv().StreamsURL)
	if err != nil {
		log.WithError(err).Error("Failed to fetch streams data")
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.WithFields(log.Fields{
			"status_code": resp.StatusCode,
		}).Error("Failed to fetch streams data")
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithError(err).Error("Failed to read streams data")
		return "", err
	}

	return string(body), nil
}
