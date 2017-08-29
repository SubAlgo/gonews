package view

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var (
	//tpIndex = template.New("") //สร้าง temp เปล่า
	// tpIndex = parseTemplate("template/root.tmpl", "template/index.tmpl")
	tpIndex      = parseTemplate("root.tmpl", "index.tmpl")
	tpAdminLogin = parseTemplate("root.tmpl", "admin/login.tmpl")
)

const templateDir = "template"

func joinTemplateDir(files ...string) []string { // func สำหรับจัดการ Path [template/]
	r := make([]string, len(files))
	for i, f := range files {
		r[i] = filepath.Join(templateDir, f)
	}
	return r
}

/*
func init() {
	tpIndex.Funcs(template.FuncMap{}) //ใส่ func เปล่า เพราะถ้าไม่ใส่แล้วไป ParseFile มันจะไม่เห็น
	_, err := tpIndex.ParseFiles("template/root.tmpl", "template/index.tmpl")

	if err != nil {
		panic(err)
	}
	tpIndex = tpIndex.Lookup("root")
}
*/

func parseTemplate(file ...string) *template.Template {
	t := template.New("")       // Create emtpy template
	t.Funcs(template.FuncMap{}) //ใส่ func เปล่า เพราะถ้าไม่ใส่แล้วไป ParseFile มันจะไม่เห็น
	//_, err := t.ParseFiles(file...)
	_, err := t.ParseFiles(joinTemplateDir(file...)...)
	if err != nil {
		panic(err)
	}
	t = t.Lookup("root")
	return t
}

func render(t *template.Template, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}

// Index render index view
func Index(w http.ResponseWriter, data interface{}) {
	render(tpIndex, w, data)
}

// AdminLogin renders admin login view
func AdminLogin(w http.ResponseWriter, data interface{}) {
	render(tpAdminLogin, w, data)
}
