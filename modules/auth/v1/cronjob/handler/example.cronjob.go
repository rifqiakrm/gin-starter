package handler

import (
	"log"

	"gin-starter/config"
	"gin-starter/modules/auth/v1/service"
)

// ExampleCronjob struct
type ExampleCronjob struct {
	userFinderSvc service.UserFinderUseCase
}

// NewExampleCronjobHandler create cronjob handler
func NewExampleCronjobHandler(userFinderSvc service.UserFinderUseCase) *ExampleCronjob {
	return &ExampleCronjob{
		userFinderSvc: userFinderSvc,
	}
}

// ProcessCronjob is a function for processing cronjob
func (c *ExampleCronjob) ProcessCronjob(_ config.Config) error {
	log.Println("Running cronjob: example")

	// TODO: create business logic

	return nil
}
