package idea

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	entity "jpmenezes.com/idebo/gen/entity"
)

// entity service example implementation.
// The example methods log the requests and return zero values.
type entitysrvc struct {
	logger *log.Logger
}

// NewEntity returns the entity service implementation.
func NewEntity(logger *log.Logger) entity.Service {
	return &entitysrvc{logger}
}

// List all stored entities
func (s *entitysrvc) List(ctx context.Context, p *entity.ListPayload) (res entity.EntityResultCollection, err error) {
	methodName := "entity.list"
	s.logger.Print(methodName)

	if p == nil || p.Authentication == nil {
		return nil, nil
	}

	userDetails := GetUserAuthInfo(*p.Authentication)
	// DEBUG fmt.Println(userDetails)

	db, err := getDB()
	if err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}
	defer db.Close()

	var entities entity.EntityResultCollection
	if err = db.Find(&entities).Error; err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}
	// log.Println(viewers)

	userIsAdmin := userDetails.IsAdmin()
	var entitiesReturn entity.EntityResultCollection
	for _, entity := range entities {
		if userIsAdmin || userDetails.IsAdminOfEntity(entity.Folder) {
			entitiesReturn = append(entitiesReturn, entity)
		}
	}
	return entitiesReturn, nil
}

// Show entity by ID
func (s *entitysrvc) Show(ctx context.Context, p *entity.ShowPayload) (res *entity.EntityResult, err error) {
	res = &entity.EntityResult{}
	methodName := "entity.show"
	s.logger.Print(methodName)

	db, err := getDB()
	if err != nil {
		return
	}
	defer db.Close()

	var entityResult = &entity.EntityResult{}
	if err = db.First(&entityResult, p.ID).Error; err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}

	return entityResult, nil
}

type UserGroupDetails struct {
	UserGroup UserGroup `json:"UserGroup"`
}

type UserGroup struct {
	GroupName   string `json:"groupName"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

func createUserGroup(p *entity.AddPayload, admin bool) {
	userGroup := &UserGroup{
		GroupName:   p.Entity.Folder,
		Status:      "modified",
		Description: p.Entity.Name,
	}
	if admin {
		userGroup.GroupName += " Admins"
		userGroup.Description += " - Administradores"
	} else {
		userGroup.GroupName += " Users"
		userGroup.Description += " - Utilizadores"
	}
	userGroupDetails := &UserGroupDetails{
		UserGroup: *userGroup,
	}
	userGroupDetailsJSON, err := json.Marshal(userGroupDetails)
	if err != nil {
		return
	}

	groupsURL := contextsServiceURL + "/group"
	client := &http.Client{}
	req, err := http.NewRequest("POST", groupsURL, bytes.NewBuffer(userGroupDetailsJSON))
	req.Header.Add("Authorization", *p.Authentication)
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	_, err = client.Do(req)
	if err != nil {
		return
	}
}

// Show entity by field
func (s *entitysrvc) Showbyfield(ctx context.Context, p *entity.ShowbyfieldPayload) (res *entity.EntityResult, err error) {
	res = &entity.EntityResult{}
	methodName := "entity.showbyfield"
	s.logger.Print(methodName)

	res = &entity.EntityResult{}

	db, err := getDB()
	if err != nil {
		s.logger.Print("entity.showbyfield: " + err.Error())
		return
	}
	defer db.Close()

	var entityResult = &entity.EntityResult{}
	if err = db.Where(p.Fieldname + "='" + p.Fieldvalue + "'").First(&entityResult).Error; err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}

	return entityResult, nil
}

// Add new entity and return its ID.
func (s *entitysrvc) Add(ctx context.Context, p *entity.AddPayload) (res string, err error) {
	methodName := "entity.add"
	s.logger.Print(methodName)

	createUserGroup(p, true)
	createUserGroup(p, false)

	db, err := getDB()
	if err != nil {
		return
	}
	defer db.Close()

	db.NewRecord(p.Entity)
	if err = db.Create(p.Entity).Error; err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}
	return fmt.Sprintf("%v", p.Entity.ID), nil
}

// Update the entity.
func (s *entitysrvc) Update(ctx context.Context, p *entity.UpdatePayload) (err error) {
	methodName := "entity.update"
	s.logger.Print(methodName)

	showPayload := &entity.ShowPayload{ID: p.ID}
	resShow, err := s.Show(nil, showPayload)

	db, err := getDB()
	if err != nil {
		return
	}
	defer db.Close()

	resShow.Name = p.Entity.Name
	resShow.Folder = p.Entity.Folder
	resShow.Inactive = p.Entity.Inactive

	if err = db.Save(resShow).Error; err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}

	return nil
}

// Remove entity
func (s *entitysrvc) Remove(ctx context.Context, p *entity.RemovePayload) (err error) {
	methodName := "entity.remove"
	s.logger.Print(methodName)

	db, err := getDB()
	if err != nil {
		return
	}
	defer db.Close()

	if err = db.Where("id = ?", p.ID).Delete(&entity.Entity{}).Error; err != nil {
		s.logger.Print(methodName + ": " + err.Error())
		return
	}

	return
}
