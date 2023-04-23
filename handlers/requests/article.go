package requests

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type ArticleForm struct{}

type CreateArticleForm struct {
	Title   string `form:"title" json:"title" binding:"required,min=3,max=100"`
	Content string `form:"content" json:"content" binding:"required,min=3,max=1000"`
}

func getMessage(field, tag string, errMsg ...string) string {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the " + field
		}
		return errMsg[0]
	case "min", "max":
		return field + " should be between 3 to 1000 characters"
	default:
		return "Something went wrong, please try again later"
	}
}

func (f ArticleForm) Create(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Title":
				return getMessage("Title", err.Tag())
			case "Content":
				return getMessage("Content", err.Tag())
			}
		}
	default:
		return "Invalid request"
	}
	return "Something went wrong, please try again later"
}

func (f ArticleForm) Update(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Title":
				return getMessage("Title", err.Tag())
			case "Content":
				return getMessage("Content", err.Tag())
			}
		}
	default:
		return "Invalid request"
	}
	return "Something went wrong, please try again later"
}
