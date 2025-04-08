package db

import (
	"encoding/json"
	"io/ioutil"
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

func GetPortThroughput() ([]*PortThroughput, error) {
	jsonData, err := ioutil.ReadFile("location.json")
	if err != nil {
		logrus.Errorf("failed to reading file: %v", err)
		return nil, err
	}

	type Port struct {
		PortName string `json:"portName,omitempty"`
		PortCode string `json:"portCode,omitempty"`
	}
	var ports []Port

	err = json.Unmarshal(jsonData, &ports)
	if err != nil {
		logrus.Errorf("Error unmarshalling JSON: %v", err)
		return nil, err
	}

	inSql := `SELECT
    SUM(containerTonnage)
FROM
    port_traffic_monthly
WHERE
    arrivalPortCode = ? 
GROUP BY
    MONTH(month)
ORDER BY
    MONTH(month);`
	outSql := `
SELECT
    SUM(containerTonnage)
FROM
    port_traffic_monthly
WHERE
    departurePortCode = ?
GROUP BY
    MONTH(month)
ORDER BY
    MONTH(month);
`
	logrus.Infof("port to retrieve: %v", ports)
	res := make([]*PortThroughput, len(ports))
	for i, port := range ports {
		res[i] = &PortThroughput{
			PortName: port.PortName,
			PortCode: port.PortCode,
		}
		var in, out []float64
		err := db.Raw(inSql, port.PortCode).Scan(&in).Error
		if err != nil {
			logrus.Errorf("failed to retrieve throughput for port : %v, err: %v", port.PortName, err)
			continue
		}
		res[i].In = in
		err = db.Raw(outSql, port.PortCode).Scan(&out).Error
		if err != nil {
			logrus.Errorf("failed to retrieve throughput for port : %v, err: %v", port.PortName, err)
			continue
		}
		res[i].Out = out
	}
	return res, nil
}
