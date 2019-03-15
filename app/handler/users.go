package handler

import (
	"bitbucket.org/go-app/app/common"
	"bitbucket.org/go-app/app/model"
	"bitbucket.org/go-app/helper"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)
 
func GetAllUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	users := []model.User{}
	db.Find(&users)
	common.RespondJSON(w, http.StatusOK, users)
}
 
func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := model.User{}
 
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		common.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
 
	if err := db.Save(&user).Error; err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.RespondJSON(w, http.StatusCreated, user)
}
 
func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	user := getUserOr404(db, id, w, r)
	if user == nil {
		return
	}
	common.RespondJSON(w, http.StatusOK, user)
}
 
func UpdateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	user := getUserOr404(db, id, w, r)
	if user == nil {
		return
	}
 
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		common.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
 
	if err := db.Save(&user).Error; err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.RespondJSON(w, http.StatusOK, user)
}
 
func DeleteUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	user := getUserOr404(db, id, w, r)
	if user == nil {
		return
	}
	if err := db.Delete(&user).Error; err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.RespondJSON(w, http.StatusNoContent, nil)
}
 
func DisableUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	user := getUserOr404(db, id, w, r)
	if user == nil {
		return
	}
	user.Disable()
	if err := db.Save(&user).Error; err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.RespondJSON(w, http.StatusOK, user)
}
 
func EnableUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	user := getUserOr404(db, id, w, r)
	if user == nil {
		return
	}
	user.Enable()
	if err := db.Save(&user).Error; err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.RespondJSON(w, http.StatusOK, user)
}
 
// getUserOr404 gets a user instance if exists, or respond the 404 error otherwise
func getUserOr404(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *model.User {
	user := model.User{}
	if err := db.First(&user, id).Error; err != nil {
		common.RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}

//ExportUser.. export all user to csv file.
func ExportUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	users := []model.User{}
	db.Find(&users)
	if users==nil{
		common.RespondError(w, http.StatusNotFound, errors.New("no data found").Error())
		return
	}
	header := []string{"ID","Name", "Email", "City","Age"}
	data := [][]string{header}
	for _,row :=range users{
		var user = []string{}
		user = append(user, strconv.FormatUint(uint64(row.ID), 10))
		user = append(user, row.Name)
		user = append(user, row.Email)
		user = append(user, row.City)
		user = append(user, strconv.FormatUint(uint64(row.Age), 10))

		data = append(data,user)
	}
	err := helper.GenerateCSV("users", data)
	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	message := []string{"CSV Generated Successfully"}
	common.RespondJSON(w, http.StatusOK, message)
}