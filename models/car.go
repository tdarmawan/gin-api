package models

// type Car struct{
// 	Merk string `json:"merk"`
// 	Harga int `json:"harga"`
// 	Typecars string `json:"typecars"`
// 	Id uint `json:"id"`
// }

type Car struct{
	Merk string `gorm:"varchar(100)"`
	Harga int `gorm:"integer(11)"`
	Typecars string `gorm:"varchar(100)"`
	Id uint `gorm:"primaryKey"`
}
