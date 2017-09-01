package view

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
)

var (
	//tpIndex = template.New("") //สร้าง temp เปล่า
	// tpIndex = parseTemplate("template/root.tmpl", "template/index.tmpl")
	tpIndex       = parseTemplate("root.tmpl", "index.tmpl")
	tpNews        = parseTemplate("root.tmpl", "news.tmpl")
	tpAdminLogin  = parseTemplate("root.tmpl", "admin/login.tmpl")
	tpAdminList   = parseTemplate("root.tmpl", "admin/list.tmpl")
	tpAdminCreate = parseTemplate("root.tmpl", "admin/create.tmpl")
	tpAdminEdit   = parseTemplate("root.tmpl", "admin/edit.tmpl")
)

var m = minify.New()

const templateDir = "template"

func joinTemplateDir(files ...string) []string { // func สำหรับจัดการ Path [template/]
	r := make([]string, len(files))
	for i, f := range files {
		r[i] = filepath.Join(templateDir, f)
	}
	return r
}

func init() {
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/javascript", js.Minify)
}

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

/*--------------------------------------- RENDER ---------------------------------------
--------------------------------------- RENDER ---------------------------------------
--------------------------------------- RENDER ---------------------------------------*/

/* วิธี render แบบที่ 1 หลังจากเขียน Header ก็จะสร้างตัวแปร err มารับค่าที่ได้จากการ Execute
func render(t *template.Template, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}
*/

//วิธี render แบบที่ 2 สร้าง buffer มาเก็บค่าก่อน แล้วค่อยสั่ง Execute ที่ Pointer buf
//ให้ Template เขียนเข้ามาให้ตัว buffer แทนการเขียน w http.ResponseWriter โดยตรง
func render(t *template.Template, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf := bytes.Buffer{} //สร้าง buffer
	err := t.Execute(&buf, data)
	// t.Excure => buf => m.minify => w
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m.Minify("text/html", w, &buf)
}

/*RENDER ด้วย Pipe
func render(t *template.Template, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	pr, pw := io.Pipe()
	go func() {
		m.Minify("text/html", w, pr)
	}()

	err := t.Execute(pw, data)
	// t.Excure => buf => m.minify => w
	if err != nil {
		log.Println(err)
		return
	}
}
*/
