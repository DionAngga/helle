package profilecontroller

import (
	"encoding/json"
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

func (c *controller) GetProfilebyUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user *request.Name
	json.NewDecoder(r.Body).Decode(&user)
	validate := validator.New()
	errs := validate.Struct(user)
	if errs != nil {
		w.Header().Add("Content-Type", "application/json")
		responseBody := errs
		json.NewEncoder(w).Encode(responseBody)
	} else {
		User, err := c.usecase.GetProfile(user)
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
}
