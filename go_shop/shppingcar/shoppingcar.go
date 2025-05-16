package shppingcar

type Carts struct {
	Cart_id    int `gorm:"primaryKey;autoIncrement"` // 主键，自增
	User_id    int `gorm:"not null"`                 // 不允许为空
	Product_id int `gorm:"not null"`                 // 不允许为空
	Quantity   int `gorm:"default:1"`                // 默认值为 1
}
