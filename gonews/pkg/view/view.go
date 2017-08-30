package view

import "net/http"

// --------------------------------------- RENDER ---------------------------------------

// Index render index view
func Index(w http.ResponseWriter, data interface{}) {
	render(tpIndex, w, data)
}

// AdminList renders admin list view
func AdminList(w http.ResponseWriter, data interface{}) {
	render(tpAdminLogin, w, data)
}

// AdminCreate renders admin create view
func AdminCreate(w http.ResponseWriter, data interface{}) {
	render(tpAdminLogin, w, data)
}

// AdminEdit renders admin edit view
func AdminEdit(w http.ResponseWriter, data interface{}) {
	render(tpAdminLogin, w, data)
}

// AdminLogin renders admin login view
func AdminLogin(w http.ResponseWriter, data interface{}) {
	render(tpAdminLogin, w, data)
}
