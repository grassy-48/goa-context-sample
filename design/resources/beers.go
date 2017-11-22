package resources

import (
	mb "goa-context-sample/design/mediatypes/beers"
	pb "goa-context-sample/design/payloads/beers"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// BeerResource
var _ = Resource("beer", func() {
	BasePath("/beers")
	Files("/index.html", "public/index.html", func() {
		Description("Serve home page")
	})
	Action("teapot", func() {
		Routing(GET("/teapot"))
		Response("TeapotResponse", func() {
			Description("I'm a teapot")
			Headers(func() {
				Header("X-Request-DrinkType", func() {
					Description("Request only tea")
					Enum("Tea", "EarlGrey", "Caylon")
				})
				Required("X-Request-DrinkType")
			})
			Status(418)
		})
	})
	DefaultMedia(mb.BeerMedia)
	Action("empty", func() {
		Routing(GET("/empty"))
	})
	Action("list", func() {
		Description("Get all beers")
		Routing(GET(""))
		Response(OK, func() {
			Media(CollectionOf(mb.BeerMedia))
		})
		Response(NotFound)
	})
	Action("show", func() {
		Description("Get beer by id")
		Routing(GET("/:beerID"))
		Params(func() {
			Param("beerID", Integer, "Beer ID")
		})
		Response(OK)
		Response(NotFound)
	})
	Action("create", func() {
		Description("create new beer")
		Routing(POST(""))
		Payload(pb.BeerPayload)
		Response(Created, func() {
			Media(mb.BeerMedia)
		})
	})
	Action("change", func() {
		Description("change exist beer")
		Routing(PATCH("/:beerID"))
		Params(func() {
			Param("beerID", Integer, "Beer ID")
		})
		Payload(pb.BeerPayload)
		Response(OK)
		Response(NotFound)
	})
	Action("remove", func() {
		Description("remove exist beer")
		Routing(DELETE("/:beerID"))
		Params(func() {
			Param("beerID", Integer, "Beer ID")
		})
		Response(NoContent)
		Response(NotFound)
	})
	Response(BadRequest)
})
