package main

import (
	"a1-bms-project/config"
	"a1-bms-project/controller"
	"log"
	"net/http"
)

func main() {
	config.InitDB()

	http.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controller.CreateBlogHandler(w, r)
		} else if r.Method == http.MethodGet {
			id := r.URL.Query().Get("id")
			if id == "" {
				controller.GetAllBlogsHandler(w, r)
			} else {
				controller.GetBlogByIDHandler(w, r)
			}
		} else if r.Method == http.MethodPut {
			controller.UpdateBlogHandler(w, r)
		} else if r.Method == http.MethodDelete {
			controller.DeleteBlogHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
