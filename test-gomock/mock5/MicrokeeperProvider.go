package mock5

import "context"

type MicrokeeperProvider struct {
	// *providers.Provider
	BaseURL string
}

func (ip *MicrokeeperProvider) Request(ctx context.Context, action string, payloadBytes []byte, accessToken string) ([]byte, error) {
	sampleResponse := []byte(`{"status": "success"}`)
	return sampleResponse, nil
}
