package schema

type Area struct {
	ID        int64  `gorm:"column:id;primaryKey;"`
	AreaValue int64  `gorm:"column:area_value"`
	AreaType  string `gorm:"column:type"`
}

func (a Area) TableName() string {
	return "areas"
}
