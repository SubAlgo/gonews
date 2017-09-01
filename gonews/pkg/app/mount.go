package app

import "net/http"

// Mount mount handler to mux (ชื่อ func ที่เป็นตัวใหญ่นำหน้าจะมีการ Export จำเป็นต้องมีการ commnet)
func Mount(mux *http.ServeMux) {
	mux.HandleFunc("/", index) // list all news
	//mux.HandleFunc("/news/", newsView) // /news/:path
	mux.Handle("/news/", http.StripPrefix("/news", http.HandlerFunc(newsView)))

	/*จัดการ mux สำหรับ news.go แบบที่2
	mux.Handle("/news/", http.StripPrefix("/news", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[1:]
		newsView(id).ServerHTTP(w, r)
	})))
	*/

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/login", adminLogin)   // /admin/login
	adminMux.HandleFunc("/list", adminList)     // /admin/list
	adminMux.HandleFunc("/create", adminCreate) // /admin/create
	adminMux.HandleFunc("/edit", adminEdit)     // /admin/edit

	// mux.Handle("/admin/", onlyAdmin(adminMux))
	//ถ้าไม่ใช้ StripPrefix path ต้องเป็น adminMux.HandleFunc("/admin/login", adminLogin)
	mux.Handle("/admin/", http.StripPrefix("/admin", onlyAdmin(adminMux)))

}

//Middleware สำหรับ check การ Login ของ Admin

func onlyAdmin(h http.Handler) http.Handler {
	return h
}
