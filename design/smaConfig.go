package design

import . "goa.design/goa/v3/dsl"

var _ = Service("smaConfig", func() {
	Description("The smaConfig service makes it possible to update the SMA Configuration.")

	HTTP(func() {
		Path("/bo/smaConfig")
	})

	Method("update", func() {
		Description("Update the SMA Configuration")
		Payload(SMAConfig)
		HTTP(func() {
			PUT("/")
			Response(StatusOK)
		})
	})
})

// SMAConfig describes the SMA Configuration.
var SMAConfig = Type("SMAConfig", func() {
	Description("SMAConfig describes the SMA Configuration.")
	Attribute("title", String, "Title of the main SMA page.", func() {
		MaxLength(255)
		Example("Sistema de Metadados dos Açores")
		Meta("rpc:tag", "1")
	})
	Required("title")
})
