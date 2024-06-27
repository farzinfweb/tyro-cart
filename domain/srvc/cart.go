package srvc

import (
	"cart/domain/model"
	"golang.org/x/net/context"
)

type ICartService interface {
	AddItem(context.Context, model.CartItem) (model.CartItem, error)
}
