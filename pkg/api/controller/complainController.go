package controller

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"projeto-ra-api-go/pkg/api/model"
	"projeto-ra-api-go/pkg/api/service"
)

type ComplainCtrl interface {
	SaveComplain(w http.ResponseWriter, r *http.Request)
	DeleteById(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	FindByParam(w http.ResponseWriter, r *http.Request)
}

type ComplainCtrlImpl struct {
	service service.Complain
}

func NewComplainController(service service.Complain) *ComplainCtrlImpl {
	return &ComplainCtrlImpl{service: service}
}

// CreateComplain godoc
// @Summary Create a new complain
// @Description Create a new complain with the input paylod
// @Tags complains
// @Accept  json
// @Produce  json
// @Param complain body model.ComplainIn true "Create complain"
// @Success 200
// @Router /complains [post]
func (p *ComplainCtrlImpl) SaveComplain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	//Unmarsshal vc transforma o json em objeto
	// Marshal vc trasforma o objeto em json
	var in model.ComplainIn
	err = json.Unmarshal(body, &in)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	error := ValidateStruct(&in)
	if error != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	complain, err := p.service.Save(context.Background(), &in)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	encoder := json.NewEncoder(w)
	encoder.Encode(complain)
	w.WriteHeader(http.StatusCreated)

}

// swagger:operation GET /complains/{id} complains getComplain godoc
// ---
// @Summary Returns a single complain.
// @Description Returns a single complain by its ID.
// @Tags complains
// @Accept  json
// @Produce  json
// params:
// - name: id
//   in: path
//   description: id of complain
//   type: int
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/model.ComplainIn"
//   "400":
//     "$ref": "#/responses/err"
//   "401":
//     "$ref": "#/responses/err"
//   "403":
//     "$ref": "#/responses/err"
//   "404":
//     "$ref": "#/responses/err"
//   "500":
//     "$ref": "#/responses/err"
func (p *ComplainCtrlImpl) FindById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	complain, err := p.service.FindByIdWithExternal(context.Background(), vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(complain)
	w.WriteHeader(http.StatusOK)
}

func (p *ComplainCtrlImpl) Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var complainIn model.ComplainIn
	json.NewDecoder(r.Body).Decode(&complainIn)

	error := ValidateStruct(&complainIn)
	if error != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	complain := p.service.Update(context.Background(), complainIn, vars["id"])

	if complain != nil {
		log.Println("Error updating the complain", complain)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (p *ComplainCtrlImpl) DeleteById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	err := p.service.DeleteById(context.Background(), vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

// GetComplains godoc
// @Summary Get details of all complains
// @Description Get details of all complains
// @Tags complains
// @Accept  json
// @Produce  json
// @Success 200 {array} model.ComplainOut
// @Router /complains [get]
func (p *ComplainCtrlImpl) FindByParam(w http.ResponseWriter, r *http.Request) {
	var companyTitle, complainCity string
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	if vars["company"] != "" {
		companyTitle = vars["company"]
	} else {
		companyTitle = r.URL.Query().Get("company")
	}
	if vars["city"] != "" {
		complainCity = vars["city"]
	} else {
		complainCity = r.URL.Query().Get("city")
	}

	complainTitle := r.URL.Query().Get("title")
	complainDescription := r.URL.Query().Get("desc")
	complainState := r.URL.Query().Get("state")

	complains, err := p.service.FindByParamWithExternal(context.Background(), &model.ComplainIn{
		Title:       complainTitle,
		Description: complainDescription,
		Company: model.Company{
			Title:       companyTitle,
			Description: companyTitle,
		},
		Locale: model.Locale{
			City:  complainCity,
			State: complainState,
		},
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	//encoder.Encode(make([]model.ComplainOut, len(complains)))
	encoder.Encode(complains)

	w.WriteHeader(http.StatusOK)
	return

}

func ValidateStruct(v *model.ComplainIn) error {
	var validate *validator.Validate
	validate = validator.New()

	errs := validate.Struct(v)
	if errs != nil {
		return errs
	}
	return nil
}
