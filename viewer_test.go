package idea

import (
	"log"
	"os"
	"testing"

	viewer "jpmenezes.com/idebo/gen/viewer"

	. "github.com/smartystreets/goconvey/convey"
)

func TestViewer(t *testing.T) {
	Convey("package viewer", t, func() {

		Convey("Client", func() {
			dropDB()

			viewerService := NewViewer(log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile))

			resList, _, err := viewerService.List(nil, &viewer.ListPayload{})
			So(err, ShouldBeError)
			So(resList, ShouldBeNil)

			createDB()

			resList, _, err = viewerService.List(nil, &viewer.ListPayload{})
			So(err, ShouldBeNil)
			So(len(resList), ShouldBeZeroValue)

			name := "name"
			title := "title"
			folder := "folder"
			centerX := "-28.193665"
			centerY := "38.676933"
			centerCRS := "EPSG:4326"
			newViewer := &viewer.Viewer{
				Name:      name,
				Title:     &title,
				Folder:    folder,
				Centerx:   &centerX,
				Centery:   &centerY,
				Centercrs: &centerCRS,
			}
			resAdd, err := viewerService.Add(nil, &viewer.AddPayload{Viewer: newViewer})
			So(err, ShouldBeNil)
			So(resAdd, ShouldEqual, "1")

			resList, _, err = viewerService.List(nil, &viewer.ListPayload{})
			So(err, ShouldBeNil)
			So(len(resList), ShouldEqual, 1)
			So(resList[0].ID, ShouldEqual, 1)
			So(resList[0].Name, ShouldEqual, name)
			So(*resList[0].Title, ShouldEqual, title)
			So(resList[0].Folder, ShouldEqual, folder)

			showPayload := &viewer.ShowPayload{ID: "1"}
			resShow, view, err := viewerService.Show(nil, showPayload)
			So(err, ShouldBeNil)
			So(view, ShouldEqual, "tiny")
			So(resShow.Name, ShouldEqual, "name")
			So(*resShow.Title, ShouldEqual, "title")
			So(resShow.Folder, ShouldEqual, "folder")

			showbyfieldPayload := &viewer.ShowbyfieldPayload{Fieldname: "Folder", Fieldvalue: "folder"}
			resShow, view, err = viewerService.Showbyfield(nil, showbyfieldPayload)
			So(err, ShouldBeNil)
			So(view, ShouldEqual, "tiny")
			So(resShow.Name, ShouldEqual, "name")
			So(*resShow.Title, ShouldEqual, "title")
			So(resShow.Folder, ShouldEqual, "folder")

			removePayload := &viewer.RemovePayload{ID: "1"}
			err = viewerService.Remove(nil, removePayload)
			So(err, ShouldBeNil)

			resList, _, err = viewerService.List(nil, &viewer.ListPayload{})
			So(err, ShouldBeNil)
			So(len(resList), ShouldBeZeroValue)

			dropDB()
		})
	})
}
