package brands

import "go_shop/database"

type Brands struct {
	Brand_id     int    `gorm:""`
	English_name string `gorm:"size:100"`
	Chinese_name string `gorm:"size:100"`
	Grade        string `gorm:"size:2"`
}

func GetBrands() []Brands {
	var brands []Brands

	database.Gdb.Raw("select b.brand_id,b.english_name,b.chinese_name,b.grade from brands b").Scan(&brands)

	return brands
}
