package impl

import (
	"cart/domain/model"
	"cart/domain/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type mongoCartItem struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	ProductId   string             `json:"product_id" bson:"product_id"`
	ProductName string             `json:"product_name" bson:"product_name"`
	Quantity    int32              `json:"quantity" bson:"quantity"`
	Price       int32              `json:"price" bson:"price"`
}

type mongoCartRepo struct {
	db *mongo.Database
}

func (m mongoCartRepo) GetItemByProductId(ctx context.Context, id string) (*model.CartItem, error) {
	var item mongoCartItem
	if err := m.db.Collection("items").FindOne(ctx, bson.M{"product_id": id}).Decode(&item); err != nil {
		return nil, err
	}
	return &model.CartItem{
		Id:          item.Id.Hex(),
		ProductId:   item.ProductId,
		ProductName: item.ProductName,
		Quantity:    int(item.Quantity),
		Price:       item.Price,
	}, nil
}

func (m mongoCartRepo) IncreaseItemQuantity(ctx context.Context, id string, q int) error {
	_, err := m.db.Collection("items").UpdateOne(
		ctx,
		bson.M{"product_id": id},
		bson.M{"$inc": bson.M{"quantity": q}},
	)
	return err
}

func (m mongoCartRepo) AddItem(ctx context.Context, item model.CartItem) (model.CartItem, error) {
	mci := mongoCartItem{
		Id:          primitive.NewObjectID(),
		ProductId:   item.ProductId,
		ProductName: item.ProductName,
		Quantity:    int32(item.Quantity),
		Price:       item.Price,
	}
	_, err := m.db.Collection("items").InsertOne(ctx, mci)
	if err != nil {
		return model.CartItem{}, err
	}
	return model.CartItem{
		Id:          mci.Id.Hex(),
		ProductId:   mci.ProductId,
		ProductName: mci.ProductName,
		Quantity:    int(mci.Quantity),
		Price:       mci.Price,
	}, nil
}

func (m mongoCartRepo) RemoveItem(context context.Context, s string) error {
	//TODO implement me
	panic("implement me")
}

func (m mongoCartRepo) GetItems(context context.Context) ([]model.CartItem, error) {
	//TODO implement me
	panic("implement me")
}

func (m mongoCartRepo) Clear(context context.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewCartRepo(db *mongo.Database) repo.ICartRepo {
	return &mongoCartRepo{db}
}
