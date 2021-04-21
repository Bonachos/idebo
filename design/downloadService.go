package design

import . "goa.design/goa/v3/dsl"

var _ = Service("downloadService", func() {
	Description("The downloadService service makes it possible to view, add or remove downloadServices.")

	HTTP(func() {
		Path("/bo/downloadService")
	})

	Method("list", func() {
		Description("List all stored downloadServices")
		Result(CollectionOf(DownloadServiceResult), func() {
			View("tiny")
		})
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})

	Method("show", func() {
		Description("Show downloadService by ID")
		Payload(func() {
			Field(1, "id", String, "ID of downloadService to show")
			Field(2, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("id")
		})
		Result(DownloadServiceResult)
		Error("not_found", NotFound, "DownloadService not found")
		HTTP(func() {
			GET("/{id}")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("add", func() {
		Description("Add new downloadService and return its ID.")
		Payload(DownloadService)
		Result(String)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})

	Method("remove", func() {
		Description("Remove downloadService")
		Payload(func() {
			Field(1, "id", String, "ID of downloadService to remove")
			Required("id")
		})
		Error("not_found", NotFound, "DownloadService not found")
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})
	})
})

// DownloadService describes a Download Service.
var DownloadService = Type("DownloadService", func() {
	Description("DownloadService describes a Download Service.")
	Attribute("name", String, "Name of downloadService", func() {
		MaxLength(255)
		Example("Serviço WFS dos Açores")
		Meta("rpc:tag", "1")
	})
	Attribute("url", String, "Address (URL) of downloadService", func() {
		MaxLength(255)
		Example("http://visualizador.sig.azores.gov.pt/sigazores")
		Meta("rpc:tag", "3")
	})
	Required("name", "url")
})

// DownloadServiceResult describes a downloadService retrieved by the downloadService service.
var DownloadServiceResult = ResultType("application/vnd.idea.downloadService", func() {
	Description("DownloadServiceResult describes a downloadService retrieved by the downloadService service.")
	Reference(DownloadService)
	TypeName("DownloadServiceResult")

	Attributes(func() {
		Attribute("id", String, "ID is the unique id of the downloadService.", func() {
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
