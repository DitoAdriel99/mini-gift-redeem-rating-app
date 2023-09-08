package logs_service_sdk

import (
	"fmt"
	"go-learn/entities"
	"go-learn/library/httpclient"
	"log"
	"net/http"
	"net/url"
	"os"
)

func (s *_SDK) CreateLogs(payload entities.LogsPayload) error {
	var (
		urlLogs          = fmt.Sprintf("%s/logs", os.Getenv("LOGS_SERVICE"))
		urlLogsParsed, _ = url.Parse(urlLogs)
	)

	reqOptions := httpclient.RequestOptions{
		URL:     urlLogsParsed.String(),
		Payload: payload,
		Method:  http.MethodPost,
	}

	resp, err := httpclient.Request(reqOptions)
	if err != nil {
		log.Println("hit endpoint create logs got err: ", err)
		return err
	}

	switch resp.Status() {
	case http.StatusOK, http.StatusCreated:
		log.Println("success create logs")
		return nil
	default:
		return entities.NewErrRequestWithMessage("error")
	}

}
