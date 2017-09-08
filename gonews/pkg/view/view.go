package view

import (
	"net/http"
	"net/url"

	"github.com/subalgo/gonews/pkg/model"
)

// IndexData fff
type IndexData struct {
	List []*model.News
}

// --------------------------------------- RENDER ---------------------------------------

// Index render index view
func Index(w http.ResponseWriter, data *IndexData) {
	render(tpIndex, w, data)
}

// News render news view
func News(w http.ResponseWriter, data *model.News) {
	render(tpNews, w, data)
}

// AdminListData ff
type AdminListData struct {
	List []*model.News
	//CurrentUser
}

type AdminLoginData struct {
	Flash url.Values
}

// AdminLogin renders admin login view
func AdminLogin(w http.ResponseWriter, data *AdminLoginData) {
	render(tpAdminLogin, w, data)
	data.Flash.Del("errors") //ใช้เคลีย Flash message ไม่ให้มัน append ในหน้า login
}

// AdminRegister renders admin login view
func AdminRegister(w http.ResponseWriter, data interface{}) {
	render(tpAdminRegister, w, data)
}

// AdminList renders admin list view
func AdminList(w http.ResponseWriter, data interface{}) {
	render(tpAdminList, w, data)
}

// AdminCreate renders admin create view
func AdminCreate(w http.ResponseWriter, data interface{}) {
	render(tpAdminCreate, w, data)
}

// AdminEdit renders admin edit view
func AdminEdit(w http.ResponseWriter, data interface{}) {
	render(tpAdminEdit, w, data)
}
