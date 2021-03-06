//go:generate goagen bootstrap -d github.com/minhkhiemm/hcm/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/minhkhiemm/hcm/app"
)

func main() {
	// Create service
	service := goa.New("hcm")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "partients" controller
	c := NewPartientsController(service)
	app.MountPartientsController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
