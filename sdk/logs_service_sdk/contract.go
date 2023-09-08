package logs_service_sdk

import "go-learn/entities"

type _SDK struct {
}

type Contract interface {
	CreateLogs(payload entities.LogsPayload) error
}

func NewSDK() Contract {
	return &_SDK{}
}
