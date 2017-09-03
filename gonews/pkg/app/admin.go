package app

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/subalgo/gonews/pkg/model"
	"github.com/subalgo/gonews/pkg/view"
)

func adminLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		userID, err := model.Login(username, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("UserID: ", userID)
		http.Redirect(w, r, "/admin/list", http.StatusSeeOther)
		return
	}
	view.AdminLogin(w, nil)
}

func adminRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		err := model.Register(username, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "admin/login", http.StatusSeeOther)
		return
	}
	view.AdminRegister(w, nil)
}

func adminList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		action := r.FormValue("action") //<input type="hidden" name="action" value="delete">
		id := r.FormValue("id")         // <input type="hidden" name="id" value="{{.ID.Hex}}">
		if action == "delete" {
			err := model.DeleteNews(id) //Method DeleteNews return error
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		http.Redirect(w, r, "/admin/list", http.StatusSeeOther)
		return
	}
	list, err := model.ListNews()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	view.AdminList(w, &view.AdminListData{
		List: list,
	})
}

func adminCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		n := model.News{
			Title:  r.FormValue("title"),
			Detail: r.FormValue("detail"),
		}

		if file, fileHeader, err := r.FormFile("image"); err == nil {
			defer file.Close()
			//fileName := time.Now().Format(time.Stamp) + "-" + fileHeader.Filename

			fileName := fileHeader.Filename
			fp, err := os.Create("upload/" + fileName) //สร้างไฟล์ไว้ที่ โฟลเดอร์ upload

			if err == nil {
				io.Copy(fp, file)
			}

			fp.Close() //Close file
			n.Image = "/upload/" + fileName
		}

		err := model.CreateNews(n)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/list", http.StatusSeeOther)
		return
	}
	view.AdminCreate(w, nil)
}

func adminEdit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	n, err := model.GetNews(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		n.Title = r.FormValue("title")
		n.Detail = r.FormValue("detail")

		//---- Edit Pic ----
		if file, fileHeader, err := r.FormFile("image"); err == nil {
			defer file.Close()
			//fileName := time.Now().Format(time.Stamp) + "-" + fileHeader.Filename

			fileName := fileHeader.Filename
			fp, err := os.Create("upload/" + fileName) //สร้างไฟล์ไว้ที่ โฟลเดอร์ upload

			if err == nil {
				io.Copy(fp, file)
			}

			fp.Close() //Close file
			n.Image = "/upload/" + fileName
		}

		err := model.UpdateNews(n)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/list", http.StatusSeeOther)
		return
	}

	view.AdminEdit(w, n)
}
