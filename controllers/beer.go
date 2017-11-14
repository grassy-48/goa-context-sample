package controllers

import (
	"goa-context-sample/app"

	"github.com/goadesign/goa"
)

// BeerController implements the beer resource.
type BeerController struct {
	*goa.Controller
}

// NewBeerController creates a beer controller.
func NewBeerController(service *goa.Service) *BeerController {
	return &BeerController{Controller: service.NewController("BeerController")}
}

// Change runs the change action.
func (c *BeerController) Change(ctx *app.ChangeBeerContext) error {
	// BeerController_Change: start_implement

	// Put your logic here

	// BeerController_Change: end_implement
	res := &app.GoaContextSampleBeer{}
	return ctx.OK(res)
}

// Create runs the create action.
func (c *BeerController) Create(ctx *app.CreateBeerContext) error {
	// BeerController_Create: start_implement

	// Put your logic here

	// BeerController_Create: end_implement
	return nil
}

// List runs the list action.
func (c *BeerController) List(ctx *app.ListBeerContext) error {
	// BeerController_List: start_implement

	// Put your logic here

	// BeerController_List: end_implement
	res := app.GoaContextSampleBeerCollection{}
	return ctx.OK(res)
}

// Remove runs the remove action.
func (c *BeerController) Remove(ctx *app.RemoveBeerContext) error {
	// BeerController_Remove: start_implement

	// Put your logic here

	// BeerController_Remove: end_implement
	return nil
}

// Show runs the show action.
func (c *BeerController) Show(ctx *app.ShowBeerContext) error {
	// BeerController_Show: start_implement

	// Put your logic here

	// BeerController_Show: end_implement
	res := &app.GoaContextSampleBeer{}
	return ctx.OK(res)
}
