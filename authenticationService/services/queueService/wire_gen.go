// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package queueservice

import (
	"github.com/AFORANURAG/microServices/authenticationService/types/queueServiceTypes"
)

// Injectors from wire.go:

func InitializeProducerService(phrase queueservicetypes.QueueServicePhrase) *Producer {
	queueService := NewQueueServiceProvider(phrase)
	producer := NewProducerServiceProvider(queueService)
	return producer
}