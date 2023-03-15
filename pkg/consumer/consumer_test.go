package consumer

import (
	"fmt"
	"go/types"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
)

func createPact() dsl.Pact {
	return dsl.Pact{
		Consumer: "consumer",
		Provider: "provider",
		LogLevel: "INFO",
	}
}

func TestConsumer(t *testing.T) {
	var pact = createPact()
	var msgHandlerWrapper = func(m dsl.Message) error {
		fmt.Println(m.Content.(string))
		return nil
	}

	var msg string = "Hello Worldss"

	message := pact.AddMessage()
	message.
		Given("say hello world").
		ExpectsToReceive("produce hello world").
		WithContent(msg).AsType(types.String)

	// unit test for consumer , handle the message/response

	pact.VerifyMessageConsumer(t, message, msgHandlerWrapper)
}
