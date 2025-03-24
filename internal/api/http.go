package api

import (
	"go-automation/internal/domain/device"
	"go-automation/internal/domain/user"
	"go-automation/pkg/config"

	"github.com/gin-gonic/gin"
)

func StartHTTPServer(cfg *config.Config) error {
	r := gin.Default()

	user.RegisterRoutes(r.Group("/users"))
	device.RegisterRoutes(r.Group("/devices"))

	return r.Run(cfg.Server.Address)
}
