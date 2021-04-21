package design

import . "goa.design/goa/v3/dsl"

var _ = Service("geoserverStyle", func() {
	Description("The geoserverStyle service makes it possible to view, add or remove styles in a geoserver.")

	HTTP(func() {
		Path("/bo/geoserverstyle")
	})

	Method("list", func() {
		Description("List all stored styles in a specific geoserver")
		Payload(func() {
			Field(1, "geoserverURL", String, "URL of geoserver to check for styles in")
			Field(2, "entityid", String, "Entity ID")
			Field(3, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("geoserverURL", "entityid")
		})
		Result(CollectionOf(StyleResult), func() {
			View("tiny")
		})
		Error("not_found", NotFound, "Geoserver not found")
		HTTP(func() {
			GET("/")
			Param("geoserverURL")
			Param("entityid")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	// Method("show", func() {
	// 	Description("Show style by ID")
	// 	Payload(func() {
	// 		Field(1, "id", String, "ID of style to show")
	// 		Field(2, "view", String, "View to render", func() {
	// 			Enum("default", "tiny")
	// 		})
	// 		Required("id")
	// 	})
	// 	Result(StyleResult)
	// 	Error("not_found", NotFound, "Style not found")
	// 	HTTP(func() {
	// 		GET("/{id}")
	// 		Param("view")
	// 		Response(StatusOK)
	// 		Response("not_found", StatusNotFound)
	// 	})
	// })

	// Method("add", func() {
	// 	Description("Add new style and return its ID.")
	// 	Payload(Style)
	// 	Result(String)
	// 	HTTP(func() {
	// 		POST("/")
	// 		Response(StatusCreated)
	// 	})
	// })

	// Method("remove", func() {
	// 	Description("Remove style")
	// 	Payload(func() {
	// 		Field(1, "id", String, "ID of style to remove")
	// 		Required("id")
	// 	})
	// 	Error("not_found", NotFound, "Style not found")
	// 	HTTP(func() {
	// 		DELETE("/{id}")
	// 		Response(StatusNoContent)
	// 	})
	// })
})
