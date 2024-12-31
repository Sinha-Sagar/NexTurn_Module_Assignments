package service

import (
	"a1-bms-project/model"
	"a1-bms-project/repository"
)

func CreateBlog(blog model.Blog) error {
	return repository.CreateBlog(blog)
}

func GetBlogByID(id int) (model.Blog, error) {
	return repository.GetBlogByID(id)
}

func GetAllBlogs() ([]model.Blog, error) {
	return repository.GetAllBlogs()
}

func UpdateBlog(blog model.Blog) error {
	return repository.UpdateBlog(blog)
}

func DeleteBlog(id int) error {
	return repository.DeleteBlog(id)
}
