package controller

import (
	"a1-bms-project/model"
	"a1-bms-project/service"
	"encoding/json"
	"net/http"
	"strconv"
)

func CreateBlogHandler(w http.ResponseWriter, r *http.Request) {
	var blog model.Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := service.CreateBlog(blog); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetBlogByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	blog, err := service.GetBlogByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(blog)
}

func GetAllBlogsHandler(w http.ResponseWriter, r *http.Request) {
	blogs, err := service.GetAllBlogs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(blogs)
}

func UpdateBlogHandler(w http.ResponseWriter, r *http.Request) {
	var blog model.Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := service.UpdateBlog(blog); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteBlogHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := service.DeleteBlog(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
