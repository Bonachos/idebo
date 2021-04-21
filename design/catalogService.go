package design

import . "goa.design/goa/v3/dsl"

var _ = Service("catalogService", func() {
	Description("The catalogService service makes it possible to view, add or remove catalogServices.")

	HTTP(func() {
		Path("/bo/catalogService")
	})

	Method("list", func() {
		Description("List all stored catalogServices")
		Result(CollectionOf(CatalogServiceResult), func() {
			View("tiny")
		})
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})

	Method("show", func() {
		Description("Show catalogService by ID")
		Payload(func() {
			Field(1, "id", String, "ID of catalogService to show")
			Field(2, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("id")
		})
		Result(CatalogServiceResult)
		Error("not_found", NotFound, "CatalogService not found")
		HTTP(func() {
			GET("/{id}")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("add", func() {
		Description("Add new catalogService and return its ID.")
		Payload(CatalogService)
		Result(String)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})

	Method("remove", func() {
		Description("Remove catalogService")
		Payload(func() {
			Field(1, "id", String, "ID of catalogService to remove")
			Required("id")
		})
		Error("not_found", NotFound, "CatalogService not found")
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})
	})
})

// CatalogService describes a Catalog Service.
var CatalogService = Type("CatalogService", func() {
	Description("CatalogService describes a Catalog Service.")
	Attribute("name", String, "Name of catalogService", func() {
		MaxLength(255)
		Example("Serviço CSW dos Açores")
		Meta("rpc:tag", "1")
	})
	Attribute("url", String, "Address (URL) of catalogService", func() {
		MaxLength(255)
		Example("http://visualizador.sig.azores.gov.pt/sigazores")
		Meta("rpc:tag", "3")
	})
	Required("name", "url")
})

// CatalogServiceResult describes a catalogService retrieved by the catalogService service.
var CatalogServiceResult = ResultType("application/vnd.idea.catalogService", func() {
	Description("CatalogServiceResult describes a catalogService retrieved by the catalogService service.")
	Reference(CatalogService)
	TypeName("CatalogServiceResult")

	Attributes(func() {
		Attribute("id", String, "ID is the unique id of the catalogService.", func() {
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
