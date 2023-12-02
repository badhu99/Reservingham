package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/badhu99/content-management-service/internal/dto"
	"github.com/badhu99/content-management-service/internal/entity"
	"github.com/badhu99/content-management-service/internal/utility"
	"github.com/gorilla/mux"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
)

func (data *HandlerData) GetDrafts(w http.ResponseWriter, r *http.Request) {

	pageNumber, err := strconv.Atoi(r.URL.Query().Get("pageNumber"))
	if err != nil {
		pageNumber = 1
	}
	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil {
		pageSize = 12
	}

	companyIdString, _ := r.Context().Value("companyId").(string)
	companyId := mssql.UniqueIdentifier{}
	companyId.UnmarshalJSON([]byte(companyIdString))

	entityDrafts := []entity.Draft{}
	var itemsCount int64

	data.Database.Model(entityDrafts).
		Where(entity.Draft{CompanyID: companyId}).
		Count(&itemsCount).
		Offset((pageNumber - 1) * pageSize).Limit(pageSize).
		Scan(&entityDrafts)

	dtoDrafts := []dto.Draft{}

	for _, k := range entityDrafts {
		dtoDrafts = append(dtoDrafts, dto.Draft{
			Id:   k.ID.String(),
			Name: k.Name,
		})
	}

	jsonResponse, _ := json.Marshal(dto.Pagination[dto.Draft]{
		Count: int(itemsCount),
		Page:  pageNumber,
		Size:  pageSize,
		Items: dtoDrafts,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(jsonResponse))
}

func (data *HandlerData) CreateDraft(w http.ResponseWriter, r *http.Request) {

	companyIdString, _ := r.Context().Value("companyId").(string)
	companyId := mssql.UniqueIdentifier{}
	companyId.UnmarshalJSON([]byte(companyIdString))

	userIdString, _ := r.Context().Value("userId").(string)
	userId := mssql.UniqueIdentifier{}
	userId.UnmarshalJSON([]byte(userIdString))

	dtoDraft := dto.DraftCreate{}
	err, code := utility.ValidateBody(&dtoDraft, r.Body)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	entityDraft := entity.Draft{}
	errorGorm := data.Database.Where(entity.Draft{Name: dtoDraft.Name, CompanyID: companyId}).First(&entityDraft).Error

	log.Println(entityDraft.ID)
	if errorGorm != gorm.ErrRecordNotFound {
		http.Error(w, fmt.Sprintf("Draft with name %s already exists", dtoDraft.Name), http.StatusNotFound)
		return
	}

	entityDraft.Name = dtoDraft.Name
	entityDraft.CompanyID = companyId

	entityDocument := entity.Document{
		Name: fmt.Sprintf("%s_%s", dtoDraft.Name, getTimeStamp()),
		Path: "",
	}

	data.Database.Create(&entityDraft)
	data.Database.Create(&entityDocument)

	log.Println(entityDraft.ID.String())
	log.Println(entityDocument.ID.String())

	entityDraftHistory := entity.DraftHistory{
		DraftID:  entityDraft.ID,
		FileID:   entityDocument.ID,
		UserID:   userId,
		Date:     time.Now(),
		Title:    dtoDraft.Title,
		Message:  dtoDraft.Message,
		Draft:    entityDraft,
		Document: entityDocument,
	}

	data.Database.Create(&entityDraftHistory)

	w.WriteHeader(http.StatusCreated)
}

func (data *HandlerData) UpdateDraftData(w http.ResponseWriter, r *http.Request) {

}

func (data *HandlerData) AddDraftHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	draftIdString := vars["draftId"]

	draftId := mssql.UniqueIdentifier{}
	draftId.UnmarshalJSON([]byte(draftIdString))

	userIdString, _ := r.Context().Value("userId").(string)
	userId := mssql.UniqueIdentifier{}
	userId.UnmarshalJSON([]byte(userIdString))

	dtoDraftHistory := dto.DraftHistory{}
	err, code := utility.ValidateBody(&dtoDraftHistory, r.Body)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	entityDraft := entity.Draft{}

	errorGorm := data.Database.Where(entity.Draft{ID: draftId}).First(&entityDraft).Error

	if errorGorm == gorm.ErrRecordNotFound {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	} else if errorGorm != nil {
		http.Error(w, errorGorm.Error(), http.StatusInternalServerError)
		return
	}

	entityDocument := entity.Document{
		Name: fmt.Sprintf("%s_%s", entityDraft.Name, getTimeStamp()),
		Path: "",
	}

	data.Database.Create(&entityDocument)

	entityDraftHistory := entity.DraftHistory{
		DraftID: draftId,
		FileID:  entityDocument.ID,
		UserID:  userId,
		Date:    time.Now(),
		Title:   dtoDraftHistory.Title,
		Message: dtoDraftHistory.Title,
	}

	data.Database.Create(&entityDraftHistory)

	w.WriteHeader(http.StatusCreated)
}

func getTimeStamp() string {
	return time.Now().Format("20060102_150405")
}
