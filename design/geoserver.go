package design

import . "goa.design/goa/v3/dsl"

var _ = Service("geoserver", func() {
	Description("The geoserver service makes it possible to view, add or remove geoservers.")

	HTTP(func() {
		Path("/bo/geoserver")
	})

	Error("unauthorized", String, "Credentials are invalid")

	Method("list", func() {
		Description("List all stored geoservers")
		Payload(func() {
			Field(1, "geoserverURL", String, "URL of geoserver")
			Attribute("authentication", String, "Authentication header")
		})
		Result(CollectionOf(GeoserverResult))
		HTTP(func() {
			GET("/")
			Param("geoserverURL")
			Header("authentication:Authorization") // Authorization header
			Response(StatusOK)
		})
	})

	Method("show", func() {
		Description("Show geoserver by ID")
		Payload(func() {
			Field(1, "id", String, "ID of geoserver to show")
			Field(2, "view", String, "Geoserver to render", func() {
				Enum("default", "tiny")
			})
			Required("id")
		})
		Result(GeoserverResult)
		Error("not_found", NotFound, "Geoserver not found")
		HTTP(func() {
			GET("/{id}")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("add", func() {
		Description("Add new geoserver and return its ID.")
		Payload(func() {
			Field(1, "geoserver", Geoserver, "Geoserver to add")
			Attribute("authentication", String, "Authentication header")
		})
		Result(String)
		HTTP(func() {
			POST("/")
			Header("authentication:Authorization") // Authorization header
			Response(StatusCreated)
			Response("unauthorized", StatusUnauthorized)
		})
	})

	Method("update", func() {
		Description("Update the geoserver.")
		Payload(func() {
			Field(1, "id", String, "ID of geoserver to update")
			Field(2, "geoserver", Geoserver, "Geoserver to update")
			Required("id")
		})
		HTTP(func() {
			PUT("/{id}")
			Response(StatusOK)
		})
	})

	Method("remove", func() {
		Description("Remove geoserver")
		Payload(func() {
			Field(1, "id", String, "ID of geoserver to remove")
			Required("id")
		})
		Error("not_found", NotFound, "Geoserver not found")
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})
	})
})

// Geoserver describes a GIS Geoserver.
var Geoserver = Type("Geoserver", func() {
	Description("Geoserver describes a GIS Geoserver.")
	Attribute("name", String, "Name of geoserver", func() {
		MaxLength(255)
		Example("Geo Server dos Açores")
		Meta("rpc:tag", "1")
	})
	Attribute("url", String, "Address (URL) of geoserver", func() {
		MaxLength(255)
		Example("http://visualizador.sig.azores.gov.pt/geoserver")
		Meta("rpc:tag", "2")
	})
	Attribute("username", String, "Username of an administrator of geoserver", func() {
		MaxLength(255)
		Example("username123")
		Meta("rpc:tag", "3")
	})
	Attribute("password", String, "Password of an administrator of geoserver", func() {
		MaxLength(255)
		Example("aStr0ngP@ssworD!")
		Meta("rpc:tag", "4")
	})
	Attribute("entity", String, "Entity to which the viewer belongs", func() {
		MaxLength(64)
		Example("Department of the Government")
		Meta("rpc:tag", "14")
	})
	Required("name", "url", "entity")
})

// GeoserverResult describes a geoserver retrieved by the geoserver service.
var GeoserverResult = ResultType("application/vnd.idea.geoserver", func() {
	Description("GeoserverResult describes a geoserver retrieved by the geoserver service.")
	Reference(Geoserver)
	TypeName("GeoserverResult")

	Attributes(func() {
		Attribute("id", UInt, "ID is the unique id of the geoserver.", func() {
			Example(123)
			Meta("rpc:tag", "6")
		})
		Field(2, "name")
		Field(3, "url")
		Field(4, "username")
		Field(4, "password")
		Field(5, "entity")
		Field(6, "entityname")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("url")
		Attribute("username")
		Attribute("password")
		Attribute("entity")
		Attribute("entityname")
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("name")
		Attribute("url", func() {
			View("tiny")
		})
		Attribute("entity")
		Attribute("entityname")
	})

	Required("id", "name", "url")
})
