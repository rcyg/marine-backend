package model

// Port
type Port struct {
	ID        int32  `gorm:"column:ID;primary_key;AUTO_INCREMENT"`
	PortName  string `gorm:"column:portName;NOT NULL"`
	PortCode  string `gorm:"column:portCode;NOT NULL"`
	Latitude  string `gorm:"column:latitude"`
	Longitude string `gorm:"column:longitude"`
}

// TableName 表名
func (p *Port) TableName() string {
	return "port"
}
