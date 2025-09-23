package interfaces

import "gin-starter/config"

// Cronjob is an interface that defines the methods that a cronjob must implement.
type Cronjob interface {
	// ProcessCronjob is a function interface to run cronjob
	ProcessCronjob(cfg config.Config) error
}
