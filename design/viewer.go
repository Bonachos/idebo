package design

import . "goa.design/goa/v3/dsl"

var _ = Service("viewer", func() {
	Description("The viewer service makes it possible to view, add or remove GIS viewers.")

	HTTP(func() {
		Path("/bo/viewer")
	})

	Error("unauthorized", String, "Credentials are invalid")

	Method("list", func() {
		Description("List all stored viewers")
		Payload(func() {
			Field(1, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Attribute("authentication", String, "Authentication header")
		})
		Result(CollectionOf(ViewerResult))
		HTTP(func() {
			GET("/")
			Param("view")
			Header("authentication:Authorization") // Authorization header
			Response(StatusOK)
		})
	})

	Method("listlayers", func() {
		Description("List all layers of a viewer")
		Payload(func() {
			Field(1, "id", String, "ID of viewer")
			Required("id")
		})
		Result(CollectionOf(ViewerLayerResult), func() {
			View("full")
		})
		HTTP(func() {
			GET("/{id}/layers")
			Response(StatusOK)
		})
	})

	Method("listgroups", func() {
		Description("List all groups of a viewer")
		Payload(func() {
			Field(1, "id", String, "ID of viewer")
			Required("id")
		})
		Result(CollectionOf(ViewerGroupResult), func() {
			View("default")
		})
		HTTP(func() {
			GET("/{id}/groups")
			Response(StatusOK)
		})
	})

	Method("show", func() {
		Description("Show viewer by ID")
		Payload(func() {
			Field(1, "id", String, "ID of viewer to show")
			Field(2, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("id")
		})
		Result(ViewerResult)
		Error("not_found", NotFound, "Viewer not found")
		HTTP(func() {
			GET("/{id}")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("showbyfield", func() {
		Description("Show viewer by field")
		Payload(func() {
			Field(1, "fieldname", String, "field name to filter by")
			Field(2, "fieldvalue", String, "field value to filter by")
			Field(3, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("fieldname", "fieldvalue")
		})
		Result(ViewerResult)
		Error("not_found", NotFound, "Viewer not found")
		HTTP(func() {
			GET("/{fieldname}/{fieldvalue}")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("showbyfield2", func() {
		Description("Show viewer by 2 fields")
		Payload(func() {
			Field(1, "fieldname", String, "field name to filter by")
			Field(2, "fieldvalue", String, "field value to filter by")
			Field(3, "fieldname2", String, "field name 2 to filter by")
			Field(4, "fieldvalue2", String, "field value 2 to filter by")
			Field(5, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("fieldname", "fieldvalue", "fieldname2", "fieldvalue2")
		})
		Result(ViewerResult)
		Error("not_found", NotFound, "Viewer not found")
		HTTP(func() {
			GET("/{fieldname}/{fieldvalue}/{fieldname2}/{fieldvalue2}")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("add", func() {
		Description("Add new viewer and return its ID.")
		Payload(Viewer)
		Payload(func() {
			Field(1, "viewer", Viewer, "Viewer to add")
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
		Description("Update the viewer")
		Payload(func() {
			Field(1, "id", String, "ID of viewer to update")
			Field(2, "viewer", Viewer, "Viewer to update")
			Required("id")
		})
		HTTP(func() {
			PUT("/{id}")
			Response(StatusOK)
		})
	})

	Method("updatelayers", func() {
		Description("Update the layers of the viewer")
		Payload(func() {
			Field(1, "id", String, "ID of viewer to update")
			Field(2, "viewerlayers", CollectionOf(ViewerLayerResult), "Layers of viewer to update")
			Required("id")
		})
		HTTP(func() {
			PUT("/{id}/layers")
			Response(StatusOK)
		})
	})

	Method("updategroups", func() {
		Description("Update the groups of the viewer")
		Payload(func() {
			Field(1, "id", String, "ID of viewer to update")
			Field(2, "viewergroups", CollectionOf(ViewerGroupResult), "Groups of viewer to update")
			Required("id")
		})
		HTTP(func() {
			PUT("/{id}/groups")
			Response(StatusOK)
		})
	})

	Method("remove", func() {
		Description("Remove viewer")
		Payload(func() {
			Field(1, "id", String, "ID of viewer to remove")
			Required("id")
		})
		Error("not_found", NotFound, "Viewer not found")
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})
	})
})

// Viewer describes a GIS Viewer.
var Viewer = Type("Viewer", func() {
	Description("Viewer describes a GIS Viewer.")
	Attribute("name", String, "Name of viewer", func() {
		MaxLength(255)
		Example("SIG dos Açores")
		Meta("rpc:tag", "1")
	})
	Attribute("title", String, "Title of viewer", func() {
		MaxLength(255)
		Example("Sistema de Informação Geográfica dos Açores")
		Meta("rpc:tag", "2")
	})
	Attribute("url", String, "Address (URL) of viewer", func() {
		MaxLength(255)
		Example("http://visualizador.sig.azores.gov.pt/sigazores")
		Meta("rpc:tag", "3")
	})
	Attribute("image", String, "Image (thumbnail) URL of viewer", func() {
		MaxLength(255)
		Example("http://visualizador.sig.azores.gov.pt/images/sigazores.png")
		Meta("rpc:tag", "4")
	})
	Attribute("folder", String, "Folder that will define the URL of viewer", func() {
		MaxLength(255)
		Example("visualizador1")
		Meta("rpc:tag", "5")
	})
	Attribute("hostname", String, "Hostname that can also be used to reference the viewer", func() {
		MaxLength(255)
		Example("visualizador1.azores.gov.pt")
		Meta("rpc:tag", "6")
	})
	Attribute("centerx", String, "X coordinate of the center of the viewer", func() {
		MaxLength(64)
		Example("-28.193665")
		Meta("rpc:tag", "7")
	})
	Attribute("centery", String, "Y coordinate of the center of the viewer", func() {
		MaxLength(64)
		Example("38.676933")
		Meta("rpc:tag", "8")
	})
	Attribute("centercrs", String, "CRS of the center of the viewer", func() {
		MaxLength(64)
		Example("EPSG:4326")
		Meta("rpc:tag", "9")
	})
	Attribute("anonymous", Boolean, "Public viewer (anonymous access)", func() {
		Example(true)
		Meta("rpc:tag", "10")
	})
	Attribute("inactive", Boolean, "Inactive viewer (in maintenance)", func() {
		Example(false)
		Meta("rpc:tag", "11")
	})
	Attribute("projection", String, "CRS of the viewer", func() {
		MaxLength(64)
		Example("EPSG:4326")
		Meta("rpc:tag", "12")
	})
	Attribute("zoom", String, "Initial zoom of the viewer", func() {
		MaxLength(64)
		Example("1")
		Meta("rpc:tag", "13")
	})
	Attribute("entity", String, "Entity to which the viewer belongs", func() {
		MaxLength(64)
		Example("Department of the Government")
		Meta("rpc:tag", "14")
	})
	Attribute("emailaddress", String, "Email Address to which notifications will be sent", func() {
		MaxLength(64)
		Example("department@azores.gov.pt")
		Meta("rpc:tag", "15")
	})
	Attribute("xminextent", String, "Extent: Min X", func() {
		MaxLength(64)
		Example("-31.12972061137288")
		Meta("rpc:tag", "16")
	})
	Attribute("xmaxextent", String, "Extent: Max X", func() {
		MaxLength(64)
		Example("-31.081590117941197")
		Meta("rpc:tag", "17")
	})
	Attribute("yminextent", String, "Extent: Min Y", func() {
		MaxLength(64)
		Example("39.66872113303264")
		Meta("rpc:tag", "18")
	})
	Attribute("ymaxextent", String, "Extent: Max Y", func() {
		MaxLength(64)
		Example("39.726511657336125")
		Meta("rpc:tag", "19")
	})
	Attribute("mapheader", String, "URL to Map Header Image", func() {
		MaxLength(255)
		Example("http://visualizador.sig.azores.gov.pt/images/sigazores.png")
		Meta("rpc:tag", "20")
	})
	Attribute("mapnorth", String, "URL to Map North SVG Image", func() {
		MaxLength(255)
		Example("http://visualizador.sig.azores.gov.pt/images/sigazores.svg")
		Meta("rpc:tag", "21")
	})
	Required("name", "folder", "entity")
})

// ViewerResult describes a viewer retrieved by the viewer service.
var ViewerResult = ResultType("application/vnd.idea.viewer", func() {
	Description("ViewerResult describes a viewer retrieved by the viewer service.")
	Reference(Viewer)
	TypeName("ViewerResult")

	Attributes(func() {
		Attribute("id", UInt, "ID is the unique id of the viewer.", func() {
			Example(123)
			Meta("rpc:tag", "6")
		})
		Field(2, "name")
		Field(3, "title")
		Field(4, "url")
		Field(5, "image")
		Field(6, "folder")
		Field(7, "hostname")
		Field(8, "centerx")
		Field(9, "centery")
		Field(10, "centercrs")
		Field(11, "anonymous")
		Field(12, "inactive")
		Field(13, "projection")
		Field(14, "zoom")
		Field(15, "entity")
		Field(16, "entityname")
		Field(17, "emailaddress")
		Field(18, "xminextent")
		Field(19, "xmaxextent")
		Field(20, "yminextent")
		Field(21, "ymaxextent")
		Field(22, "mapheader")
		Field(23, "mapnorth")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("title")
		Attribute("url")
		Attribute("image")
		Attribute("folder")
		Attribute("hostname")
		Attribute("centerx")
		Attribute("centery")
		Attribute("centercrs")
		Attribute("anonymous")
		Attribute("inactive")
		Attribute("projection")
		Attribute("zoom")
		Attribute("entity")
		Attribute("entityname")
		Attribute("emailaddress")
		Attribute("xminextent")
		Attribute("xmaxextent")
		Attribute("yminextent")
		Attribute("ymaxextent")
		Attribute("mapheader")
		Attribute("mapnorth")
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("name")
		Attribute("url")
		Attribute("folder")
		Attribute("entity")
		Attribute("entityname")
		Attribute("xminextent")
		Attribute("xmaxextent")
		Attribute("yminextent")
		Attribute("ymaxextent")
	})

	Required("id", "name", "folder", "entity")
})

// ViewerLayerResult describes a layer of a viewer, retrieved by the viewer service.
var ViewerLayerResult = ResultType("application/vnd.idea.viewerlayer", func() {
	Description("ViewerLayerResult describes a layer of a viewer, retrieved by the viewer service.")
	Reference(Viewer)
	TypeName("ViewerLayerResult")

	Attributes(func() {
		Attribute("id", UInt, "ID is the unique id of the layer.", func() {
			Example(123)
			Meta("rpc:tag", "6")
		})
		Field(2, "type")
		Field(3, "title")
		Field(4, "name")
		Field(5, "source")
		Field(6, "group")
		Field(7, "visibility")
		Field(8, "url")
		Field(9, "format")
		Field(10, "style")
		Field(11, "catalogurl")
		Field(12, "mapheader")
		Field(13, "mapnorth")
	})

	View("default", func() {
		Attribute("id")
		Attribute("format")
		Attribute("group")
		Attribute("name")
		Attribute("source")
		Attribute("style")
		Attribute("title")
		Attribute("type")
		Attribute("visibility")
		Attribute("catalogurl")
		Attribute("mapheader")
		Attribute("mapnorth")
	})

	View("full", func() {
		Attribute("id")
		Attribute("format")
		Attribute("group")
		Attribute("name")
		Attribute("source")
		Attribute("style")
		Attribute("title")
		Attribute("type")
		Attribute("url")
		Attribute("visibility")
		Attribute("catalogurl")
		Attribute("mapheader")
		Attribute("mapnorth")
	})

	Required("id", "type", "title", "name", "group")
})

// ViewerGroupResult describes a group of a viewer, retrieved by the viewer service.
var ViewerGroupResult = ResultType("application/vnd.idea.viewergroup", func() {
	Description("ViewerGroupResult describes a group of a viewer, retrieved by the viewer service.")
	Reference(Viewer)
	TypeName("ViewerGroupResult")

	Attributes(func() {
		Attribute("id", String, "ID is the unique id of the group.", func() {
			Example("123")
			Meta("rpc:tag", "6")
		})
		Field(2, "title")
	})

	View("default", func() {
		Attribute("id")
		Attribute("title")
	})

	Required("id", "title")
})
