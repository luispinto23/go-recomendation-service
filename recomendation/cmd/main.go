package main

import (
	"log"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"

	"github.com/luispinto23/go-recomendation-service/internal/recommendation"
	"github.com/luispinto23/go-recomendation-service/internal/transport"
)

func main() {
	c := retryablehttp.NewClient()
	c.RetryMax = 10

	partnerAdaptor, err := recommendation.NewPartnerShipAdaptor(
		c.StandardClient(),
		"http://localhost:3031",
	)
	if err != nil {
		log.Fatal("failed to create a partnerAdaptor: ", err)
	}

	svc, err := recommendation.NewService(partnerAdaptor)
	if err != nil {
		log.Fatal("failed to create a service: ", err)
	}

	handler, err := recommendation.NewHandler(*svc)
	if err != nil {
		log.Fatal("failed to create a handler: ", err)
	}

	m := transport.NewMux(*handler)

	if err := http.ListenAndServe(":4040", m); err != nil {
		log.Fatal("server errored: ", err)
	}
}
