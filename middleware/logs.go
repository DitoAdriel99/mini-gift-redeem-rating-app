package middleware

import (
	"bytes"
	"encoding/json"
	"go-learn/entities"
	"go-learn/library/jwt_parse"
	"go-learn/sdk"
	"io"
	"net/http"

	"github.com/jmoiron/sqlx/types"
)

type LogsIntegrate struct {
	SDK sdk.SDK
}

func NewLogsIntegrate(SDK sdk.SDK) *LogsIntegrate {
	return &LogsIntegrate{
		SDK: SDK,
	}
}

func (l *LogsIntegrate) CreateLogs(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		endpointPath := r.URL.Path

		rw := &responseWriterWithStatus{ResponseWriter: w}

		reqBody, _ := io.ReadAll(r.Body)
		defer r.Body.Close()
		r.Body = io.NopCloser(bytes.NewBuffer(reqBody))

		// Continue processing the request
		next.ServeHTTP(rw, r)

		// Capture the response body from the custom response writer
		respBody := rw.responseBuf.Bytes()
		// Create variables to hold the JSON data
		var reqJSON types.JSONText
		var respJSON types.JSONText

		// Unmarshal the request and response bodies into the JSON variables
		json.Unmarshal(reqBody, &reqJSON)
		json.Unmarshal(respBody, &respJSON)

		payload := entities.LogsPayload{
			Event:              endpointPath,
			UserAgent:          r.Header.Get("User-Agent"),
			HttpStatusCode:     rw.statusCode,
			HttpMethod:         r.Method,
			ClientRequestData:  reqJSON,
			ClientResponseData: respJSON,
		}

		bearer := r.Header.Get("Authorization")
		if bearer != "" {
			claims, _ := jwt_parse.GetClaimsFromToken(bearer)
			payload.Fullname = claims.Username
			payload.Email = claims.Email
		}

		payload.Fullname = "-"
		payload.Email = "-"

		l.SDK.LogsSDK.CreateLogs(payload)

	})
}

// Define a custom response writer that captures the status code
type responseWriterWithStatus struct {
	http.ResponseWriter
	statusCode  int
	responseBuf bytes.Buffer // To capture the response data
}

func (rw *responseWriterWithStatus) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriterWithStatus) Write(data []byte) (int, error) {
	// Capture the response data
	n, err := rw.responseBuf.Write(data)
	if err != nil {
		return n, err
	}

	// Write the data to the original ResponseWriter
	return rw.ResponseWriter.Write(data)
}
