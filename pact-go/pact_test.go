package main

import (
	"log"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestPactProvider(t *testing.T) {
	go main() // Inicie seu servidor aqui

	pact := &dsl.Pact{
		Consumer: "NodeService",
		Provider: "GoService",
		Host:     "127.0.0.1",
	}
	defer pact.Teardown()

	pact.
		AddInteraction().
		Given("Dado um cenário específico").
		UponReceiving("Uma requisição para criar um usuário").
		WithRequest(dsl.Request{
			Method: "POST",
			Path:   dsl.String("/user"),
			Body: map[string]interface{}{
				"email": "test@example.com",
				"cpf":   "12345678909",
			},
		}).
		WillRespondWith(dsl.Response{
			Status: 200,
		})

	verifyReq := types.VerifyRequest{
		ProviderBaseURL: "http://127.0.0.1:8080",
		BrokerURL:       "http://127.0.0.1:80",
		ProviderVersion: "1.0.0",
	}

	_, err := pact.VerifyProvider(t, verifyReq)
	if err != nil {
		log.Fatalf("Error on Verify Provider: %v", err)
	}

}
