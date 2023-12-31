package admin

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yzimhao/trading_engine/cmd/haoadm/models"
	"github.com/yzimhao/trading_engine/cmd/haobase/assets"
	"github.com/yzimhao/trading_engine/cmd/haobase/orders"
	"github.com/yzimhao/trading_engine/utils"
	"github.com/yzimhao/trading_engine/utils/app"
)

func SystemMenu(ctx *gin.Context) {
	if ctx.Query("v") == "v2" {
		s := models.SystemMenu{}
		data := s.GetV1SystemInit()
		ctx.JSON(200, data)
	} else {
		s := models.SystemMenu{}
		data := s.GetV1SystemInit()
		ctx.JSON(200, data)
	}

}

func SystemInfo(ctx *gin.Context) {
	info := app.KeepaliveInfo()
	data := make([]string, 0)

	for _, item := range info {
		name := strings.Split(item, ".")[1]
		data = append(data, name)
	}
	db := app.Database().NewSession()
	defer db.Close()

	utils.ResponseOkJson(ctx, gin.H{
		"module":       data,
		"version":      app.Version,
		"build_at":     app.Build,
		"assets_stats": assets.AssetsCheck(),
		"trade_stats":  orders.TradeResultStats(),
	})
}
