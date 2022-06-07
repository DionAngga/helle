package controller

import (
	"encoding/json"
	"fmt"
	"helle/entity/database"
	"helle/usecase"
	"net/http"
)

type controller struct {
	usecase usecase.AccUsecase
}

func New(usecase usecase.AccUsecase) *controller {
	return &controller{usecase}
}

//dari fungsi ini kita membuat macam2 error untuk menghandle error
func handleError(err interface{}) {
	if err != nil {
		fmt.Println("error")
	}
}

func (c *controller) GetUsernameByAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user *database.TblUserAccount
	err := json.NewDecoder(r.Body).Decode(&user)
	handleError(err)

	User, err := c.usecase.GetUsername(user.Account)
	handleError(err)

	err = json.NewEncoder(w).Encode(&User)
	handleError(err)
}
