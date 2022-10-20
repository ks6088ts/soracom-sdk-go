package test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ks6088ts/soracom-sdk-go/generated/api"
)

func getServerIndex(coverageType string) (int, error) {
	switch coverageType {
	case "jp":
		return 0, nil
	case "g":
		return 1, nil
	}
	return -1, fmt.Errorf("Invalid coverageType = %v specified", coverageType)
}

type SoracomClient struct {
	Client       *api.APIClient
	AuthResponse *api.AuthResponse
	serverIndex  int
}

func (c *SoracomClient) GetContext(ctx context.Context) context.Context {
	contextApiKeys := map[string]api.APIKey{
		"api_key": {
			Key: *c.AuthResponse.ApiKey,
		},
		"api_token": {
			Key: *c.AuthResponse.Token,
		},
	}
	authCtx := context.WithValue(ctx, api.ContextAPIKeys, contextApiKeys)
	return context.WithValue(authCtx, api.ContextServerIndex, c.serverIndex)
}

var (
	globalSoracomClient *SoracomClient
)

//nolint:staticcheck, no need to call os.Exit() from go 1.15+
func TestMain(m *testing.M) {
	authKeyId := os.Getenv("SORACOM_AUTH_KEY_ID")
	authKey := os.Getenv("SORACOM_AUTH_KEY")
	coverageType := os.Getenv("COVERAGE_TYPE")

	authRequest := api.AuthRequest{
		AuthKey:   &authKey,
		AuthKeyId: &authKeyId,
	}

	configuration := api.NewConfiguration()
	apiClient := api.NewAPIClient(configuration)

	resp, r, err := apiClient.AuthApi.Auth(context.Background()).AuthRequest(authRequest).Execute()
	if err != nil {
		log.Fatalf("AuthApi.Auth() failed with error = %v", err)
	}
	if r.StatusCode != 200 {
		log.Fatalf("AuthApi.Auth() failed with status code = %v", r.StatusCode)
	}
	serverIndex, err := getServerIndex(coverageType)
	if err != nil {
		log.Fatalf("coverageType is invalid, err = %v", err)
	}

	globalSoracomClient = &SoracomClient{
		Client:       apiClient,
		AuthResponse: resp,
		serverIndex:  serverIndex,
	}

	m.Run()
}
