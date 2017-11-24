package beers

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var BeerTitleDSL = func() {
	Description("beer name")
	MaxLength(30)
	MinLength(0)
	Default("なにか美味しいビール")
	Example("銀河高原ビール")
}
