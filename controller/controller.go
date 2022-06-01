package controller

import "net/http"

type Controller interface {
	GetInquirybyaccount(w http.ResponseWriter, r *http.Request)
	GetProfilebyUsername(w http.ResponseWriter, r *http.Request)
	GetUsernameByAccount(w http.ResponseWriter, r *http.Request)
	GetUserPhoneNumber(w http.ResponseWriter, r *http.Request)
}
