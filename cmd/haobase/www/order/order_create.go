package order

import (
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/yzimhao/trading_engine/cmd/haobase/assets"
	"github.com/yzimhao/trading_engine/cmd/haobase/base/varieties"
	"github.com/yzimhao/trading_engine/cmd/haobase/orders"
	"github.com/yzimhao/trading_engine/trading_core"
	"github.com/yzimhao/trading_engine/utils"
)

type order_create_request_args struct {
	Symbol    string                 `json:"symbol" binding:"required"`
	Side      trading_core.OrderSide `json:"side" binding:"required"`
	OrderType trading_core.OrderType `json:"order_type" binding:"required"`
	Price     string                 `json:"price" example:"1.00"`
	Quantity  string                 `json:"qty" example:"12"`
	Amount    string                 `json:"amount" example:"100.00"`
}

func Create(ctx *gin.Context) {
	var req order_create_request_args
	if err := ctx.BindJSON(&req); err != nil {
		utils.ResponseFailJson(ctx, err.Error())
		return
	}

	var info *orders.Order
	var err error

	user_id := ctx.MustGet("user_id").(string)

	//todo基础的验证
	cfg := varieties.NewTradingVarieties(req.Symbol)
	if cfg == nil {
		utils.ResponseFailJson(ctx, "非法交易对")
		return
	}

	if req.Side == trading_core.OrderSideSell {
		if assets.BalanceOfAvailable(user_id, cfg.Target.Symbol).Cmp(utils.D("0")) <= 0 {
			utils.ResponseFailJson(ctx, "持有数量不足")
			return
		}
	} else if req.Side == trading_core.OrderSideBuy {
		if assets.BalanceOfAvailable(user_id, cfg.Base.Symbol).Cmp(utils.D("0")) <= 0 {
			utils.ResponseFailJson(ctx, "可用金额不足")
			return
		}
	}

	if req.OrderType == trading_core.OrderTypeLimit {
		info, err = orders.NewLimitOrder(user_id, req.Symbol, req.Side, req.Price, req.Quantity)
	} else if req.OrderType == trading_core.OrderTypeMarket {
		if utils.D(req.Amount).Cmp(decimal.Zero) > 0 {
			info, err = orders.NewMarketOrderByAmount(user_id, req.Symbol, req.Side, req.Amount)
		} else if utils.D(req.Quantity).Cmp(decimal.Zero) > 0 {
			info, err = orders.NewMarketOrderByQty(user_id, req.Symbol, req.Side, req.Quantity)
		}
	}

	if err != nil {
		utils.ResponseFailJson(ctx, err.Error())
		return
	}
	utils.ResponseOkJson(ctx, info.OrderId)
}
