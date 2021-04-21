package design

import . "goa.design/goa/v3/dsl"

var _ = Service("viewService", func() {
	Description("The viewService service makes it possible to view, add or remove viewServices.")

	HTTP(func() {
		Path("/bo/viewService")
	})

	Method("list", func() {
		Description("List all stored viewServices")
		Result(CollectionOf(ViewServiceResult), func() {
			View("tiny")
		})
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})

	Method("show", func() {
		Description("Show viewService by ID")
		Payload(func() {
			Field(1, "id", String, "ID of viewService to show")
			Field(2, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("id")
		})
		Result(ViewServiceResult)
		Error("not_found", NotFound, "ViewService not found")
		HTTP(func() {
			GET("/{id}")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("add", func() {
		Description("Add new viewService and return its ID.")
		Payload(ViewService)
		Result(String)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})

	Method("remove", func() {
		Description("Remove viewService")
		Payload(func() {
			Field(1, "id", String, "ID of viewService to remove")
			Required("id")
		})
		Error("not_found", NotFound, "ViewService not found")
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})
	})
})

// ViewService describes a View Service.
var ViewService = Type("ViewService", func() {
	Description("ViewService describes a View Service.")
	Attribute("name", String, "Name of viewService", func() {
		MaxLength(255)
		Example("Serviço WMS dos Açores")
		Meta("rpc:tag", "1")
	})
	Attribute("url", String, "Address (URL) of viewService", func() {
		MaxLength(255)
		Example("http://visualizador.sig.azores.gov.pt/sigazores")
		Meta("rpc:tag", "3")
	})
	Required("name", "url")
})

// ViewServiceResult describes a viewService retrieved by the viewService service.
var ViewServiceResult = ResultType("application/vnd.idea.viewService", func() {
	Description("ViewServiceResult describes a viewService retrieved by the viewService service.")
	Reference(ViewService)
	TypeName("ViewServiceResult")

	Attributes(func() {
		Attribute("id", String, "ID is the unique id of the viewService.", func() {
			Example("123abc")
			Meta("rpc:tag", "6")
		})
		Field(2, "name")
		Field(4, "url")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("url")
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("name")
		Attribute("url", func() {
			View("tiny")
		})
	})

	Required("id", "name", "url")
})
