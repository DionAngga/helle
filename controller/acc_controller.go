package controller

import "net/http"

type AccController interface {
	GetUsernameByAccount(w http.ResponseWriter, r *http.Request)
	GetUserPhoneNumber(w http.ResponseWriter, r *http.Request)
}
