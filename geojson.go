package idea

type Gpx struct {
	Type     string    `json:"type"`
	Name     string    `json:"name"`
	CRS      CRS       `json:"crs"`
	Features []Feature `json:"features"`
}

type CRS struct {
	Type       string        `json:"type"`
	Properties CRSProperties `json:"properties"`
}

type CRSProperties struct {
	Name string `json:"name"`
}

type GpxFeature struct {
	Type       string            `json:"type"`
	Properties FeatureProperties `json:"properties"`
	Geometry   GpxGeometry       `json:"geometry"`
}

type GpxGeometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type FeatureProperties struct {
	Ele                       float64 `json:"ele"`
	Time                      string  `json:"time"`
	GpxtpxTrackPointExtension int64   `json:"gpxtpx_TrackPointExtension"`
}
