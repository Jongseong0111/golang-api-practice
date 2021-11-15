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
		err = errors.New("1")
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
		err = errors.New("2")
		return
	}

	productId, err := result.LastInsertId()
	if err != nil {
		err = errors.New("3")
		return
	}
	fmt.Println(productId)

	newProduct, err := dao.GetProductFromProductID(context.Background(), int32(productId))
	if err != nil {
		fmt.Println(err)
		err = errors.New("4")
		return
	}

	return newProduct, err
}

func (receiver ProductService) GetProductList(userId int32) (res []int32, err error) {
	productList, err := dao.GetProductIDFromUser(context.Background(), userId)
	if err != nil {
		return
	}
	return productList, err
}