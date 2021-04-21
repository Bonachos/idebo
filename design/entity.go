package design

import . "goa.design/goa/v3/dsl"

var _ = Service("entity", func() {
	Description("The entity service makes it possible to view, add or remove entities.")

	HTTP(func() {
		Path("/bo/entity")
	})

	Method("list", func() {
		Description("List all stored entities")
		Payload(func() {
			Attribute("authentication", String, "Authentication header")
		})
		Result(CollectionOf(EntityResult))
		HTTP(func() {
			GET("/")
			Response(StatusOK)
			Header("authentication:Authorization") // Authorization header
		})
	})

	Method("show", func() {
		Description("Show entity by ID")
		Payload(func() {
			Field(1, "id", String, "ID of entity to show")
			Field(2, "view", String, "Entity to render", func() {
				Enum("default")
			})
			Required("id")
		})
		Result(EntityResult)
		Error("not_found", NotFound, "Entity not found")
		HTTP(func() {
			GET("/{id}")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("showbyfield", func() {
		Description("Show entity by field")
		Payload(func() {
			Field(1, "fieldname", String, "field name to filter by")
			Field(2, "fieldvalue", String, "field value to filter by")
			Field(3, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("fieldname", "fieldvalue")
		})
		Result(EntityResult)
		Error("not_found", NotFound, "Entity not found")
		HTTP(func() {
			GET("/{fieldname}/{fieldvalue}")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("add", func() {
		Description("Add new entity and return its ID.")
		Payload(func() {
			Field(1, "entity", Entity, "Entity to add")
			Attribute("authentication", String, "Authentication header")
		})
		Result(String)
		HTTP(func() {
			POST("/")
			Header("authentication:Authorization") // Authorization header
			Response(StatusCreated)
		})
	})

	Method("update", func() {
		Description("Update the entity.")
		Payload(func() {
			Field(1, "id", String, "ID of entity to update")
			Field(2, "entity", Entity, "Entity to update")
			Required("id")
		})
		HTTP(func() {
			PUT("/{id}")
			Response(StatusOK)
		})
	})

	Method("remove", func() {
		Description("Remove entity")
		Payload(func() {
			Field(1, "id", String, "ID of entity to remove")
			Required("id")
		})
		Error("not_found", NotFound, "Entity not found")
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})
	})
})

// Entity describes an Entity or department of an organization.
var Entity = Type("Entity", func() {
	Description("Entity describes an Entity or department of an organization.")
	Attribute("name", String, "Name of entity", func() {
		MaxLength(255)
		Example("Departamento do Governo dos Açores")
		Meta("rpc:tag", "1")
	})
	Attribute("folder", String, "Folder name of entity. Serves as entity identifier.", func() {
		MaxLength(255)
		Example("ABCDE")
		Meta("rpc:tag", "2")
	})
	Attribute("inactive", Boolean, "Inactive entity (in maintenance)", func() {
		Example(false)
		Meta("rpc:tag", "3")
	})
	Required("name", "folder")
})

// EntityResult describes a entity retrieved by the entity service.
var EntityResult = ResultType("application/vnd.idea.entity", func() {
	Description("EntityResult describes an entity retrieved by the entity service.")
	Reference(Entity)
	TypeName("EntityResult")

	Attributes(func() {
		Attribute("id", UInt, "ID is the unique id of the entity.", func() {
			Example(123)
			Meta("rpc:tag", "6")
		})
		Field(2, "name")
		Field(3, "folder")
		Field(4, "inactive")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("folder")
		Attribute("inactive")
	})

	Required("id", "name", "folder")
})
