package controller

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

func AddressesController(ctx *gin.Context) {
	addrs, _ := net.InterfaceAddrs()
	var result []string
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				result = append(result, ipnet.IP.String())
			}
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"addresses": result})
}
