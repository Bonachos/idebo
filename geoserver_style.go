package idea

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	geoserver "jpmenezes.com/idebo/gen/geoserver"
	geoserverstyle "jpmenezes.com/idebo/gen/geoserver_style"
)

// geoserverStyle service example implementation.
// The example methods log the requests and return zero values.
type geoserverStylesrvc struct {
	logger *log.Logger
}

// NewGeoserverStyle returns the geoserverStyle service implementation.
func NewGeoserverStyle(logger *log.Logger) geoserverstyle.Service {
	return &geoserverStylesrvc{logger}
}

// List all stored styles in a specific geoserver
func (s *geoserverStylesrvc) List(ctx context.Context, p *geoserverstyle.ListPayload) (res geoserverstyle.StyleResultCollection, err error) {
	s.logger.Print("geoserverStyle.list")
	// gsURL := "https://sigmar.dram.azores.gov.pt/geoserver/"
	gsURL := p.GeoserverURL
	var gsusername string
	var gspassword string
	gsproxyurl := "http://10.0.0.139:1080"
	// folderName := "car.biodiversidade_conservacao_marinha"

	gsURL = strings.TrimRight(gsURL, "/")
	if strings.LastIndex(gsURL, "/wms") == len(gsURL)-4 {
		gsURL = gsURL[:len(gsURL)-4]
	}
	u, err := url.Parse(gsURL)
	if err != nil {
		log.Println(err)
	}
	u.Path = path.Join(u.Path, "rest/styles")
	stylesURL := u.String()

	var authToken = jwtSecretToken
	listPayload := &geoserver.ListPayload{GeoserverURL: &p.GeoserverURL, Authentication: &authToken}

	geoserverService := NewGeoserver(log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile))
	resList, _, err := geoserverService.List(nil, listPayload)

	for _, geoserver := range resList {
		if *geoserver.Entityname == p.Entityid {
			if geoserver.Username == nil {
				return res, nil
			}

			gsusername = *geoserver.Username
			gspassword = *geoserver.Password

			if gsusername == "" {
				return res, nil
			}

			//creating the proxyURL
			client := &http.Client{}
			// proxyURL, err := url.Parse(gsproxyurl)
			if gsproxyurl != "" && err == nil {
				//adding the proxy settings to the Transport object
				client = &http.Client{Transport: &http.Transport{
					// Proxy: http.ProxyURL(proxyURL),
				}}
			}

			req, err := http.NewRequest("GET", stylesURL, nil)
			if err != nil {
				log.Println(err)
			}
			if gsusername != "" {
				req.SetBasicAuth(gsusername, gspassword)
			}
			req.Header.Add("Content-Type", "application/json")
			resp, err := client.Do(req)
			if err != nil {
				log.Println(err)
			}
			respBody, _ := ioutil.ReadAll(resp.Body)
			// respBodyString := string(respBody)

			// fmt.Println(respBodyString)

			var Styles *StylesResponse
			err = json.Unmarshal(respBody, &Styles)
			if err != nil {
				log.Println(err)
				return res, nil
			}
			// fmt.Println(Styles.Styles.Style[0].Name)

			for _, currentStyle := range Styles.Styles.Style {
				res = append(res, &geoserverstyle.StyleResult{
					ID:   currentStyle.Name,
					Name: currentStyle.Name,
				})
			}
		}
	}

	return res, nil
}

type StylesResponse struct {
	Styles StylesList `json:"styles"`
}

type StylesList struct {
	Style []Style `json:"style"`
}

type Style struct {
	Name string `json:"name"`
	Href string `json:"href,omitempty"`
}

type LanguageVersion struct {
	Version string `json:"version"`
}

type Legend struct {
	Width          string `json:"width,omitempty"`
	Height         string `json:"height,omitempty"`
	Format         string `json:"format"`
	OnlineResource string `json:"onlineResource"`
}
