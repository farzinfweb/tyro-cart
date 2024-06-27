package handler

import (
	"cart/domain/model"
	"cart/domain/srvc"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
)

type AddCartItemReq struct {
	ProductId   string `json:"productId"`
	ProductName string `json:"productName"`
	Quantity    int    `json:"quantity"`
	Price       int32  `json:"price"`
}

type AddCartItemRespData struct {
	Id string `json:"id"`
}

type AddCartItemResp struct {
	Data  *AddCartItemRespData `json:"data"`
	Error ErrorResp            `json:"error"`
}

type ErrorResp struct {
	Message string `json:"message"`
}

func AddCartItem(c echo.Context, cartService srvc.ICartService) error {
	req := new(AddCartItemReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(422, AddCartItemResp{
			Error: ErrorResp{Message: "Unmarshalling problem"},
		})
	}
	i, err := cartService.AddItem(context.TODO(), model.CartItem{
		ProductId:   req.ProductId,
		ProductName: req.ProductName,
		Quantity:    req.Quantity,
		Price:       req.Price,
	})
	if err != nil {
		return c.String(422, err.Error())
	}

	return c.JSON(201, AddCartItemResp{Data: &AddCartItemRespData{Id: i.Id}})
}
