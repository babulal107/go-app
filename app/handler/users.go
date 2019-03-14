package handler
 
import (
	"bitbucket.org/go-app/app/common"
	"bitbucket.org/go-app/app/model"
	"encoding/json"
	"net/http"

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
 
	name := vars["name"]
	user := getUserOr404(db, name, w, r)
	if user == nil {
		return
	}
	common.RespondJSON(w, http.StatusOK, user)
}
 
func UpdateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
 
	name := vars["name"]
	user := getUserOr404(db, name, w, r)
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
 
	name := vars["name"]
	user := getUserOr404(db, name, w, r)
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
 
	name := vars["name"]
	user := getUserOr404(db, name, w, r)
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
 
	name := vars["name"]
	user := getUserOr404(db, name, w, r)
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
func getUserOr404(db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *model.User {
	user := model.User{}
	if err := db.First(&user, model.User{Name: name}).Error; err != nil {
		common.RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}
