package model

import (
	"projeto-ra-api-go/pkg/api/provider/mongo/document"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestComplainIn_ToComplain(t *testing.T) {
	//
	type fields struct {
		ComplainIn ComplainIn
	}
	tests := []struct {
		name   string
		fields fields
		want   *Complain
	}{
		{
			name: " parse complainIn to complain",
			fields: fields{
				ComplainIn: ComplainIn{
					Title:    "Terra",
					Description: "frio",
					Locale: Locale{
						City: "arid",
						State: "Bahia",
					},
					Company: Company{
						Title: "Company 1",
						Description: "Description 1",
					},
				}},
			want: &Complain{

				Title:    "Terra",
				Description: "frio",
				Locale: Locale{
					City: "arid",
					State: "Bahia",
				},
				Company: Company{
					Title: "Company 1",
					Description: "Description 1",
				},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.fields.ComplainIn
			if got := p.ToComplain(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToComplain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplain_ToComplainOut(t *testing.T) {
	id := primitive.NewObjectID()
	type fields struct {
		Complain Complain
	}
	tests := []struct {
		name   string
		fields fields
		want   *ComplainOut
	}{{
		name: " parse complain to complainOut",
		fields: fields{
			Complain: Complain{
				ID:      id,
				Title:    "Terra",
				Description: "frio",
				Locale: Locale{
					City: "arid",
					State: "Bahia",
				},
				Company: Company{
					Title: "Company 1",
					Description: "Description 1",
				},
			}},
		want: &ComplainOut{
			ID:      id,
			Title:    "Terra",
			Description: "frio",
			Locale: Locale{
				City: "arid",
				State: "Bahia",
			},
			Company: Company{
				Title: "Company 1",
				Description: "Description 1",
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.fields.Complain
			if got := p.ToComplainOut(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToComplainOut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplain_ToDocument(t *testing.T) {
	id := primitive.NewObjectID()
	type fields struct {
		Complain Complain
	}
	type args struct {
		id primitive.ObjectID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *document.Complain
	}{
		{name: " parse complain to document.plenet",
			fields: fields{
				Complain: Complain{

					Title:    "Terra",
					Description: "frio",
					Locale: Locale{
						City: "arid",
						State: "Bahia",
					},
					Company: Company{
						Title: "Company 1",
						Description: "Description 1",
					},
				}},
			args: args{id: id},
			want: &document.Complain{
				ID:      id,
				Title:    "Terra",
				Description: "frio",
				Locale: Locale{
					City: "arid",
					State: "Bahia",
				},
				Company: Company{
					Title: "Company 1",
					Description: "Description 1",
				},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.fields.Complain
			if got := p.ToDocument(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}
