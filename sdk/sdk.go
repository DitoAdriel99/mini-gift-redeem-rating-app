package sdk

import "go-learn/sdk/logs_service_sdk"

type SDK struct {
	LogsSDK logs_service_sdk.Contract
}

func NewSDK() *SDK {
	return &SDK{
		LogsSDK: logs_service_sdk.NewSDK(),
	}
}
