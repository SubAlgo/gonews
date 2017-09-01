package app

import (
	"net/http"

	"github.com/subalgo/gonews/pkg/model"

	"github.com/subalgo/gonews/pkg/view"
)

func newsView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]
	n, err := model.GetNews(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	view.News(w, n)
}

/*--- วิธีที่ 2
type newsView string

func (id newsView) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(id)
	view.News(w, nil)
}
*/
