package handles

import (
	"marine-backend/internal/db"
	"marine-backend/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type PortTrafficByMonthReq struct {
}
type PortTrafficByMonthResp struct {
	Traffic []*model.PortTrafficMonthly `json:"traffic"`
}

func GetPortTrafficByMonth(c *gin.Context) {
	portCode := c.Query("port_code")
	if portCode == "" {
		c.JSON(http.StatusBadRequest, nil)
	}
	logrus.Debugf("port_code : %v", portCode)

	traffic, err := db.Get12MonthTrafficByPort(portCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, PortTrafficByMonthResp{Traffic: traffic})
	return
}

type PortThroughputByMonthReq struct{}
type PortThroughputByMonthResp struct {
	Data []*db.PortThroughput `json:"data,omitempty"`
}

func GetPortThroughputByMonth(c *gin.Context) {
	data, err := db.GetPortThroughput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, PortThroughputByMonthResp{Data: data})
	return
}
