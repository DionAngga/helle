package controller

import "net/http"

type ProfileController interface {
	GetProfilebyUsername(w http.ResponseWriter, r *http.Request)
	GetInquirybyaccount(w http.ResponseWriter, r *http.Request)
}
