package models

// Table of dorayaki stock from each store
type Stock struct {
	StoreID   uint64  `gorm:"primaryKey;" json:"store_id"`
	VariantID uint64  `gorm:"primaryKey;" json:"variant_id"`
	Store     Store   `gorm:"foreignKey:StoreID;References:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Variant   Variant `gorm:"foreignKey:VariantID;References:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Total     uint64  `json:"total"`
}

// Dorayaki variant table
type Variant struct {
	ID          uint64 `gorm:"primary_key;autoIncrement:true" json:"id"`
	Flavour     string `gorm:"type:varchar(255)" json:"flavour"`
	Description string `gorm:"type:text" json:"description"`
	IMG_URL     string `gorm:"type:text" json:"img_url"`
}

// Table of dorayaki store
type Store struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Street   string `gorm:"type:text" json:"street"`
	District string `gorm:"type:varchar(255)" json:"district"`
	Province string `gorm:"type:varchar(255)" json:"province"`
	// Phone_Number string `gorm:"type:varchar(255)" json:"phone_number"`
}

// Table of user
type User struct {
	ID     int64 `gorm:"primaryKey"`
	Name   string
	Status string // admin or superadmin
}
