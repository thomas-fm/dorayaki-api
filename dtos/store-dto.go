package dtos

// models when create
type StoreCreateDTO struct {
	Name         string `json:"name" form:"name" binding:"required"`
	Street       string `json:"street" form:"street" binding:"required"`
	District     string `json:"district" form:"district" binding:"required"`
	Province     string `json:"province" form:"province" binding:"required"`
	Phone_Number string `json:"phone_number" form:"phone_number" binding:"required"`
}

// models when update
type StoreUpdateDTO struct {
	ID           uint64 `json:"id" form:"id" binding:"required"`
	Name         string `json:"name" form:"name" binding:"required"`
	Street       string `json:"street" form:"street" binding:"required"`
	District     string `json:"district" form:"district" binding:"required"`
	Province     string `json:"province" form:"province" binding:"required"`
	Phone_Number string `json:"phone_number" form:"phone_number" binding:"required"`
}

type StoreFilterDTO struct {
	District string `json:"district" form:"district" binding:"required"`
	Province string `json:"province" form:"province" binding:"required"`
}
