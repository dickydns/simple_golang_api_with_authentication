package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"` //pkey
	Nama     string `json:"nama" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=4"`
	Status   bool   `gorm:"default:true" json:"status"`
}

type UpdateUser struct {
	Nama     *string `json:"nama" binding:"omitempty"` // hilang di json jika string kosong
	Email    *string `json:"email" binding:"omitempty,email"`
	Password *string `json:"password" binding:"omitempty,min=4"`
	Status   *bool   `gorm:"default:true" json:"status"`
}
