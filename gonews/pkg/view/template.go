package view

import (
	"html/template"
	"log"
	"net/http"
)

var (
	tpIndex = template.New("") //สร้าง temp เปล่า
)

func init() {
	tpIndex.Funcs(template.FuncMap{}) //ใส่ func เปล่า เพราะถ้าไม่ใส่แล้วไป ParseFile มันจะไม่เห็น
	_, err := tpIndex.ParseFiles("template/root.tmpl", "template/index.tmpl")

	if err != nil {
		panic(err)
	}
	tpIndex = tpIndex.Lookup("root")
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
