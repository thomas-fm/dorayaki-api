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
