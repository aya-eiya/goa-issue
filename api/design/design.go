package design

import (
	. "goa.design/goa/dsl"
)

var _ = API("calc", func() {
	Title("Calculator Service")
	Description("calc hosts the Calculator Service.")
})

var CalcResultType = ResultType("application/vnd.result", func() {
	TypeName("Result")
	Attributes(func() {
		Field(1, "result", String)
	})
})

var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers")

	Method("add", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})

		Result(CalcResultType)

		GRPC(func() {
			Response(CodeOK)
		})
	})

})
