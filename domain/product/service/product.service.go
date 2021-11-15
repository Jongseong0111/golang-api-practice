package service

import (
	"context"
	"errors"
	"fmt"
	"tutorial.sqlc.dev/app/db"
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

func (receiver ProductService) UpdateProduct(item dto.UpdateProductRequest) (err error) {

	tx, _ := db.GetConnection().Db.Begin()
	daoTx := dao.WithTx(tx)

	defer tx.Rollback()

	productUserId, err := daoTx.GetUserFromProduct(context.Background(), int32(item.ProductID))
	if err != nil {
		err = errors.New("product doesn't exist")
		return
	}

	userId := item.UserID
	if productUserId != int32(userId) {
		err =errors.New("product bad")
		return
	}

	var productName *string
	var productUnit *string

	if item.ProductName != nil {
		productName = item.ProductName
	}
	err = daoTx.UpdateProductName(context.Background(), model.UpdateProductNameParams{
		ProductName: productName,
		ProductID: int32(item.ProductID),
	})
	if err != nil {
		err = errors.New("1")
		return
	}

	if item.ProductUnit != nil {
		productUnit = item.ProductUnit
	}

	duplicateProduct, err := dao.CheckDuplicateProductUnit(context.Background(), productUnit)

	if err != nil {
		err = errors.New("2")
		return
	}

	if len(duplicateProduct) > 0 {
		err = errors.New("duplicate Product")
		return
	}
	err = daoTx.UpdateProductUnit(context.Background(), model.UpdateProductUnitParams{
		ProductUnit: productUnit,
		ProductID: int32(item.ProductID),
	})

	tx.Commit()
	return
}