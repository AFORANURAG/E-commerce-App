// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package authService

import (
	"github.com/AFORANURAG/microServices/authenticationService/clientProviders/emailServiceClient"
	"github.com/AFORANURAG/microServices/authenticationService/services/userService"
)

// Injectors from wire.go:

func InitializeAuthenticationService(phrase string) AuthenticationServiceServer {
	userServiceUserService := userService.InitializeUserService(phrase)
	emailServiceClient := emailclient.NewEmailServiceServiceClientProvider()
	authenticationServiceServer := NewAuthenticationServiceProvider(userServiceUserService, emailServiceClient)
	return authenticationServiceServer
}
