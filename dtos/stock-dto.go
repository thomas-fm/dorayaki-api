package dtos

import (
	"dorayaki-api/models"
)

/*
{
	id_store: 2,
	variants_quantity: [
		{
			"variant-id":
			"quantity":
		},
		{
			"variant-id":
			"quantity":
		},
	]
}
*/

type VariantsQuantity struct {
	Variant models.Variant
	Total   uint64
}

type StockUpdateDTO struct {
	StoreID          uint64
	VariantsQuantity []VariantsQuantity
}

type StockCreateDTO struct {
	StoreID          uint64
	VariantsQuantity []VariantsQuantity
}

type StockSingleUpdate struct {
	// ID        uint64 `json:"id" form:"id" binding:"required"`
	StoreID   uint64 `json:"store_id" form:"store_id" binding:"required"`
	VariantID uint64 `json:"variant_id" form:"variant_id" binding:"required"`
	Total     uint64 `json:"total" form:"total" binding:"required"`
}

type StockSingleCreate struct {
	StoreID   uint64 `json:"store_id" form:"store_id" binding:"required"`
	VariantID uint64 `json:"variant_id" form:"variant_id" binding:"required"`
	Total     uint64 `json:"total" form:"total" binding:"required"`
}
