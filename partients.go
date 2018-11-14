package main

import (
	"github.com/goadesign/goa"
	"github.com/k0kubun/pp"
	"github.com/minhkhiemm/hcm/app"
)

// PartientsController implements the partients resource.
type PartientsController struct {
	*goa.Controller
}

// NewPartientsController creates a partients controller.
func NewPartientsController(service *goa.Service) *PartientsController {
	return &PartientsController{Controller: service.NewController("PartientsController")}
}

var Patients = []*app.Patient{}

// Create runs the create action.
func (c *PartientsController) Create(ctx *app.CreatePartientsContext) error {
	// PartientsController_Create: start_implement

	// Put your logic here
	Patients = append(Patients,&app.Patient{
		Name:ctx.Payload.Name,
		DateOfBirth:ctx.Payload.DateOfBirth,
		ID:ctx.Payload.ID,
	})
	return ctx.Created()
	// PartientsController_Create: end_implement
}

// List runs the list action.
func (c *PartientsController) List(ctx *app.ListPartientsContext) error {
	// PartientsController_List: start_implement

	// Put your logic here

	pp.Println(Patients)
	return ctx.OK(Patients)
	// PartientsController_List: end_implement
}
