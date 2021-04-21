package geoserverAPI

import "encoding/xml"

type Abstract struct {
	Text    string   `xml:",chardata" json:",omitempty"` // maxLength=123
	XMLName xml.Name `xml:"http://www.opengis.net/wms Abstract,omitempty" json:"Abstract,omitempty"`
}

type AccessConstraints struct {
	Text    string   `xml:",chardata" json:",omitempty"` // maxLength=4
	XMLName xml.Name `xml:"http://www.opengis.net/wms AccessConstraints,omitempty" json:"AccessConstraints,omitempty"`
}

type Address struct {
	XMLName xml.Name `xml:"http://www.opengis.net/wms Address,omitempty" json:"Address,omitempty"`
}

type AddressType struct {
	Text    string   `xml:",chardata" json:",omitempty"` // maxLength=4
	XMLName xml.Name `xml:"http://www.opengis.net/wms AddressType,omitempty" json:"AddressType,omitempty"`
}

type BoundingBox struct {
	AttrCRS  string   `xml:"CRS,attr"  json:",omitempty"`  // maxLength=9
	Attrmaxx string   `xml:"maxx,attr"  json:",omitempty"` // maxLength=19
	Attrmaxy string   `xml:"maxy,attr"  json:",omitempty"` // maxLength=18
	Attrminx string   `xml:"minx,attr"  json:",omitempty"` // maxLength=19
	Attrminy string   `xml:"miny,attr"  json:",omitempty"` // maxLength=18
	XMLName  xml.Name `xml:"http://www.opengis.net/wms BoundingBox,omitempty" json:"BoundingBox,omitempty"`
}

type CRS struct {
	Text    string   `xml:",chardata" json:",omitempty"` // maxLength=9
	XMLName xml.Name `xml:"http://www.opengis.net/wms CRS,omitempty" json:"CRS,omitempty"`
}

type Capability struct {
	Exception *Exception `xml:"http://www.opengis.net/wms Exception,omitempty" json:"Exception,omitempty"` // ZZmaxLength=0
	Layer     *Layer     `xml:"http://www.opengis.net/wms Layer,omitempty" json:"Layer,omitempty"`         // ZZmaxLength=0
	Request   *Request   `xml:"http://www.opengis.net/wms Request,omitempty" json:"Request,omitempty"`     // ZZmaxLength=0
	XMLName   xml.Name   `xml:"http://www.opengis.net/wms Capability,omitempty" json:"Capability,omitempty"`
}

type City struct {
	Text    string   `xml:",chardata" json:",omitempty"` // maxLength=10
	XMLName xml.Name `xml:"http://www.opengis.net/wms City,omitempty" json:"City,omitempty"`
}

type ContactAddress struct {
	Address         *Address         `xml:"http://www.opengis.net/wms Address,omitempty" json:"Address,omitempty"`                 // ZZmaxLength=0
	AddressType     *AddressType     `xml:"http://www.opengis.net/wms AddressType,omitempty" json:"AddressType,omitempty"`         // ZZmaxLength=0
	City            *City            `xml:"http://www.opengis.net/wms City,omitempty" json:"City,omitempty"`                       // ZZmaxLength=0
	Country         *Country         `xml:"http://www.opengis.net/wms Country,omitempty" json:"Country,omitempty"`                 // ZZmaxLength=0
	PostCode        *PostCode        `xml:"http://www.opengis.net/wms PostCode,omitempty" json:"PostCode,omitempty"`               // ZZmaxLength=0
	StateOrProvince *StateOrProvince `xml:"http://www.opengis.net/wms StateOrProvince,omitempty" json:"StateOrProvince,omitempty"` // ZZmaxLength=0
	XMLName         xml.Name         `xml:"http://www.opengis.net/wms ContactAddress,omitempty" json:"ContactAddress,omitempty"`
}

type ContactElectronicMailAddress struct {
	Text    string   `xml:",chardata" json:",omitempty"` // maxLength=29
	XMLName xml.Name `xml:"http://www.opengis.net/wms ContactElectronicMailAddress,omitempty" json:"ContactElectronicMailAddress,omitempty"`
}

type ContactFacsimileTelephone struct {
	XMLName xml.Name `xml:"http://www.opengis.net/wms ContactFacsimileTelephone,omitempty" json:"ContactFacsimileTelephone,omitempty"`
}

type ContactInformation struct {
	ContactAddress               *ContactAddress               `xml:"http://www.opengis.net/wms ContactAddress,omitempty" json:"ContactAddress,omitempty"`                             // ZZmaxLength=0
	ContactElectronicMailAddress *ContactElectronicMailAddress `xml:"http://www.opengis.net/wms ContactElectronicMailAddress,omitempty" json:"ContactElectronicMailAddress,omitempty"` // ZZmaxLength=0
	ContactFacsimileTelephone    *ContactFacsimileTelephone    `xml:"http://www.opengis.net/wms ContactFacsimileTelephone,omitempty" json:"ContactFacsimileTelephone,omitempty"`       // ZZmaxLength=0
	ContactPersonPrimary         *ContactPersonPrimary         `xml:"http://www.opengis.net/wms ContactPersonPrimary,omitempty" json:"ContactPersonPrimary,omitempty"`                 // ZZmaxLength=0
	ContactPosition              *ContactPosition              `xml:"http://www.opengis.net/wms ContactPosition,omitempty" json:"ContactPosition,omitempty"`                           // ZZmaxLength=0
	ContactVoiceTelephone        *ContactVoiceTelephone        `xml:"http://www.opengis.net/wms ContactVoiceTelephone,omitempty" json:"ContactVoiceTelephone,omitempty"`               // ZZmaxLength=0
	XMLName                      xml.Name                      `xml:"http://www.opengis.net/wms ContactInformation,omitempty" json:"ContactInformation,omitempty"`
}

type ContactOrganization struct {
	Text    string   `xml:",chardata" json:",omitempty"` // maxLength=23
	XMLName xml.Name `xml:"http://www.opengis.net/wms ContactOrganization,omitempty" json:"ContactOrganization,omitempty"`
}

type ContactPerson struct {
	Text    string   `xml:",chardata" json:",omitempty"` // maxLength=19
	XMLName xml.Name `xml:"http://www.opengis.net/wms ContactPerson,omitempty" json:"ContactPerson,omitempty"`
}

type ContactPersonPrimary struct {
	ContactOrganization *ContactOrganization `xml:"http://www.opengis.net/wms ContactOrganization,omitempty" json:"ContactOrganization,omitempty"` // ZZmaxLength=0
	ContactPerson       *ContactPerson       `xml:"http://www.opengis.net/wms ContactPerson,omitempty" json:"ContactPerson,omitempty"`             // ZZmaxLength=0
	XMLName             xml.Name             `xml:"http://www.opengis.net/wms ContactPersonPrimary,omitempty" json:"ContactPersonPrimary,omitempty"`
}

type ContactPosition struct {
	Text    string   `xml:",chardata" json:",omitempty"` // maxLength=16
	XMLName xml.Name `xml:"http://www.opengis.net/wms ContactPosition,omitempty" json:"ContactPosition,omitempty"`
}

type ContactVoiceTelephone struct {
	XMLName xml.Name `xml:"http://www.opengis.net/wms ContactVoiceTelephone,omitempty" json:"ContactVoiceTelephone,omitempty"`
}

type Country struct {
	Text    string   `xml:",chardata" json:",omitempty"` // maxLength=5
	XMLName xml.Name `xml:"http://www.opengis.net/wms Country,omitempty" json:"Country,omitempty"`
}

type DCPType struct {
	HTTP    *HTTP    `xml:"http://www.opengis.net/wms HTTP,omitempty" json:"HTTP,omitempty"` // ZZmaxLength=0
	XMLName xml.Name `xml:"http://www.opengis.net/wms DCPType,omitempty" json:"DCPType,omitempty"`
}

type EX_GeographicBoundingBox struct {
	EastBoundLongitude *EastBoundLongitude `xml:"http://www.opengis.net/wms eastBoundLongitude,omitempty" json:"eastBoundLongitude,omitempty"` // ZZmaxLength=0
	NorthBoundLatitude *NorthBoundLatitude `xml:"http://www.opengis.net/wms northBoundLatitude,omitempty" json:"northBoundLatitude,omitempty"` // ZZmaxLength=0
	SouthBoundLatitude *SouthBoundLatitude `xml:"http://www.opengis.net/wms southBoundLatitude,omitempty" json:"southBoundLatitude,omitempty"` // ZZmaxLength=0
	WestBoundLongitude *WestBoundLongitude `xml:"http://www.opengis.net/wms westBoundLongitude,omitempty" json:"westBoundLongitude,omitempty"` // ZZmaxLength=0
	XMLName            xml.Name            `xml:"http://www.opengis.net/wms EX_GeographicBoundingBox,omitempty" json:"EX_GeographicBoundingBox,omitempty"`
}

type Exception struct {
	Format  []*Format `xml:"http://www.opengis.net/wms Format,omitempty" json:"Format,omitempty"` // ZZmaxLength=0
	XMLName xml.Name  `xml:"http://www.opengis.net/wms Exception,omitempty" json:"Exception,omitempty"`
}

type Fees struct {
	Text    string   `xml:",chardata" json:",omitempty"` // maxLength=4
	XMLName xml.Name `xml:"http://www.opengis.net/wms Fees,omitempty" json:"Fees,omitempty"`
}

type Format struct {
	Text    string   `xml:",chardata" json:",omitempty"` // maxLength=53
	XMLName xml.Name `xml:"http://www.opengis.net/wms Format,omitempty" json:"Format,omitempty"`
}

type Get struct {
	OnlineResource *OnlineResource `xml:"http://www.opengis.net/wms OnlineResource,omitempty" json:"OnlineResource,omitempty"` // ZZmaxLength=0
	XMLName        xml.Name        `xml:"http://www.opengis.net/wms Get,omitempty" json:"Get,omitempty"`
}

type GetCapabilities struct {
	DCPType *DCPType  `xml:"http://www.opengis.net/wms DCPType,omitempty" json:"DCPType,omitempty"` // ZZmaxLength=0
	Format  []*Format `xml:"http://www.opengis.net/wms Format,omitempty" json:"Format,omitempty"`   // ZZmaxLength=0
	XMLName xml.Name  `xml:"http://www.opengis.net/wms GetCapabilities,omitempty" json:"GetCapabilities,omitempty"`
}

type GetFeatureInfo struct {
	DCPType *DCPType  `xml:"http://www.opengis.net/wms DCPType,omitempty" json:"DCPType,omitempty"` // ZZmaxLength=0
	Format  []*Format `xml:"http://www.opengis.net/wms Format,omitempty" json:"Format,omitempty"`   // ZZmaxLength=0
	XMLName xml.Name  `xml:"http://www.opengis.net/wms GetFeatureInfo,omitempty" json:"GetFeatureInfo,omitempty"`
}

type GetMap struct {
	DCPType *DCPType  `xml:"http://www.opengis.net/wms DCPType,omitempty" json:"DCPType,omitempty"` // ZZmaxLength=0
	Format  []*Format `xml:"http://www.opengis.net/wms Format,omitempty" json:"Format,omitempty"`   // ZZmaxLength=0
	XMLName xml.Name  `xml:"http://www.opengis.net/wms GetMap,omitempty" json:"GetMap,omitempty"`
}

type HTTP struct {
	Get     *Get     `xml:"http://www.opengis.net/wms Get,omitempty" json:"Get,omitempty"`   // ZZmaxLength=0
	Post    *Post    `xml:"http://www.opengis.net/wms Post,omitempty" json:"Post,omitempty"` // ZZmaxLength=0
	XMLName xml.Name `xml:"http://www.opengis.net/wms HTTP,omitempty" json:"HTTP,omitempty"`
}

type Keyword struct {
	Text    string   `xml:",chardata" json:",omitempty"` // maxLength=54
	XMLName xml.Name `xml:"http://www.opengis.net/wms Keyword,omitempty" json:"Keyword,omitempty"`
}

type KeywordList struct {
	Keyword []*Keyword `xml:"http://www.opengis.net/wms Keyword,omitempty" json:"Keyword,omitempty"` // ZZmaxLength=0
	XMLName xml.Name   `xml:"http://www.opengis.net/wms KeywordList,omitempty" json:"KeywordList,omitempty"`
}

type Layer struct {
	Attropaque               string                      `xml:"opaque,attr"  json:",omitempty"`                                                                          // maxLength=1
	Attrqueryable            string                      `xml:"queryable,attr"  json:",omitempty"`                                                                       // maxLength=1
	Abstract                 *Abstract                   `xml:"http://www.opengis.net/wms Abstract,omitempty" json:"Abstract,omitempty"`                                 // ZZmaxLength=0
	BoundingBox              []*BoundingBox              `xml:"http://www.opengis.net/wms BoundingBox,omitempty" json:"BoundingBox,omitempty"`                           // ZZmaxLength=0
	CRS                      []*CRS                      `xml:"http://www.opengis.net/wms CRS,omitempty" json:"CRS,omitempty"`                                           // ZZmaxLength=0
	EX_GeographicBoundingBox []*EX_GeographicBoundingBox `xml:"http://www.opengis.net/wms EX_GeographicBoundingBox,omitempty" json:"EX_GeographicBoundingBox,omitempty"` // ZZmaxLength=0
	KeywordList              *KeywordList                `xml:"http://www.opengis.net/wms KeywordList,omitempty" json:"KeywordList,omitempty"`                           // ZZmaxLength=0
	Layer                    []*Layer                    `xml:"http://www.opengis.net/wms Layer,omitempty" json:"Layer,omitempty"`                                       // ZZmaxLength=0
	Name                     *Name                       `xml:"http://www.opengis.net/wms Name,omitempty" json:"Name,omitempty"`                                         // ZZmaxLength=0
	Style                    []*Style                    `xml:"http://www.opengis.net/wms Style,omitempty" json:"Style,omitempty"`                                       // ZZmaxLength=0
	Title                    *Title                      `xml:"http://www.opengis.net/wms Title,omitempty" json:"Title,omitempty"`                                       // ZZmaxLength=0
	XMLName                  xml.Name                    `xml:"http://www.opengis.net/wms Layer,omitempty" json:"Layer,omitempty"`
}

type LegendURL struct {
	Attrheight     string          `xml:"height,attr"  json:",omitempty"`                                                      // maxLength=2
	Attrwidth      string          `xml:"width,attr"  json:",omitempty"`                                                       // maxLength=3
	Format         []*Format       `xml:"http://www.opengis.net/wms Format,omitempty" json:"Format,omitempty"`                 // ZZmaxLength=0
	OnlineResource *OnlineResource `xml:"http://www.opengis.net/wms OnlineResource,omitempty" json:"OnlineResource,omitempty"` // ZZmaxLength=0
	XMLName        xml.Name        `xml:"http://www.opengis.net/wms LegendURL,omitempty" json:"LegendURL,omitempty"`
}

type Name struct {
	Text    string   `xml:",chardata" json:",omitempty"` // maxLength=81
	XMLName xml.Name `xml:"http://www.opengis.net/wms Name,omitempty" json:"Name,omitempty"`
}

type OnlineResource struct {
	AttrXlinkSpacehref string   `xml:"http://www.w3.org/1999/xlink href,attr"  json:",omitempty"` // maxLength=264
	AttrXlinkSpacetype string   `xml:"http://www.w3.org/1999/xlink type,attr"  json:",omitempty"` // maxLength=6
	AttrXmlnsxlink     string   `xml:"xmlns xlink,attr"  json:",omitempty"`                       // maxLength=28
	XMLName            xml.Name `xml:"http://www.opengis.net/wms OnlineResource,omitempty" json:"OnlineResource,omitempty"`
}

type Post struct {
	OnlineResource *OnlineResource `xml:"http://www.opengis.net/wms OnlineResource,omitempty" json:"OnlineResource,omitempty"` // ZZmaxLength=0
	XMLName        xml.Name        `xml:"http://www.opengis.net/wms Post,omitempty" json:"Post,omitempty"`
}

type PostCode struct {
	XMLName xml.Name `xml:"http://www.opengis.net/wms PostCode,omitempty" json:"PostCode,omitempty"`
}

type Request struct {
	GetCapabilities *GetCapabilities `xml:"http://www.opengis.net/wms GetCapabilities,omitempty" json:"GetCapabilities,omitempty"` // ZZmaxLength=0
	GetFeatureInfo  *GetFeatureInfo  `xml:"http://www.opengis.net/wms GetFeatureInfo,omitempty" json:"GetFeatureInfo,omitempty"`   // ZZmaxLength=0
	GetMap          *GetMap          `xml:"http://www.opengis.net/wms GetMap,omitempty" json:"GetMap,omitempty"`                   // ZZmaxLength=0
	XMLName         xml.Name         `xml:"http://www.opengis.net/wms Request,omitempty" json:"Request,omitempty"`
}

type Service struct {
	Abstract           *Abstract           `xml:"http://www.opengis.net/wms Abstract,omitempty" json:"Abstract,omitempty"`                     // ZZmaxLength=0
	AccessConstraints  *AccessConstraints  `xml:"http://www.opengis.net/wms AccessConstraints,omitempty" json:"AccessConstraints,omitempty"`   // ZZmaxLength=0
	ContactInformation *ContactInformation `xml:"http://www.opengis.net/wms ContactInformation,omitempty" json:"ContactInformation,omitempty"` // ZZmaxLength=0
	Fees               *Fees               `xml:"http://www.opengis.net/wms Fees,omitempty" json:"Fees,omitempty"`                             // ZZmaxLength=0
	KeywordList        *KeywordList        `xml:"http://www.opengis.net/wms KeywordList,omitempty" json:"KeywordList,omitempty"`               // ZZmaxLength=0
	Name               *Name               `xml:"http://www.opengis.net/wms Name,omitempty" json:"Name,omitempty"`                             // ZZmaxLength=0
	OnlineResource     *OnlineResource     `xml:"http://www.opengis.net/wms OnlineResource,omitempty" json:"OnlineResource,omitempty"`         // ZZmaxLength=0
	Title              *Title              `xml:"http://www.opengis.net/wms Title,omitempty" json:"Title,omitempty"`                           // ZZmaxLength=0
	XMLName            xml.Name            `xml:"http://www.opengis.net/wms Service,omitempty" json:"Service,omitempty"`
}

type StateOrProvince struct {
	XMLName xml.Name `xml:"http://www.opengis.net/wms StateOrProvince,omitempty" json:"StateOrProvince,omitempty"`
}

type Style struct {
	LegendURL *LegendURL `xml:"http://www.opengis.net/wms LegendURL,omitempty" json:"LegendURL,omitempty"` // ZZmaxLength=0
	Name      *Name      `xml:"http://www.opengis.net/wms Name,omitempty" json:"Name,omitempty"`           // ZZmaxLength=0
	Title     *Title     `xml:"http://www.opengis.net/wms Title,omitempty" json:"Title,omitempty"`         // ZZmaxLength=0
	XMLName   xml.Name   `xml:"http://www.opengis.net/wms Style,omitempty" json:"Style,omitempty"`
}

type Title struct {
	Text    string   `xml:",chardata" json:",omitempty"` // maxLength=81
	XMLName xml.Name `xml:"http://www.opengis.net/wms Title,omitempty" json:"Title,omitempty"`
}

type WMS_Capabilities struct {
	AttrXmlnsinspire_common    string      `xml:"xmlns inspire_common,attr"  json:",omitempty"`                                     // maxLength=46
	AttrXmlnsinspire_vs        string      `xml:"xmlns inspire_vs,attr"  json:",omitempty"`                                         // maxLength=50
	AttrXsiSpaceschemaLocation string      `xml:"http://www.w3.org/2001/XMLSchema-instance schemaLocation,attr"  json:",omitempty"` // maxLength=238
	AttrupdateSequence         string      `xml:"updateSequence,attr"  json:",omitempty"`                                           // maxLength=4
	Attrversion                string      `xml:"version,attr"  json:",omitempty"`                                                  // maxLength=5
	AttrXmlnsxlink             string      `xml:"xmlns xlink,attr"  json:",omitempty"`                                              // maxLength=28
	Attrxmlns                  string      `xml:"xmlns,attr"  json:",omitempty"`                                                    // maxLength=26
	AttrXmlnsxsi               string      `xml:"xmlns xsi,attr"  json:",omitempty"`                                                // maxLength=41
	Capability                 *Capability `xml:"http://www.opengis.net/wms Capability,omitempty" json:"Capability,omitempty"`      // ZZmaxLength=0
	Service                    *Service    `xml:"http://www.opengis.net/wms Service,omitempty" json:"Service,omitempty"`            // ZZmaxLength=0
	XMLName                    xml.Name    `xml:"http://www.opengis.net/wms WMS_Capabilities,omitempty" json:"WMS_Capabilities,omitempty"`
}

type EastBoundLongitude struct {
	Text    string   `xml:",chardata" json:",omitempty"` // maxLength=19
	XMLName xml.Name `xml:"http://www.opengis.net/wms eastBoundLongitude,omitempty" json:"eastBoundLongitude,omitempty"`
}

type NorthBoundLatitude struct {
	Text    string   `xml:",chardata" json:",omitempty"` // maxLength=18
	XMLName xml.Name `xml:"http://www.opengis.net/wms northBoundLatitude,omitempty" json:"northBoundLatitude,omitempty"`
}

type SouthBoundLatitude struct {
	Text    string   `xml:",chardata" json:",omitempty"` // maxLength=18
	XMLName xml.Name `xml:"http://www.opengis.net/wms southBoundLatitude,omitempty" json:"southBoundLatitude,omitempty"`
}

type WestBoundLongitude struct {
	Text    string   `xml:",chardata" json:",omitempty"` // maxLength=19
	XMLName xml.Name `xml:"http://www.opengis.net/wms westBoundLongitude,omitempty" json:"westBoundLongitude,omitempty"`
}
