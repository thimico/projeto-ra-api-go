package dao

import (
	"context"
	"fmt"
	"projeto-ra-api-go/pkg/api/model"
	"projeto-ra-api-go/pkg/api/provider/mongo/document"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Complain struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoComplain(client *mongo.Client, db *mongo.Database) *Complain { //retorna a interface
	return &Complain{client: client, collection: db.Collection("complain")}
}

func (p *Complain) Save(parentContext context.Context, complain *model.Complain) (*model.Complain, error) {
	if complain == nil {
		return nil, fmt.Errorf("Error on parsing")
	}
	locale := document.Locale{
		City:  complain.Locale.City,
		State: complain.Locale.State,
	}
	company := document.Company{
		Title:       complain.Company.Title,
		Description: complain.Company.Description,
	}
	doc := document.Complain{
		Title:       complain.Title,
		Description: complain.Description,
		Locale:      locale,
		Company:     company,
	}

	one, err := p.collection.InsertOne(parentContext, doc)
	if err != nil {
		return nil, err
	}

	if one == nil {
		return nil, err
	}

	result := p.collection.FindOne(parentContext, bson.M{"_id": one.InsertedID.(primitive.ObjectID).String()})

	var docOut document.Complain
	err = result.Decode(&docOut)
	if err != nil {
		return nil, err
	}

	return ToComplain(docOut), nil
}

func (p *Complain) DeleteById(ctx context.Context, id string) error {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = p.collection.DeleteOne(ctx, bson.M{"_id": oID})
	if err != nil {
		return err
	}

	return nil
}

func (p *Complain) Update(ctx context.Context, complain *model.Complain, id string) error {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	document := complain.ToDocument(oID)

	opts := options.Update().SetUpsert(true)

	_, err = p.collection.UpdateOne(ctx, bson.M{"_id": document.ID}, bson.D{{"$set", document}}, opts)

	if err != nil {
		return err
	}

	return nil
}

func (p *Complain) FindById(ctx context.Context, id string) (*model.Complain, error) {

	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := p.collection.FindOne(ctx, bson.M{"_id": oID})

	var doc document.Complain
	err = result.Decode(&doc)
	if err != nil {
		return nil, err
	}

	return ToComplain(doc), nil
}

func (p *Complain) FindByParam(ctx context.Context, param *model.ComplainIn) ([]model.Complain, error) {

	filter := p.getFilter(param)

	result, err := p.collection.Find(ctx, filter)
	if err != nil { // se o erro nao for nulo
		return nil, err
	}

	var documents []document.Complain
	err = result.All(ctx, &documents)
	if err != nil {
		return nil, err
	}

	var complains []model.Complain

	for _, doc := range documents {
		complain := ToComplain(doc)

		complains = append(complains, *complain)
	}

	return complains, nil
}
func (p *Complain) getFilter(params *model.ComplainIn) bson.D {
	filter := bson.D{}

	filter = p.appendFilter("title", params.Title, &filter)
	filter = p.appendFilter("description", params.Description, &filter)
	filter = p.appendFilter("locale.city", params.Locale.City, &filter)
	filter = p.appendFilter("company.title", params.Company.Title, &filter)

	return filter

}
func (p *Complain) appendFilter(field, value string, filter *bson.D) bson.D {
	if len(value) > 0 {
		*filter = append(*filter, bson.E{Key: field, Value: value})
	}

	return *filter
}
func (p *Complain) Check(ctx context.Context) error {
	ctx, _ = context.WithTimeout(ctx, time.Second)
	return p.client.Ping(ctx, nil)
}

func ToComplain(p document.Complain) *model.Complain {

	return &model.Complain{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Locale:      model.Locale{
			City:  p.Locale.City,
			State: p.Locale.State,
		},
		Company:     model.Company{
			Title:       p.Company.Title,
			Description: p.Company.Description,
		},
	}
}
