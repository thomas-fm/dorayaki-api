package dtos

// models when create
type VariantCreateDTO struct {
	Flavour     string `json:"flavour" form:"flavour" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	IMG_URL     string `json:"img_url" form:"img_url" binding:"required"`
}

// models when update
type VariantUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Flavour     string `json:"flavour" form:"flavour" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	IMG_URL     string `json:"img_url" form:"img_url" binding:"required"`
}
