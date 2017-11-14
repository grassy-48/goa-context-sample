package taps

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var TapPayload = Type("TypePayload", func() {
	Member("beerID", Integer, "beer ID related tap")
	Required("beerID")
})
