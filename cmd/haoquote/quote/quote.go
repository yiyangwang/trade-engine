package quote

import (
	"github.com/gookit/goutil/arrutil"
	"github.com/yzimhao/trading_engine/cmd/haobase/base"
	"github.com/yzimhao/trading_engine/cmd/haoquote/quote/tradelog"
	"github.com/yzimhao/trading_engine/cmd/haoquote/quote/www"
	"github.com/yzimhao/trading_engine/utils/app/config"
	"github.com/yzimhao/trading_engine/utils/filecache"
)

func Run() {
	init_symbols_quote()
	www.Run()
}

func init_symbols_quote() {
	filecache.NewStorage(config.App.Haoquote.Cache, 1)

	local_config_symbols := config.App.Local.Symbols
	db_symbols := base.NewTSymbols().All()
	for _, item := range db_symbols {
		if len(local_config_symbols) > 0 && arrutil.Contains(local_config_symbols, item.Symbol) || len(local_config_symbols) == 0 {
			go tradelog.Monitor(item.Symbol, int64(item.PricePrecision), int64(item.QtyPrecision))
		}
	}
}
