package model

// PortTrafficMonthly
type PortTrafficMonthly struct {
	ID                int32  `gorm:"column:ID;primary_key;AUTO_INCREMENT"`
	DeparturePortCode string `gorm:"column:departurePortCode;NOT NULL"`
	ArrivalPortCode   string `gorm:"column:arrivalPortCode;NOT NULL"`
	Month             string `gorm:"column:month;NOT NULL"`
	VoyageCount       int32  `gorm:"column:voyageCount;NOT NULL"`
	ContainerTonnage  string `gorm:"column:containerTonnage;NOT NULL"`
}

// TableName 表名
func (p *PortTrafficMonthly) TableName() string {
	return "port_traffic_monthly"
}
