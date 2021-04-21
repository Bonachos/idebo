package idea

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	geoserver "jpmenezes.com/idebo/gen/geoserver"
	geoserverdata "jpmenezes.com/idebo/gen/geoserver_data"

	"jpmenezes.com/idebo/geoserverAPI"
)

// geoserverData service example implementation.
// The example methods log the requests and return zero values.
type geoserverDatasrvc struct {
	logger *log.Logger
}

// NewGeoserverData returns the geoserverData service implementation.
func NewGeoserverData(logger *log.Logger) geoserverdata.Service {
	return &geoserverDatasrvc{logger}
}

// List all layers from a geoserver
func (s *geoserverDatasrvc) Layers(ctx context.Context, p *geoserverdata.LayersPayload) (res geoserverdata.GeoserverLayerResultCollection, err error) {
	s.logger.Print("geoserverData.layers")

	var geoserverURL string
	if p.Geoserverid == "12345" {
		geoserverURL = "https://visualizador.idea.azores.gov.pt/contexts/mapproxy"
	} else {
		geoserverService := NewGeoserver(log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile))

		showPayload := &geoserver.ShowPayload{ID: p.Geoserverid}
		geoserver, _, _ := geoserverService.Show(nil, showPayload)
		geoserverURL = geoserver.URL
		fmt.Println(geoserverURL)
	}
	geoserverURL += "?service=wms&version=1.3.0&request=GetCapabilities"

	resp, err := http.Get(geoserverURL)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var getCapabilities geoserverAPI.WMS_Capabilities
	err = xml.Unmarshal(body, &getCapabilities)
	if err != nil {
		return
	}

	var result geoserverdata.GeoserverLayerResultCollection

	for _, layer := range getCapabilities.Capability.Layer.Layer {
		if layer.Name != nil {
			var srs string
			for _, crs := range layer.CRS {
				srs += "," + crs.Text
			}
			if srs == "" {
				for _, bbox := range layer.BoundingBox {
					srs += "," + bbox.AttrCRS
				}
			}
			srs = strings.TrimPrefix(srs, ",")
			geoserverLayerResult := &geoserverdata.GeoserverLayerResult{
				Title: layer.Title.Text,
				Name:  layer.Name.Text,
				Srs:   &srs,
			}
			result = append(result, geoserverLayerResult)
			fmt.Println(layer.Name.Text + " | " + layer.Title.Text)
		}
	}

	return result, nil
}
