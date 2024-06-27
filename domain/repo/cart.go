package repo

import (
	"cart/domain/model"
	"golang.org/x/net/context"
)

type ICartRepo interface {
	AddItem(context.Context, model.CartItem) (model.CartItem, error)
	RemoveItem(context.Context, string) error
	GetItems(context.Context) ([]model.CartItem, error)
	Clear(context.Context) error
	GetItemByProductId(context.Context, string) (*model.CartItem, error)
	IncreaseItemQuantity(context.Context, string, int) error
}
