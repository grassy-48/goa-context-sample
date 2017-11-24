package taps

import (
	tt "goa-context-sample/design/types/taps"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var TapMedia = MediaType("application/vnd.goa.context.sample.tap+json", func() {
	Description("A beer info of tap")
	Reference(tt.TapType)
	Attributes(func() {
		Attribute("id")
		Attribute("beerID")
		Attribute("beerTitle")
		Attribute("created_at")
		Attribute("updated_at")
		//		Attribute("id", Integer, "beer tap ID")
		//		Attribute("beerID", Integer, "beer name")
		//		Attribute("beerTitle", String, "beer name")
		//		Attribute("created_at", String, "create time stamp")
		//		Attribute("updated_at", String, "update time stamp")
		Required("id", "beerTitle")
	})
	View("default", func() {
		Attribute("id")
		Attribute("beerID")
		Attribute("beerTitle")
		Attribute("created_at")
		Attribute("updated_at")
	})
})
