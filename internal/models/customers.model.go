package models

import "time"

type CustomerInfo struct {
	Id              int64     `gorm:"AUTO_INCREMENT:yes" sql:"bigint;not null;primary_key"  json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	CustomerId      int64     `json:"customerId"`
	CustomerGroupId int64     `json:"customerGroupId"`
	PhoneNumber     string    `json:"phoneNumber"`
	NickName        string    `json:"nickname"`
	AccessToken     string    `json:"accessToken"`
	RefreshToken    string    `json:"refreshToken"`
	Province        string    `json:"province"`
	ProvinceId      string    `json:"provinceId"`
	Vertical        string    `json:"vertical"`
	VerticalId      string    `json:"verticalId"`
	WarehouseId     string    `json:"warehouseId"`
	Pronouns        string    `json:"pronouns"`
	CartId          string    `json:"cartId"`
}
