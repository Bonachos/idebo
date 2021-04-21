package design

import . "goa.design/goa/v3/dsl"

var _ = Service("geodata", func() {
	Description("The geodata service makes it possible to add geodata.")

	HTTP(func() {
		Path("/bo/geodata")
	})

	Method("list", func() {
		Description("List all stored geodata")
		Payload(func() {
			Attribute("authentication", String, "Authentication header")
		})
		Result(CollectionOf(GeodataResult))
		HTTP(func() {
			GET("/")
			Header("authentication:Authorization") // Authorization header
			Response(StatusOK)
		})
	})

	Method("upload", func() {
		HTTP(func() {
			POST("/")
			MultipartRequest()
		})
		Payload(FilesUpload)
		Result(String)
	})

	Method("remove", func() {
		Description("Remove geodata")
		Payload(func() {
			Field(1, "id", String, "ID of geodata to remove")
			Required("id")
		})
		Error("not_found", NotFound, "Geodata not found")
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})
	})
})

// FileUpload is a single file upload element
var FileUpload = Type("FileUpload", func() {
	Description("A single File Upload type")
	Attribute("type", String)
	Attribute("bytes", Bytes)
	Attribute("name", String)
})

// FilesUpload is a list of files
var FilesUpload = Type("FilesUpload", func() {
	Description("Collection of uploaded files")

	Attribute("Files", ArrayOf(FileUpload), "Collection of uploaded files")
})

// Geodata describes a Geodata file.
var Geodata = Type("Geodata", func() {
	Description("Geodata describes a Geodata file.")
	Attribute("name", String, "Name of geodata file", func() {
		MaxLength(255)
		Example("nomeDoFicheiro.extensao")
		Meta("rpc:tag", "1")
	})
	Attribute("entity", String, "Entity to which the geodata file belongs", func() {
		MaxLength(64)
		Example("Department of the Government")
		Meta("rpc:tag", "14")
	})
	Required("name", "entity")
})

// GeodataResult describes a geodata file retrieved by the geodata service.
var GeodataResult = ResultType("application/vnd.idea.geodata", func() {
	Description("GeodataResult describes a geodata file retrieved by the geodata service.")
	Reference(Geodata)
	TypeName("GeodataResult")

	Attributes(func() {
		Attribute("id", UInt, "ID is the unique id of the geodata file.", func() {
			Example(123)
			Meta("rpc:tag", "6")
		})
		Field(2, "name")
		Field(5, "entity")
		Field(6, "entityname")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("entity")
		Attribute("entityname")
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("name")
		Attribute("entity")
		Attribute("entityname")
	})

	Required("id", "name")
})
