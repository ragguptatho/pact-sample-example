package provider

import (
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
)

// Setup the Pact client.
func createPact() dsl.Pact {
	return dsl.Pact{
		Provider: "provider",
		LogLevel: "DEBUG",
	}
}

func TestProvider(t *testing.T) {
	var state string

	// fn -> message
	functionMappings := dsl.MessageHandlers{
		"produce hello world": func(m dsl.Message) (interface{}, error) {
			return state, nil
		},
	}

	// fn --> setup
	stateMappings := dsl.StateHandlers{
		"say hello world": func(s dsl.State) error {
			state = "Hello World"
			return nil
		},
	}

	var pact = createPact()

	// Verify the Provider with local Pact Files
	pact.VerifyMessageProvider(t, dsl.VerifyMessageRequest{
		PactURLs:                   []string{"http://localhost:9292/pacts/provider/provider/consumer/consumer/latest"},
		MessageHandlers:            functionMappings,
		StateHandlers:              stateMappings,
		PactLogLevel:               "DEBUG",
		PublishVerificationResults: true,
		ProviderVersion:            "0.0.2",
	})

}
