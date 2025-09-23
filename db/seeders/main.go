// Package auth provides auth
package main

import (
	"github.com/gin-gonic/gin"

	"gin-starter/common/helper"
	"gin-starter/config"
)

func main() {
	cfg, err := config.LoadConfig(".env")
	checkError(err)

	_, err = helper.NewPostgresGormDB(&cfg.Postgres)
	checkError(err)

	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
