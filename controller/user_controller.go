package controller

import "net/http"

type UserController interface {
	GetInquirybyaccount(w http.ResponseWriter, r *http.Request)
}
