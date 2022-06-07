package profilecontroller

import (
	"encoding/json"
	"fmt"
	"helle/entity/request"
	"helle/entity/response"
	"helle/usecase"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type controller struct {
	usecase usecase.ProfileUseCase
}

func New(usecase usecase.ProfileUseCase) *controller {
	return &controller{usecase}
}

func handleError(err interface{}) {
	if err != nil {
		fmt.Println("error")
	}
}

func (c *controller) GetProfilebyUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user *request.Name
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("error")
	}
	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		responseBody := err
		err = json.NewEncoder(w).Encode(responseBody)
		handleError(err)
		return
	}

	User, err := c.usecase.GetProfile(user)
	if err != nil {
		responseBody := response.Validate{Validation: "error", Field: "gorm"}
		err = json.NewEncoder(w).Encode(responseBody)
		handleError(err)
		return
	}
	if User.Username == "" {
		responseBody := response.Validate{Validation: "error", Field: "username not found"}
		err = json.NewEncoder(w).Encode(responseBody)
		handleError(err)
	}
	err = json.NewEncoder(w).Encode(&User)
	handleError(err)

}
