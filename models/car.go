package models


// Car represents the model for an Car
type Car struct{
	Pemilik 		string 		`gorm:"varchar(100)"`
	Merk 				string 		`gorm:"varchar(100)"`
	Harga 			int 			`gorm:"integer(11)"`
	Type 				string 		`gorm:"varchar(100)"`
	Id 					uint 			`gorm:"primaryKey"` 
}