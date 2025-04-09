package model

// Port
type Port struct {
	ID        int32  `gorm:"column:ID;primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	PortName  string `gorm:"column:portName;NOT NULL" json:"port_name,omitempty"`
	PortCode  string `gorm:"column:portCode;NOT NULL" json:"port_code,omitempty"`
	Latitude  string `gorm:"column:latitude" json:"latitude,omitempty"`
	Longitude string `gorm:"column:longitude" json:"longitude,omitempty"`
}

// TableName 表名
func (p *Port) TableName() string {
	return "port"
}
