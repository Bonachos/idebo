package idea

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	entity "jpmenezes.com/idebo/gen/entity"
	viewer "jpmenezes.com/idebo/gen/viewer"

	// Postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// contextsFolder = "/app/"
	// contextsFolder = "/home/jpmen/MapStore2/web/client/"
	contextsFolder = os.Getenv("CONTEXTSFOLDER")
	projectName    = os.Getenv("PROJECTNAME")

	errUnauthorized viewer.Unauthorized
)

// viewer service example implementation.
// The example methods log the requests and return zero values.
type viewersrvc struct {
	logger *log.Logger
}

// NewViewer returns the viewer service implementation.
func NewViewer(logger *log.Logger) viewer.Service {
	return &viewersrvc{logger}
}

// List all stored viewers
func (s *viewersrvc) List(ctx context.Context, p *viewer.ListPayload) (res viewer.ViewerResultCollection, view string, err error) {
	methodName := "viewer.list"
	s.logger.Print(methodName)

	if p.View != nil {
		view = *p.View
	} else {
		view = "tiny"
	}

	if p == nil || p.Authentication == nil {
		return nil, view, nil
	}

	userDetails := GetUserAuthInfo(*p.Authentication)
	// DEBUG fmt.Println(userDetails)

	db, err := getDB()
	if err != nil {
		s.logger.Print("viewer.list: " + err.Error())
		return
	}
	defer db.Close()

	var viewersDB viewer.ViewerResultCollection
	if err = db.Find(&viewersDB).Error; err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}

	userIsAdmin := userDetails.IsAdmin()
	var viewersReturn viewer.ViewerResultCollection
	for _, viewerDB := range viewersDB {
		if userIsAdmin || userDetails.IsAdminOfEntity(*&viewerDB.Entity) {
			viewerDB.Entityname = &viewerDB.Entity
			entityName := getEntityName(viewerDB.Entity)
			if entityName != "" {
				viewerDB.Entityname = &entityName
			}
			viewersReturn = append(viewersReturn, viewerDB)
		}
	}
	return viewersReturn, view, nil
}

func getEntityName(entityID string) string {
	if entityID == "default" {
		return "default"
	}

	showPayload := &entity.ShowbyfieldPayload{Fieldname: "folder", Fieldvalue: entityID}
	entityService := NewEntity(log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile))
	resShow, err := entityService.Showbyfield(nil, showPayload)
	if err != nil {
		// s.logger.Print("getEntityName: " + err.Error())
		return ""
	}
	return resShow.Name
}

// List all layers of a viewer
func (s *viewersrvc) Listlayers(ctx context.Context, p *viewer.ListlayersPayload) (res viewer.ViewerLayerResultCollection, err error) {
	s.logger.Print("viewer.listlayers")

	showPayload := &viewer.ShowbyfieldPayload{Fieldname: "id", Fieldvalue: p.ID}
	resShow, _, err := s.Showbyfield(nil, showPayload)
	if err != nil {
		s.logger.Print("viewer.listlayers: " + err.Error())
		return
	}

	folder := resShow.Folder

	getFilePath := filepath.Join(contextsFolder, *&resShow.Entity, folder+".json")
	mapConfigJSON, err := ioutil.ReadFile(getFilePath)
	if err != nil {
		s.logger.Print("viewer.listlayers: " + err.Error())
		return
	}

	var mapConfig MapConfig
	err = json.Unmarshal(mapConfigJSON, &mapConfig)
	if err != nil {
		s.logger.Print("viewer.listlayers: " + err.Error())
		return
	}

	var layers viewer.ViewerLayerResultCollection
	for layerIndex, layer := range mapConfig.Map.Layers {
		// fmt.Println(layer.Name)

		format := layer.Format
		source := layer.Source
		style := fmt.Sprint(layer.Style)
		url := layer.URL
		visibility := strconv.FormatBool(layer.Visibility)
		layers = append(layers, &viewer.ViewerLayerResult{
			ID:     uint(layerIndex),
			Format: &format,
			Group:  layer.Group,
			Name:   layer.Name,
			// Opacity:    layer.Opacity,
			Source:     &source,
			Style:      &style,
			Title:      layer.Title,
			Type:       layer.Type,
			URL:        &url,
			Visibility: &visibility,
			Catalogurl: &layer.CatalogURL,
		})
	}

	return layers, nil
}

// List all groups of a viewer
func (s *viewersrvc) Listgroups(ctx context.Context, p *viewer.ListgroupsPayload) (res viewer.ViewerGroupResultCollection, err error) {
	s.logger.Print("viewer.listgroups")

	showPayload := &viewer.ShowbyfieldPayload{Fieldname: "id", Fieldvalue: p.ID}
	resShow, _, err := s.Showbyfield(nil, showPayload)
	if err != nil {
		s.logger.Print("viewer.listlayers: " + err.Error())
		return
	}

	folder := resShow.Folder

	getFilePath := filepath.Join(contextsFolder, *&resShow.Entity, folder+".json")
	mapConfigJSON, err := ioutil.ReadFile(getFilePath)
	if err != nil {
		s.logger.Print("viewer.listgroups: " + err.Error())
		return
	}

	var mapConfig MapConfig
	err = json.Unmarshal(mapConfigJSON, &mapConfig)
	if err != nil {
		s.logger.Print("viewer.listgroups: " + err.Error())
		return
	}

	var groups viewer.ViewerGroupResultCollection
	for _, group := range mapConfig.Map.Groups {
		// DEBUG fmt.Println(group.ID + " " + group.Title)

		groups = append(groups, &viewer.ViewerGroupResult{
			ID:    group.ID,
			Title: group.Title,
		})
	}

	return groups, nil
}

// Show viewer by ID
func (s *viewersrvc) Show(ctx context.Context, p *viewer.ShowPayload) (res *viewer.ViewerResult, view string, err error) {
	methodName := "viewer.show"
	s.logger.Print(methodName)

	res = &viewer.ViewerResult{}
	if p.View != nil {
		view = *p.View
	} else {
		view = "tiny"
	}

	db, err := getDB()
	if err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}
	defer db.Close()

	var viewerResult = &viewer.ViewerResult{}

	idSplit := strings.Split(p.ID, "_-_")
	if len(idSplit) == 1 {
		if err = db.First(&viewerResult, p.ID).Error; err != nil {
			s.logger.Print(methodName + ": " + err.Error())
			return
		}
	} else {
		folder := strings.Split(idSplit[1], ".json")[0]
		if err = db.Where("entity='" + idSplit[0] + "' AND folder='" + folder + "'").First(&viewerResult).Error; err != nil {
			s.logger.Print(methodName + ": " + err.Error())
			return
		}
	}

	getFilePath := filepath.Join(contextsFolder, viewerResult.Entity, viewerResult.Folder)
	if !strings.Contains(getFilePath, ".json") {
		getFilePath += ".json"
	}
	mapConfigJSON, err := ioutil.ReadFile(getFilePath)
	if err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}

	var mapConfig MapConfig
	err = json.Unmarshal(mapConfigJSON, &mapConfig)
	if err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}

	viewerResult.Projection = &mapConfig.Projection
	zoom := fmt.Sprintf("%d", mapConfig.Zoom)
	viewerResult.Zoom = &zoom

	centerX := fmt.Sprintf("%f", mapConfig.Center.X)
	viewerResult.Centerx = &centerX
	centerY := fmt.Sprintf("%f", mapConfig.Center.Y)
	viewerResult.Centery = &centerY
	centerCRS := mapConfig.Center.Crs
	viewerResult.Centercrs = &centerCRS

	if mapConfig.MaxExtent != nil {
		xMinExtent := fmt.Sprintf("%f", mapConfig.MaxExtent[0])
		viewerResult.Xminextent = &xMinExtent
		yMinExtent := fmt.Sprintf("%f", mapConfig.MaxExtent[1])
		viewerResult.Yminextent = &yMinExtent
		xMaxExtent := fmt.Sprintf("%f", mapConfig.MaxExtent[2])
		viewerResult.Xmaxextent = &xMaxExtent
		yMaxExtent := fmt.Sprintf("%f", mapConfig.MaxExtent[3])
		viewerResult.Ymaxextent = &yMaxExtent
	}
	viewerResult.Emailaddress = &mapConfig.EmailAddress
	viewerResult.Mapheader = &mapConfig.PrintMapHeader
	viewerResult.Mapnorth = &mapConfig.PrintMapNorth

	return viewerResult, view, nil
}

// Show viewer by field
func (s *viewersrvc) Showbyfield(ctx context.Context, p *viewer.ShowbyfieldPayload) (res *viewer.ViewerResult, view string, err error) {
	methodName := "viewer.showbyfield"
	s.logger.Print(methodName)

	res = &viewer.ViewerResult{}
	if p.View != nil {
		view = *p.View
	} else {
		view = "tiny"
	}

	db, err := getDB()
	if err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}
	defer db.Close()

	var viewerResult = &viewer.ViewerResult{}
	if err = db.Where(p.Fieldname + "='" + p.Fieldvalue + "'").First(&viewerResult).Error; err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}

	return viewerResult, view, nil
}

// Show viewer by 2 fields
func (s *viewersrvc) Showbyfield2(ctx context.Context, p *viewer.Showbyfield2Payload) (res *viewer.ViewerResult, view string, err error) {
	methodName := "viewer.showbyfield2"
	s.logger.Print(methodName)

	res = &viewer.ViewerResult{}
	if p.View != nil {
		view = *p.View
	} else {
		view = "tiny"
	}

	db, err := getDB()
	if err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}
	defer db.Close()

	var viewerResult = &viewer.ViewerResult{}
	if err = db.Where(p.Fieldname + "='" + p.Fieldvalue + "' AND " + p.Fieldname2 + "='" + p.Fieldvalue2 + "'").First(&viewerResult).Error; err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}

	return viewerResult, view, nil
}

// Add new viewer and return its ID.
func (s *viewersrvc) Add(ctx context.Context, p *viewer.AddPayload) (res string, err error) {
	methodName := "viewer.add"
	s.logger.Print(methodName)

	if p == nil || p.Authentication == nil {
		s.logger.Print(methodName + ": Not authorized: no authentication info provided!")
		return "", errUnauthorized
	}

	// Confirm folder with same name does not exist
	showPayload := &viewer.Showbyfield2Payload{Fieldname: "folder", Fieldvalue: p.Viewer.Folder, Fieldname2: "entity", Fieldvalue2: p.Viewer.Entity}
	resShow, _, err := s.Showbyfield2(nil, showPayload)
	if resShow.Folder != "" && fmt.Sprint(resShow.ID) != fmt.Sprint(p.Viewer.ID) {
		s.logger.Print(methodName + ": Folder already exists!")
		return "", errUnauthorized
	}

	// TODO Confirm valid folder name (No slash, or weird characters)

	userDetails := GetUserAuthInfo(*p.Authentication)

	if !userDetails.IsAdmin() && !userDetails.IsAdminOfEntity(*&p.Viewer.Entity) {
		s.logger.Print(methodName + ": Not authorized!")
		return "", errUnauthorized
	}

	db, err := getDB()
	if err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}
	defer db.Close()

	db.NewRecord(p.Viewer)
	if err = db.Create(p.Viewer).Error; err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}
	// TODO Pass these values to the backoffice
	centerX := -28.193665
	centerY := 38.676933
	centerCRS := "EPSG:4326"

	if p.Viewer.Centerx != nil && p.Viewer.Centery != nil {
		centerX, err = strconv.ParseFloat(*p.Viewer.Centerx, 64)
		if err != nil {
			centerX = -28.193665
		}
		centerY, err = strconv.ParseFloat(*p.Viewer.Centery, 64)
		if err != nil {
			centerY = 38.676933
		}
	}
	if p.Viewer.Centercrs != nil {
		centerCRS = *p.Viewer.Centercrs
	}

	projection := "EPSG:4326"
	if p.Viewer.Projection != nil {
		projection = *p.Viewer.Projection
	}

	var zoom int64 = 7
	if p.Viewer.Zoom != nil {
		zoom, err = strconv.ParseInt(*p.Viewer.Zoom, 0, 64)
		if err != nil {
			s.logger.Print(methodName + ": " + err.Error())
		}
	}

	var xMinExtent float64
	var xMaxExtent float64
	var yMinExtent float64
	var yMaxExtent float64
	if p.Viewer.Xminextent != nil {
		xMinExtent, err = strconv.ParseFloat(*p.Viewer.Xminextent, 64)
		if err != nil {
			s.logger.Print(methodName + ": " + err.Error())
		}
	}
	if p.Viewer.Xmaxextent != nil {
		xMaxExtent, err = strconv.ParseFloat(*p.Viewer.Xmaxextent, 64)
		if err != nil {
			s.logger.Print(methodName + ": " + err.Error())
		}
	}
	if p.Viewer.Yminextent != nil {
		yMinExtent, err = strconv.ParseFloat(*p.Viewer.Yminextent, 64)
		if err != nil {
			s.logger.Print(methodName + ": " + err.Error())
		}
	}
	if p.Viewer.Ymaxextent != nil {
		yMaxExtent, err = strconv.ParseFloat(*p.Viewer.Ymaxextent, 64)
		if err != nil {
			s.logger.Print(methodName + ": " + err.Error())
		}
	}
	maxExtent := []float64{xMinExtent, yMinExtent, xMaxExtent, yMaxExtent}

	var textSerchConfig TextSerchConfig
	var catalogServices CatalogServices
	switch projectName {
	case "sigmar":
		textSerchConfig = TextSerchConfig{
			Override: true,
			Services: []TextSerchConfigServices{
				{
					Type:        "wfs",
					Name:        "Endereços Azores",
					DisplayName: "${properties.endereco}",
					SubTitle:    "Endereços Azores",
					Priority:    6,
					Options: Options{
						URL:      "https://sigmar.dram.azores.gov.pt/geoserver/wfs",
						TypeName: "spe.enderecos_azores_ine:enderecos_azores",
						QueriableAttributes: []string{
							"endereco",
						},
						SortBy:      "endereco",
						MaxFeatures: 5,
						SrsName:     "EPSG:4326",
					},
				},
				{
					Type:        "wfs",
					Name:        "Hidrotermal Vents",
					DisplayName: "${properties.nome}",
					SubTitle:    "Hidrotermal Vents",
					Priority:    6,
					Options: Options{
						URL:      "https://sigmar.dram.azores.gov.pt/geoserver/wfs",
						TypeName: "spe.toponimia_zona_maritima:hidrotermal_vents",
						QueriableAttributes: []string{
							"nome",
						},
						SortBy:      "nome",
						MaxFeatures: 5,
						SrsName:     "EPSG:4326",
					},
				},
				{
					Type:        "wfs",
					Name:        "Montes Submarinos",
					DisplayName: "${properties.designacao}",
					SubTitle:    "Montes Submarinos",
					Priority:    6,
					Options: Options{
						URL:      "https://sigmar.dram.azores.gov.pt/geoserver/wfs",
						TypeName: "spe.montes_submarinos:seamounts_2019",
						QueriableAttributes: []string{
							"designacao",
						},
						SortBy:      "designacao",
						MaxFeatures: 5,
						SrsName:     "EPSG:4326",
					},
				},
			},
		}
		catalogServices = CatalogServices{
			Services: CatalogServicesServices{
				GsWms: GsWms{
					Title: "WMS do SIGMAR",
					Type:  "wms",
					URL:   "https://sigmar.dram.azores.gov.pt/geoserver/wms",
				},
			},
		}
	case "idea":
		textSerchConfig = TextSerchConfig{
			Override: true,
			Services: []TextSerchConfigServices{
				{
					Type:        "wfs",
					Name:        "SIGEndA",
					DisplayName: "${properties.ENDERECO}",
					SubTitle:    "SIG dos Endereços dos Açores",
					Priority:    6,
					Options: Options{
						URL:      "https://wssig4.azores.gov.pt/geoserver/sigenda/wfs",
						TypeName: "sigenda:Enderecos_RAA",
						QueriableAttributes: []string{
							"ENDERECO",
						},
						SortBy:      "ENDERECO",
						MaxFeatures: 5,
						SrsName:     "EPSG:4326",
					},
				},
			},
		}
		catalogServices = CatalogServices{
			Services: CatalogServicesServices{
				GsWms: GsWms{
					Title: "Catálogo do Sistema de Metadados dos Açores",
					Type:  "csw",
					URL:   "https://visualizador.idea.azores.gov.pt/geonetwork/srv/por/csw",
				},
				GsCsw: GsWms{
					Title: "LifeNatura",
					Type:  "wms",
					URL:   "https://wssig4.azores.gov.pt/geoserver/wms",
				},
			},
		}
	}

	emailAddress := ""
	if p.Viewer.Emailaddress != nil {
		emailAddress = *p.Viewer.Emailaddress
	}
	printMapHeader := ""
	if p.Viewer.Mapheader != nil {
		printMapHeader = *p.Viewer.Mapheader
	}
	printMapNorth := ""
	if p.Viewer.Mapnorth != nil {
		printMapNorth = *p.Viewer.Mapnorth
	}

	var initialLayers []Layer
	switch projectName {
	case "sigmar":
		initialLayers = []Layer{
			{
				Type:       "wms",
				URL:        "https://sigmar.dram.azores.gov.pt/service",
				Title:      "Open Street Map",
				Name:       "openstreetmap",
				Group:      "background",
				TileSize:   2048,
				Tiled:      false,
				Visibility: true,
			},
		}
	case "idea":
		initialLayers = []Layer{
			{
				Type:       "wms",
				URL:        "https://visualizador.idea.azores.gov.pt/contexts/mapproxy",
				Title:      "Open Street Map",
				Name:       "openstreetmap",
				Group:      "background",
				TileSize:   2048,
				Tiled:      false,
				Visibility: true,
			},
		}
	}

	mapConfig := &MapConfig{
		Version:         2,
		CatalogServices: catalogServices,
		EmailAddress:    emailAddress,
		Map: Map{
			Center: Center{
				X:   centerX,
				Y:   centerY,
				Crs: centerCRS,
			},
			MapOptions: MapOptions{
				View: View{
					Resolutions: getResolutions(projection),
				},
			},
			Layers:          initialLayers,
			Projection:      projection,
			Units:           "m",
			Zoom:            int(zoom),
			TextSerchConfig: textSerchConfig,
			PrintMapHeader:  printMapHeader,
			PrintMapNorth:   printMapNorth,
		},
	}
	if xMinExtent != 0 && xMaxExtent != 0 && yMinExtent != 0 && yMaxExtent != 0 {
		mapConfig.MaxExtent = maxExtent
	}

	mapConfigJSON, err := json.Marshal(mapConfig)
	if err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		// TODO Remove from DB
		return
	}

	filePathWrite := filepath.Join(contextsFolder, p.Viewer.Entity, p.Viewer.Folder)
	if !strings.Contains(filePathWrite, ".json") {
		filePathWrite += ".json"
	}

	os.MkdirAll(filepath.Dir(filePathWrite), os.ModePerm)
	err = ioutil.WriteFile(filePathWrite, mapConfigJSON, 0644)
	if err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		// TODO Remove from DB
		return
	}

	return fmt.Sprint(p.Viewer.ID), nil
}

func getResolutions(srs string) []float64 {
	switch srs {
	case "EPSG:3857":
		fallthrough
	case "EPSG:5014":
		fallthrough
	case "EPSG:5015":
		return []float64{
			5291.66666666668,
			3968.75,
			2645.83333333334,
			1984.375,
			1322.91666666667,
			661.458333333335000,
			264.583333333334000,
			132.291666666667000,
			66.145833333333500,
			26.458333333333400,
			6.614583333333350,
			2.645833333333340,
			1.322916666666670,
			0.661458333333335,
			0.2645833333,
			0.1322916667,
		}
	case "EPSG:5013":
		fallthrough
	case "EPSG:4326":
		fallthrough
	default:
		return []float64{
			0.04758912313,
			0.03569184234,
			0.02379456156,
			0.01784592117,
			0.01189728078,
			0.005948640391,
			0.002379456156,
			0.001189728078,
			0.0005948640391,
			0.0002379456156,
			0.00005948640391,
			0.00002379456156,
			0.00001189728078,
			0.000005948640391,
			0.000002379456156,
			0.000001189728078,
		}
	}
}

// Update the viewer
func (s *viewersrvc) Update(ctx context.Context, p *viewer.UpdatePayload) (err error) {
	methodName := "viewer.update"
	s.logger.Print(methodName)

	// TODO Update this to verify not only folder but also entity
	// Confirm folder with same name does not exist
	// showPayloadFolder := &viewer.ShowbyfieldPayload{Fieldname: "folder", Fieldvalue: p.Viewer.Folder}
	// resShowFolder, _, err := s.Showbyfield(nil, showPayloadFolder)
	// if resShowFolder.Folder != "" && fmt.Sprint(resShowFolder.ID) != fmt.Sprint(p.ID) {
	// 	s.logger.Print(methodName + ": Folder already exists!")
	// 	return errUnauthorized // TODO replace with custom error
	// }

	db, err := getDB()
	if err != nil {
		return
	}
	defer db.Close()

	showPayload := &viewer.ShowPayload{ID: p.ID}
	resShow, _, err := s.Show(nil, showPayload)

	resShow.Title = p.Viewer.Title
	resShow.Name = p.Viewer.Name
	resShow.URL = p.Viewer.URL
	resShow.Anonymous = p.Viewer.Anonymous
	resShow.Inactive = p.Viewer.Inactive
	resShow.Image = p.Viewer.Image
	resShow.Emailaddress = p.Viewer.Emailaddress
	resShow.Mapheader = p.Viewer.Mapheader
	resShow.Mapnorth = p.Viewer.Mapnorth

	folder := resShow.Folder
	resShow.Folder = p.Viewer.Folder
	entity := resShow.Entity
	resShow.Entity = p.Viewer.Entity
	resShow.Entityname = nil

	if err = db.Omit("entityname").Save(resShow).Error; err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}

	// TODO Pass these values to the backoffice
	centerX := -28.193665
	centerY := 38.676933
	centerCRS := "EPSG:4326"

	if p.Viewer.Centerx != nil && p.Viewer.Centery != nil {
		centerX, err = strconv.ParseFloat(*p.Viewer.Centerx, 64)
		if err != nil {
			centerX = -28.193665
		}
		centerY, err = strconv.ParseFloat(*p.Viewer.Centery, 64)
		if err != nil {
			centerY = 38.676933
		}
	}
	if p.Viewer.Centercrs != nil {
		centerCRS = *p.Viewer.Centercrs
	}

	getFilePath := filepath.Join(contextsFolder, entity, folder)
	if !strings.Contains(getFilePath, ".json") {
		getFilePath += ".json"
	}
	mapConfigJSON, err := ioutil.ReadFile(getFilePath)
	if err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}

	var mapConfig MapConfig
	err = json.Unmarshal(mapConfigJSON, &mapConfig)
	if err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}

	center := &Center{
		X:   centerX,
		Y:   centerY,
		Crs: centerCRS,
	}
	mapConfig.Map.Center = *center

	var xMinExtent float64
	var xMaxExtent float64
	var yMinExtent float64
	var yMaxExtent float64
	if p.Viewer.Xminextent != nil {
		xMinExtent, err = strconv.ParseFloat(*p.Viewer.Xminextent, 64)
		if err != nil {
			s.logger.Print(methodName + ": " + err.Error())
		}
	}
	if p.Viewer.Xmaxextent != nil {
		xMaxExtent, err = strconv.ParseFloat(*p.Viewer.Xmaxextent, 64)
		if err != nil {
			s.logger.Print(methodName + ": " + err.Error())
		}
	}
	if p.Viewer.Yminextent != nil {
		yMinExtent, err = strconv.ParseFloat(*p.Viewer.Yminextent, 64)
		if err != nil {
			s.logger.Print(methodName + ": " + err.Error())
		}
	}
	if p.Viewer.Ymaxextent != nil {
		yMaxExtent, err = strconv.ParseFloat(*p.Viewer.Ymaxextent, 64)
		if err != nil {
			s.logger.Print(methodName + ": " + err.Error())
		}
	}
	maxExtent := []float64{xMinExtent, yMinExtent, xMaxExtent, yMaxExtent}

	if xMinExtent != 0 && xMaxExtent != 0 && yMinExtent != 0 && yMaxExtent != 0 {
		mapConfig.MaxExtent = maxExtent
	}

	mapConfig.Map.Projection = *p.Viewer.Projection
	mapConfig.Map.MapOptions.View.Resolutions = getResolutions(*p.Viewer.Projection)
	mapConfig.Map.Units = "m"

	var zoom int64 = 7
	if p.Viewer.Zoom != nil {
		zoom, err = strconv.ParseInt(*p.Viewer.Zoom, 0, 64)
		if err != nil {
			s.logger.Print(methodName + ": " + err.Error())
		}
	}
	mapConfig.Map.Zoom = int(zoom)

	mapConfig.EmailAddress = *p.Viewer.Emailaddress

	mapConfig.PrintMapHeader = *p.Viewer.Mapheader
	mapConfig.PrintMapNorth = *p.Viewer.Mapnorth

	mapConfigJSON, err = json.Marshal(mapConfig)
	if err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}

	filePathWrite := filepath.Join(contextsFolder, p.Viewer.Entity, p.Viewer.Folder)
	if !strings.Contains(filePathWrite, ".json") {
		filePathWrite += ".json"
	}

	err = ioutil.WriteFile(filePathWrite, mapConfigJSON, 0644)
	if err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}

	filePathRemove := filepath.Join(contextsFolder, resShow.Entity, resShow.Folder)
	if !strings.Contains(filePathRemove, ".json") {
		filePathRemove += ".json"
	}

	if p.Viewer.Entity != entity || p.Viewer.Folder != folder {
		err = os.Remove(filePathRemove)
		if err != nil {
			s.logger.Print(methodName + ": " + err.Error())
			return
		}
	}
	return
}

// Update the layers of the viewer
func (s *viewersrvc) Updatelayers(ctx context.Context, p *viewer.UpdatelayersPayload) (err error) {
	methodName := "viewer.updatelayers"
	s.logger.Print(methodName)

	showPayload := &viewer.ShowbyfieldPayload{Fieldname: "id", Fieldvalue: p.ID}
	resShow, _, err := s.Showbyfield(nil, showPayload)
	if err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}
	folder := resShow.Folder

	filePathRead := filepath.Join(contextsFolder, resShow.Entity, folder)
	if !strings.Contains(filePathRead, ".json") {
		filePathRead += ".json"
	}

	mapConfigJSON, err := ioutil.ReadFile(filePathRead)
	if err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}

	var mapConfig MapConfig
	err = json.Unmarshal(mapConfigJSON, &mapConfig)
	if err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}

	var layersAdd []Layer
	for _, viewerAdd := range p.Viewerlayers {
		source := viewerAdd.Source
		if source == nil {
			emptySource := ""
			source = &emptySource
		}
		var style interface{} = viewerAdd.Style
		url := viewerAdd.URL
		if url == nil {
			emptyURL := ""
			url = &emptyURL
		}
		visibility := false
		if viewerAdd.Visibility != nil {
			visibility, err = strconv.ParseBool(*viewerAdd.Visibility)
			if err != nil {
				visibility = true
			}
		}
		layerAdd := &Layer{
			Type:       viewerAdd.Type,
			Title:      viewerAdd.Title,
			Name:       viewerAdd.Name,
			Source:     *source,
			Style:      style,
			Group:      viewerAdd.Group,
			URL:        *url,
			Visibility: visibility,
			TileSize:   2048,
		}

		if viewerAdd.Type == "vector" {

			filePathRead := filepath.Join(geodataFolder, viewerAdd.Name)
			if !strings.Contains(filePathRead, ".geojson") {
				filePathRead += ".geojson"
			}

			geodataJSON, err2 := ioutil.ReadFile(filePathRead)
			if err2 != nil {
				s.logger.Print(methodName + ": " + err.Error())
				return
			}

			var gpx Gpx
			err = json.Unmarshal(geodataJSON, &gpx)
			if err != nil {
				s.logger.Print(methodName + ": " + err.Error())
				return
			}

			layerAdd.Features = gpx.Features

		} else {
			if *viewerAdd.URL == "https://sigmar.dram.azores.gov.pt/service" || *viewerAdd.URL == "https://visualizador.idea.azores.gov.pt/contexts/mapproxy" {
				layerAdd.Tiled = false
				layerAdd.TileSize = 2048
			}

			if viewerAdd.Catalogurl != nil {
				if *viewerAdd.Catalogurl != "" {
					layerAdd.CatalogURL = *viewerAdd.Catalogurl
				}
			}

			// TODO What happens when the WFS is not activated?
			if strings.Index(*viewerAdd.URL, "/geoserver/wms") > 0 {
				layerAdd.Search = Search{
					Type: "wfs",
					URL:  strings.Replace(*viewerAdd.URL, "/geoserver/wms", "/geoserver/wfs", -1),
				}
			}
			if viewerAdd.Type == "tileprovider" || viewerAdd.Type == "osm" {
				switch layerAdd.Name {
				case "DarkMatter":
					layerAdd.Provider = "CartoDB.DarkMatter"
				case "Positron":
					layerAdd.Provider = "CartoDB.Positron"
				case "Mapnik":
					layerAdd.Provider = "OpenStreetMap.Mapnik"
				case "BlackAndWhite":
					layerAdd.Provider = "OpenStreetMap.BlackAndWhite"
				case "DE":
					layerAdd.Provider = "OpenStreetMap.DE"
				case "France":
					layerAdd.Provider = "OpenStreetMap.France"
				case "HOT":
					layerAdd.Provider = "OpenStreetMap.HOT"
				case "OpenTopoMap":
					layerAdd.Provider = "OpenTopoMap"
				}
			}
		}
		layersAdd = append(layersAdd, *layerAdd)
	}

	mapConfig.Layers = layersAdd

	mapConfigJSON, err = json.Marshal(mapConfig)
	if err != nil {
		s.logger.Print("viewer.updatelayers: " + err.Error())
		return
	}

	filePathWrite := filepath.Join(contextsFolder, resShow.Entity, folder)
	if !strings.Contains(filePathWrite, ".json") {
		filePathWrite += ".json"
	}

	err = ioutil.WriteFile(filePathWrite, mapConfigJSON, 0644)
	if err != nil {
		s.logger.Print("viewer.updatelayers: " + err.Error())
		return
	}

	return
}

// Update the groups of the viewer
func (s *viewersrvc) Updategroups(ctx context.Context, p *viewer.UpdategroupsPayload) (err error) {
	methodName := "viewer.updategroups"

	showPayload := &viewer.ShowbyfieldPayload{Fieldname: "id", Fieldvalue: p.ID}
	resShow, _, err := s.Showbyfield(nil, showPayload)
	if err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}
	folder := resShow.Folder

	filePathRead := filepath.Join(contextsFolder, resShow.Entity, folder)
	if !strings.Contains(filePathRead, ".json") {
		filePathRead += ".json"
	}

	mapConfigJSON, err := ioutil.ReadFile(filePathRead)
	if err != nil {
		s.logger.Print("viewer.updategroups: " + err.Error())
		return
	}

	var mapConfig MapConfig
	err = json.Unmarshal(mapConfigJSON, &mapConfig)
	if err != nil {
		s.logger.Print("viewer.updategroups: " + err.Error())
		return
	}

	var groupsAdd []MapGroups
	for _, groupAdd := range p.Viewergroups {
		if groupAdd.ID != "background" {
			groupsAdd = append(groupsAdd, MapGroups{
				ID:       groupAdd.ID,
				Title:    groupAdd.Title,
				Expanded: false,
			})
		}
	}

	mapConfig.Groups = groupsAdd

	mapConfigJSON, err = json.Marshal(mapConfig)
	if err != nil {
		s.logger.Print("viewer.updategroups: " + err.Error())
		return
	}

	filePathWrite := filepath.Join(contextsFolder, resShow.Entity, folder)
	if !strings.Contains(filePathWrite, ".json") {
		filePathWrite += ".json"
	}

	err = ioutil.WriteFile(filePathWrite, mapConfigJSON, 0644)
	if err != nil {
		s.logger.Print("viewer.updategroups: " + err.Error())
		return
	}

	return
}

// Remove viewer
func (s *viewersrvc) Remove(ctx context.Context, p *viewer.RemovePayload) (err error) {
	methodName := "viewer.remove"
	s.logger.Print(methodName)

	db, err := getDB()
	if err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}
	defer db.Close()

	if err = db.Where("id = ?", p.ID).Delete(&viewer.Viewer{}).Error; err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}

	showPayload := &viewer.ShowPayload{ID: p.ID}
	resShow, _, err := s.Show(nil, showPayload)

	filePathRemove := filepath.Join(contextsFolder, resShow.Entity, resShow.Folder)
	if !strings.Contains(filePathRemove, ".json") {
		filePathRemove += ".json"
	}
	err = os.Remove(filePathRemove)
	if err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}

	return
}
