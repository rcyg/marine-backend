package handles

import (
	"marine-backend/internal/db"
	"marine-backend/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type PortTrafficReq struct {
	PortCode string `form:"port_code"`
}
type PortTrafficResp struct {
	Traffic []*model.PortTrafficMonthly `json:"traffic"`
}

func GetPortTraffic(c *gin.Context) {
	portCode := c.Query("port_code")
	if portCode == "" {
		c.JSON(http.StatusBadRequest, nil)
	}
	logrus.Debugf("port_code : %v", portCode)

	traffic, err := db.Get12MonthTrafficByPort(portCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, PortTrafficResp{Traffic: traffic})
	return
}
