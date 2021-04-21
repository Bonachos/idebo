package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = Service("openapi", func() {
	// Serve the file with relative path ../../gen/http/openapi.json for
	// requests sent to /swagger.json.
	Files("/swagger.json", "../../gen/http/openapi.json")
})

// API describes the global properties of the API server.
var _ = API("idea", func() {
	Title("IDE.A Service")
	Description("HTTP service for managing IDE.A")
	cors.Origin("*", func() {
		cors.Headers("*")
		cors.Methods("GET", "POST", "PUT", "DELETE")
	})
	Server("bo", func() {
		Description("bo hosts the IDE.A services.")
		Services(
			"catalogService",
			"downloadService",
			"entity",
			"geodata",
			"geoserver",
			"geoserverData",
			"geoserverStyle",
			"openapi",
			"smaConfig",
			"style",
			"transformationService",
			"viewer",
			"viewService",
		)
		Host("localhost", func() {
			Description("default host")
			URI("http://0.0.0.0:8888/bo")
		})
	})
})

// NotFound is the type returned when attempting to show or delete a viewer that does not exist.
var NotFound = Type("NotFound", func() {
	Description("NotFound is the type returned when attempting to show or delete a viewer that does not exist.")
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("viewer 1 not found")
		Meta("rpc:tag", "1")
	})
	Field(2, "id", String, "ID of missing viewer")
	Required("message", "id")
})
