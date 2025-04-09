package handles

import (
	"marine-backend/internal/db"
	"marine-backend/internal/model"
	"marine-backend/pkg/utils"
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
		return
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

type GetPortsByCodeReq struct {
	Ports []string `json:"ports,omitempty"`
}

func GetPortsByCode(c *gin.Context) {
	var req *GetPortsByCodeReq
	err := c.BindJSON(&req)
	if err != nil {
		utils.Log.Errorf("failed to bind json, err: %v", err)
		return
	}
	if len(req.Ports) == 0 {
		utils.Log.Errorf("empty ports slice!")
		return
	}
	utils.Log.Infof("ports to fetch: %v", req.Ports)
	ports, err := db.GetPortsByCode(req.Ports)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, ports)
	return
}

func GetPortByCode(c *gin.Context) {
	portCode := c.Query("port_code")
	if portCode == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	port, err := db.GetPortByCode(portCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, port)
	return
}
