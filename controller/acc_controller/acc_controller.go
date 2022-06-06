package controller

import (
	"encoding/json"
	"helle/entity/database"
	"helle/entity/response"
	"helle/usecase"
	"net/http"
)

type controller struct {
	usecase usecase.AccUsecase
}

func New(usecase usecase.AccUsecase) *controller {
	return &controller{usecase}
}

func (c *controller) GetUsernameByAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user *database.TblUserAccount
	json.NewDecoder(r.Body).Decode(user)

	User, err := c.usecase.GetUsername(user.Account)
	if err != nil {
		responseBody := response.Validate{Validation: "error", Field: "gorm"}
		json.NewEncoder(w).Encode(responseBody)
	}

	if User.Username == "" {
		responseBody := response.Validate{Validation: "error", Field: "username not found"}
		json.NewEncoder(w).Encode(responseBody)
	} else {
		json.NewEncoder(w).Encode(&User)
	}
}
