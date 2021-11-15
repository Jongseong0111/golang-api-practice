package dto

import "time"

type CreateProductRequest struct {
	ProductUnit *string `json:"productUnit"`
	ProductName *string `json:"productName"`
	ManufacturerID *int `json:"manufacturerId"`
	RegUserID *int `json:"regUserId"`
}

type CreateProductResponse struct {
	ProductID int32 `json:"productId"`
	ProductUnit *string `json:"productUnit"`
	ProductName *string `json:"productName"`
	ManufacturerID *int `json:"manufacturerId"`
	RegUserID *int `json:"regUserId"`
	RegDateTime time.Time `json:"regDateTime"`
}