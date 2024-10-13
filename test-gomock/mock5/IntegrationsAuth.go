package mock5

import "context"

type IntegrationsAuth struct {
	// *services.Base
	client ClientService
}

func (ia *IntegrationsAuth) GetAccessToken(ctx context.Context) (string, error) {
	resp, err := ia.client.GetIntegrationsV1ProvidersProviderToken("sample_params")
	if err != nil {
		return "", err
	}
	return resp, nil
}
