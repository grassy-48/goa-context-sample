package taps

import (
	dsb "goa-context-sample/design/dsls/beers"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var TapType = Type("TapType", func() {
	Attribute("id", Integer, "beer tap ID")
	Attribute("beerID", Integer, "beer name")
	Attribute("beerTitle", String, dsb.BeerTitleDSL)
	Attribute("created_at", String, "create time stamp")
	Attribute("updated_at", String, "update time stamp")
})
