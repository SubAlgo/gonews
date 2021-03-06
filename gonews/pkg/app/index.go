package app

import (
	"net/http"

	"github.com/subalgo/gonews/pkg/model"
	"github.com/subalgo/gonews/pkg/view"
)

func index(w http.ResponseWriter, r *http.Request) {

	// ปกติถ้าไม่เจอหน้าจะถูกส่งมาที่หน้า index รวมถึง favicon ด้วย
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	ctx := r.Context()
	username, _ := ctx.Value("username").(string)
	list, err := model.ListNews()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	view.Index(w, &view.IndexData{
		List:     list,
		Username: username,
	})
}
