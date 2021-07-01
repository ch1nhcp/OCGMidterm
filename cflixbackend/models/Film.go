package models

type Film struct {
	Id          int64
	Name        string `json: "name"`
	Price       int64  `json: "price"`
	Description string `json: "description"`
	Saleprice   int    `json: "salePrice"`
	Salepercent int    `json: "salePercent"`
	Img         string `json: "img"`
}

// func (f *Film) TableName() string {
// 	return "film"
// }
