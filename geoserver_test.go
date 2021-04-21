package idea

import (
	"log"
	"os"
	"testing"

	geoserver "jpmenezes.com/idebo/gen/geoserver"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGeoserver(t *testing.T) {
	Convey("package geoserver", t, func() {

		Convey("Client", func() {
			dropDB()

			geoserverService := NewGeoserver(log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile))

			resList, view, err := geoserverService.List(nil, nil)
			So(err, ShouldBeError)
			So(resList, ShouldBeNil)

			createDB()

			resList, view, err = geoserverService.List(nil, nil)
			So(err, ShouldBeNil)
			So(len(resList), ShouldBeZeroValue)

			name := "name"
			url := "http://testurl.com"
			username := "testusername"
			password := "testpassword"
			newGeoserver := &geoserver.Geoserver{
				Name:     name,
				URL:      url,
				Username: &username,
				Password: &password,
			}
			resAdd, err := geoserverService.Add(nil, &geoserver.AddPayload{Geoserver: newGeoserver})
			So(err, ShouldBeNil)
			So(resAdd, ShouldEqual, "1")

			resList, view, err = geoserverService.List(nil, nil)
			So(err, ShouldBeNil)
			So(len(resList), ShouldEqual, 1)

			showPayload := &geoserver.ShowPayload{ID: "1"}
			resShow, view, err := geoserverService.Show(nil, showPayload)
			So(err, ShouldBeNil)
			So(view, ShouldEqual, "default")
			So(resShow.Name, ShouldEqual, name)
			So(resShow.URL, ShouldEqual, url)
			So(*resShow.Username, ShouldEqual, username)

			// showbyfieldPayload := &viewer.ShowbyfieldPayload{Fieldname: "Folder", Fieldvalue: "folder"}
			// resShow, view, err = viewerService.Showbyfield(nil, showbyfieldPayload)
			// So(err, ShouldBeNil)
			// So(view, ShouldEqual, "default")
			// So(resShow.Name, ShouldEqual, "name")
			// So(*resShow.Title, ShouldEqual, "title")
			// So(resShow.Folder, ShouldEqual, "folder")

			removePayload := &geoserver.RemovePayload{ID: "1"}
			err = geoserverService.Remove(nil, removePayload)
			So(err, ShouldBeNil)

			resList, view, err = geoserverService.List(nil, nil)
			So(err, ShouldBeNil)
			So(len(resList), ShouldBeZeroValue)

			dropDB()
		})
	})
}
