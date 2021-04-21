package idea

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	geoserver "jpmenezes.com/idebo/gen/geoserver"
	// Postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	jwtSecretToken = "42isTheAnswer"
)

var (
	// TODO
	MASTERCRYPTOPASSWORD = "XPTO123456789"

	geoStoreURL        = os.Getenv("GEOSTORE_URL")
	contextsServiceURL = os.Getenv("CONTEXTSSERVICE_URL")
)

// geoserver service example implementation.
// The example methods log the requests and return zero values.
type geoserversrvc struct {
	logger *log.Logger
}

// NewGeoserver returns the geoserver service implementation.
func NewGeoserver(logger *log.Logger) geoserver.Service {
	return &geoserversrvc{logger}
}

func GetUserAuthInfo(authenticationHeader string) *UserDetails {
	if authenticationHeader == jwtSecretToken {
		return &UserDetails{
			User: User{
				Name: "ADMIN",
				Role: "ADMIN",
			},
		}
	}

	var jsonToken string
	if strings.Index(authenticationHeader, "Bearer") == 0 {
		jsonToken = authenticationHeader[7:]
	}
	geostoreUserDetailsURL := contextsServiceURL + "/user?j=" + jsonToken
	client := &http.Client{}
	req, err := http.NewRequest("GET", geostoreUserDetailsURL, nil)
	req.Header.Add("Authorization", authenticationHeader)
	req.Header.Add("Accept", "application/json, text/plain, */*")
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	log.Println(string(body))

	var userDetails UserDetails
	err = json.Unmarshal(body, &userDetails)
	if err != nil {
		return nil
	}

	fmt.Println(userDetails.User.Name)

	return &userDetails
}

// List all stored geoservers
func (s *geoserversrvc) List(ctx context.Context, p *geoserver.ListPayload) (res geoserver.GeoserverResultCollection, view string, err error) {
	view = "default"
	methodName := "geoserver.list"
	s.logger.Print(methodName)

	if p == nil || p.Authentication == nil {
		return nil, view, nil
	}

	userDetails := GetUserAuthInfo(*p.Authentication)
	// DEBUG fmt.Println(userDetails)

	db, err := getDB()
	if err != nil {
		return
	}
	defer db.Close()

	var geoservers geoserver.GeoserverResultCollection
	if p.GeoserverURL != nil && *p.GeoserverURL != "" {
		if err = db.Where("url='" + *p.GeoserverURL + "'").Find(&geoservers).Error; err != nil {
			s.logger.Print(methodName + ": " + err.Error())
			return
		}
	} else {
		if err = db.Find(&geoservers).Error; err != nil {
			s.logger.Print(methodName + ": " + err.Error())
			return
		}
	}

	userIsAdmin := userDetails.IsAdmin()
	var geoserversReturn geoserver.GeoserverResultCollection
	for _, geoserverDB := range geoservers {
		if userIsAdmin || userDetails.IsAdminOfEntity(*&geoserverDB.Entity) {
			if geoserverDB.Password != nil {
				var password string
				if *geoserverDB.Password != "" {
					base64DecodedPassword, err := base64.StdEncoding.DecodeString(*geoserverDB.Password)
					if err == nil {
						password = string(decrypt(base64DecodedPassword, MASTERCRYPTOPASSWORD))
					}
				}
				*geoserverDB.Password = password
			}

			geoserverDB.Entityname = &geoserverDB.Entity
			entityName := getEntityName(geoserverDB.Entity)
			if entityName != "" {
				geoserverDB.Entityname = &entityName
			}

			geoserversReturn = append(geoserversReturn, geoserverDB)
		}
	}
	return geoserversReturn, view, nil
}

// Show geoserver by ID
func (s *geoserversrvc) Show(ctx context.Context, p *geoserver.ShowPayload) (res *geoserver.GeoserverResult, view string, err error) {
	res = &geoserver.GeoserverResult{}
	view = "default"
	methodName := "geoserver.show"
	s.logger.Print(methodName)

	db, err := getDB()
	if err != nil {
		return
	}
	defer db.Close()

	var geoserverResult = &geoserver.GeoserverResult{}
	if err = db.First(&geoserverResult, p.ID).Error; err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}

	return geoserverResult, view, nil
}

// Add new geoserver and return its ID.
func (s *geoserversrvc) Add(ctx context.Context, p *geoserver.AddPayload) (res string, err error) {
	methodName := "geoserver.add"
	s.logger.Print(methodName)

	db, err := getDB()
	if err != nil {
		return
	}
	defer db.Close()

	if p.Geoserver.Password != nil {
		var password string
		if *p.Geoserver.Password != "" {
			password = base64.StdEncoding.EncodeToString(encrypt([]byte(*p.Geoserver.Password), MASTERCRYPTOPASSWORD))
		}
		p.Geoserver.Password = &password
	}

	db.NewRecord(p)
	if err = db.Create(p.Geoserver).Error; err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}

	return fmt.Sprintf("%v", p.Geoserver.ID), nil
}

// Update the geoserver.
func (s *geoserversrvc) Update(ctx context.Context, p *geoserver.UpdatePayload) (err error) {
	methodName := "geoserver.update"
	s.logger.Print(methodName)

	showPayload := &geoserver.ShowPayload{ID: p.ID}
	resShow, _, err := s.Show(nil, showPayload)

	db, err := getDB()
	if err != nil {
		return
	}
	defer db.Close()

	resShow.Name = p.Geoserver.Name
	url, err := url.Parse(p.Geoserver.URL)
	if err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}
	url.RawQuery = ""
	url.Fragment = ""
	resShow.URL = url.String()
	resShow.Username = p.Geoserver.Username
	resShow.Entity = p.Geoserver.Entity
	if p.Geoserver.Password != nil {
		var password string
		if *p.Geoserver.Password != "" {
			password = base64.StdEncoding.EncodeToString(encrypt([]byte(*p.Geoserver.Password), MASTERCRYPTOPASSWORD))
		}
		resShow.Password = &password
	}

	if err = db.Omit("entityname").Save(resShow).Error; err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}

	return nil
}

// Remove geoserver
func (s *geoserversrvc) Remove(ctx context.Context, p *geoserver.RemovePayload) (err error) {
	methodName := "geoserver.remove"
	s.logger.Print(methodName)

	db, err := getDB()
	if err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}
	defer db.Close()

	if err = db.Where("id = ?", p.ID).Delete(&geoserver.Geoserver{}).Error; err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}

	return
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Println(err.Error())
		return []byte("")
	}
	return plaintext
}
