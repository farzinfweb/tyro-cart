package impl

import (
	"cart/domain/model"
	"cart/domain/repo"
	"cart/domain/srvc"
	"golang.org/x/net/context"
)

type cartService struct {
	rpo repo.ICartRepo
}

func (c cartService) AddItem(ctx context.Context, item model.CartItem) (model.CartItem, error) {
	_i, err := c.rpo.GetItemByProductId(ctx, item.ProductId)
	if err != nil {
		return c.rpo.AddItem(ctx, item)
	}
	err = c.rpo.IncreaseItemQuantity(ctx, item.ProductId, item.Quantity)
	if err != nil {
		return model.CartItem{}, err
	}
	return *_i, nil
}

func NewCartService(rpo repo.ICartRepo) srvc.ICartService {
	return cartService{rpo}
}
