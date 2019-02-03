package design

import (
	. "goa.design/goa/dsl"
)

var _ = API("calc", func() {
	Title("Calculator Service")
	Description("calc hosts the Calculator Service.")
})

var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers")

	Method("add", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})

		Result(String)

		GRPC(func() {
			Response(CodeOK)
		})
	})

})
