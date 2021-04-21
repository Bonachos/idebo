package design

import . "goa.design/goa/v3/dsl"

var _ = Service("geoserverData", func() {
	Description("The geoserverData service makes it possible to view, add or remove data from geoservers.")

	HTTP(func() {
		Path("/bo/geoserverData")
	})

	Method("layers", func() {
		Description("List all layers from a geoserver")
		Payload(func() {
			Field(1, "geoserverid", String, "id of geoserver")
			Field(2, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("geoserverid")
		})
		Result(CollectionOf(GeoserverLayerResult), func() {
			View("default")
		})
		Error("not_found", NotFound, "Geoserver not found")
		HTTP(func() {
			GET("/{geoserverid}")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})
})

// GeoserverLayerResult describes a geoserver layer retrieved by the geoserverData service.
var GeoserverLayerResult = ResultType("application/vnd.idea.geoserverData", func() {
	Description("GeoserverLayerResult describes a geoserver layer retrieved by the geoserverData service.")
	Reference(Geoserver)
	TypeName("GeoserverLayerResult")

	Attributes(func() {
		Field(1, "title")
		Field(2, "name")
		Field(3, "srs")
	})

	View("default", func() {
		Attribute("title")
		Attribute("name")
		Attribute("srs")
	})

	View("tiny", func() {
		Attribute("name")
	})

	Required("title", "name")
})
