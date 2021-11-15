package service

import (
	"context"
	"errors"
	"fmt"
	"tutorial.sqlc.dev/app/domain/product/dto"
	"tutorial.sqlc.dev/app/model"
)

type ProductService struct {}

func (receiver ProductService) CreateProduct(req dto.CreateProductRequest) (res model.ScmProduct, err error){
	duplicateProduct, err := dao.CheckDuplicateProductUnit(context.Background(), req.ProductUnit)

	if err != nil {
		return
	}

	if len(duplicateProduct) > 0 {
		err = errors.New("duplicate Product")
		return
	}

	fmt.Println(req)
	result, err := dao.CreateProduct(context.Background(), model.CreateProductParams{
		ProductName: req.ProductName,
		ProductUnit: req.ProductUnit,
		ManufacturerID: req.ManufacturerID,
		RegUserID: req.RegUserID,
	})

	if err != nil {
		return
	}

	productId, err := result.LastInsertId()
	if err != nil {
		return
	}

	newProduct, err := dao.GetProductFromProductID(context.Background(), int32(productId))
	if err != nil {
		return
	}

	return newProduct, err
}

func (receiver ProductService) GetProductList(userId int) (res []int32, err error) {
	productList, err := dao.GetProductIDFromUser(context.Background(), int32(userId))
	if err != nil {
		return
	}

	if len(productList)==0 {
		err = errors.New(fmt.Sprintf("no product item with userid: %v", userId))
	}

	return productList, err
}