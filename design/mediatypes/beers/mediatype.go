package beers

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var BeerMedia = MediaType("application/vnd.goa.context.sample.beer+json", func() {
	Description("A bottle of beer")
	Attributes(func() {
		//Attribute("id", Integer, "beer ID")
		Attribute("id", Integer, func() {
			Description("id (type int64)")
			Metadata("struct:field:type", "int64")
			Example(1)
		})
		Attribute("title", String, "beer name")
		Attribute("type", String, "beer type")
		Attribute("created_at", String, "create time stamp")
		Attribute("updated_at", String, "update time stamp")
		Required("id", "title", "type")
	})
	View("default", func() {
		Attribute("id")
		Attribute("title")
		Attribute("type")
		Attribute("created_at")
		Attribute("updated_at")
	})
})
