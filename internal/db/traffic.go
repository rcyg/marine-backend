package db

import (
	"marine-backend/internal/model"

	"github.com/sirupsen/logrus"
)

func Get12MonthTrafficByPort(port string) ([]*model.PortTrafficMonthly, error) {
	var in, out, res []*model.PortTrafficMonthly
	// 入港
	err := db.Model(&model.PortTrafficMonthly{}).Where(model.PortTrafficMonthly{
		ArrivalPortCode: port,
	}).Scan(&in).Error
	if err != nil {
		logrus.Errorf("failed to retrieve port_traffic_montly for port ArrivalPortCode=%s", port)
		return res, err
	}
	err = db.Model(&model.PortTrafficMonthly{}).Where(model.PortTrafficMonthly{
		DeparturePortCode: port,
	}).Scan(&out).Error
	if err != nil {
		logrus.Errorf("failed to retrieve port_traffic_montly for port DeparturePortCode=%s", port)
		return res, err
	}
	res = append(res, in...)
	res = append(res, out...)
	return res, nil
}
