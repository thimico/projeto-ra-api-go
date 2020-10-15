package model

import (
	"projeto-ra-api-go/pkg/api/provider/mongo/document"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Complain struct {
	ID          primitive.ObjectID
	Title       string
	Description string
	Locale      Locale  `json:"locale"`
	Company     Company `json:"company"`
}

type Locale struct {
	City  string `json:"city"`
	State string `json:"state"`
}

type Company struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ComplainIn struct {
	Title       string  `json:"title"    validate:"required" `
	Description string  `json:"description" validate:"required" `
	Locale      Locale  `json:"locale"`
	Company     Company `json:"company" validate:"required" `
}

type ComplainOut struct {
	ID             primitive.ObjectID `json:"id"`
	Title          string             `json:"title"`
	Description    string             `json:"description"`
	Locale         Locale             `json:"locale"`
	Company        Company            `json:"company"`
	CountPageViews int                `json:"countPageViews"`
	IsOnTop10BadRA bool               `json:"isOnTop10BadRA"`
}

func (p *ComplainIn) ToComplain() *Complain {
	return &Complain{
		Title:       p.Title,
		Description: p.Description,
		Locale:      p.Locale,
		Company:     p.Company,
	}
}

func (p *Complain) ToComplainOut() *ComplainOut {
	return &ComplainOut{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Locale:      p.Locale,
		Company:     p.Company,
	}
}

func (p *Complain) ToDocument(id primitive.ObjectID) *document.Complain {
	locale := document.Locale{
		City:  p.Locale.City,
		State: p.Locale.State,
	}
	company := document.Company{
		Title:       p.Company.Title,
		Description: p.Company.Description,
	}
	return &document.Complain{
		ID:          id,
		Title:       p.Title,
		Description: p.Description,
		Locale:      locale,
		Company:     company,
	}
}
