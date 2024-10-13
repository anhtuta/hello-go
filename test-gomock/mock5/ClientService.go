package mock5

type ClientService interface {
	// DeleteIntegrationsV1ProvidersProviderToken(params *DeleteIntegrationsV1ProvidersProviderTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteIntegrationsV1ProvidersProviderTokenOK, error)

	// GetIntegrationsV1Authorize(params *GetIntegrationsV1AuthorizeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetIntegrationsV1AuthorizeOK, error)

	GetIntegrationsV1ProvidersProviderToken(params string) (string, error)

	// GetIntegrationsV1Token(params *GetIntegrationsV1TokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetIntegrationsV1TokenOK, error)

	// SetTransport(transport runtime.ClientTransport)
}
