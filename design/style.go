package design

import . "goa.design/goa/v3/dsl"

var _ = Service("style", func() {
	Description("The style service makes it possible to view, add or remove styles.")

	HTTP(func() {
		Path("/bo/style")
	})

	Method("list", func() {
		Description("List all stored styles")
		Result(CollectionOf(StyleResult), func() {
			View("tiny")
		})
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})

	Method("show", func() {
		Description("Show style by ID")
		Payload(func() {
			Field(1, "id", String, "ID of style to show")
			Field(2, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("id")
		})
		Result(StyleResult)
		Error("not_found", NotFound, "Style not found")
		HTTP(func() {
			GET("/{id}")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("add", func() {
		Description("Add new style and return its ID.")
		Payload(Style)
		Result(String)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})

	Method("remove", func() {
		Description("Remove style")
		Payload(func() {
			Field(1, "id", String, "ID of style to remove")
			Required("id")
		})
		Error("not_found", NotFound, "Style not found")
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})
	})
})

// Style describes a Style.
var Style = Type("Style", func() {
	Description("Style describes a Style.")
	Attribute("name", String, "Name of style", func() {
		MaxLength(255)
		Example("Estilo SLD")
		Meta("rpc:tag", "1")
	})
	Required("name")
})

// StyleResult describes a style retrieved by the style service.
var StyleResult = ResultType("application/vnd.idea.style", func() {
	Description("StyleResult describes a style retrieved by the style service.")
	Reference(Style)
	TypeName("StyleResult")

	Attributes(func() {
		Attribute("id", String, "ID is the unique id of the style.", func() {
			Example("123abc")
			Meta("rpc:tag", "6")
		})
		Field(2, "name")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("name")
	})

	Required("id", "name")
})
