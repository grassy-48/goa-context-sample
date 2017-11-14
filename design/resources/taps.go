package resources

import (
	mt "goa-context-sample/design/mediatypes/taps"
	pt "goa-context-sample/design/payloads/taps"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// TapResource
var _ = Resource("tap", func() {
	BasePath("/taps")
	DefaultMedia(mt.TapMedia)
	Action("list", func() {
		Description("Get all taps")
		Routing(GET(""))
		Response(OK, func() {
			Media(CollectionOf(mt.TapMedia))
		})
		Response(NotFound)
	})
	Action("show", func() {
		Description("Get beerInfo by tap id")
		Routing(GET("/:tapID"))
		Params(func() {
			Param("tapID", Integer, "Tap ID")
		})
		Response(OK)
		Response(NotFound)
	})
	Action("create", func() {
		Description("create bew tap")
		Routing(POST(""))
		Payload(pt.TapPayload)
		Response(OK)
		Response(NotFound)
	})
	Action("change", func() {
		Description("change beer ID related tap")
		Routing(PATCH("/:tapID"))
		Params(func() {
			Param("tapID", Integer, "Tap ID")
		})
		Payload(pt.TapPayload)
		Response(OK)
		Response(NotFound)
	})
	Action("remove", func() {
		Description("remove exist tap")
		Routing(DELETE("/:tapID"))
		Params(func() {
			Param("tapID", Integer, "Tap ID")
		})
		Response(NoContent)
		Response(NotFound)
	})
	Response(BadRequest)
})
