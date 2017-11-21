package beers

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var BeerPayload = Type("BeerPayload", func() {
	Member("title", String, "Beer name")
	Member("type", String, "beer type")
	Member("homepage", String, func() {
		Description("beer official homepage")
		Format("uri")
	})
	Required("title", "type")
})
