package app

import (
	"log"
	"net/http"

	"github.com/subalgo/gonews/pkg/view"
)

/*func newsView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]
	log.Println(id)

	view.News(w, nil)
}*/

type newsView string

func (id newsView) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(id)
	view.News(w, nil)
}
