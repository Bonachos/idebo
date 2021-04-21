package design

import . "goa.design/goa/v3/dsl"

var _ = Service("transformationService", func() {
	Description("The transformationService service makes it possible to view, add or remove transformationServices.")

	HTTP(func() {
		Path("/bo/transformationService")
	})

	Method("list", func() {
		Description("List all stored transformationServices")
		Result(CollectionOf(TransformationServiceResult), func() {
			View("tiny")
		})
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})

	Method("show", func() {
		Description("Show transformationService by ID")
		Payload(func() {
			Field(1, "id", String, "ID of transformationService to show")
			Field(2, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("id")
		})
		Result(TransformationServiceResult)
		Error("not_found", NotFound, "TransformationService not found")
		HTTP(func() {
			GET("/{id}")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("add", func() {
		Description("Add new transformationService and return its ID.")
		Payload(TransformationService)
		Result(String)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})

	Method("remove", func() {
		Description("Remove transformationService")
		Payload(func() {
			Field(1, "id", String, "ID of transformationService to remove")
			Required("id")
		})
		Error("not_found", NotFound, "TransformationService not found")
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})
	})
})

// TransformationService describes a Transformation Service.
var TransformationService = Type("TransformationService", func() {
	Description("TransformationService describes a Transformation Service.")
	Attribute("name", String, "Name of transformationService", func() {
		MaxLength(255)
		Example("Serviço WPS dos Açores")
		Meta("rpc:tag", "1")
	})
	Attribute("url", String, "Address (URL) of transformationService", func() {
		MaxLength(255)
		Example("http://visualizador.sig.azores.gov.pt/sigazores")
		Meta("rpc:tag", "3")
	})
	Required("name", "url")
})

// TransformationServiceResult describes a transformationService retrieved by the transformationService service.
var TransformationServiceResult = ResultType("application/vnd.idea.transformationService", func() {
	Description("TransformationServiceResult describes a transformationService retrieved by the transformationService service.")
	Reference(TransformationService)
	TypeName("TransformationServiceResult")

	Attributes(func() {
		Attribute("id", String, "ID is the unique id of the transformationService.", func() {
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
