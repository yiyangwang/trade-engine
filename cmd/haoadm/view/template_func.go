package view

import (
	"fmt"
	"html/template"
	"time"

	"github.com/yzimhao/trading_engine/utils/app"
	"github.com/yzimhao/trading_engine/utils/app/config"
)

func templateFuncMap() template.FuncMap {
	return template.FuncMap{
		"app_power_by": func() string {
			return fmt.Sprintf("powered by HaoTrader %s", app.Version)
		},
		"app_short_name": func() string {
			return "HT"
		},
		"version": func() string {
			return app.Version
		},
		"author": func() string {
			return "yzimhao"
		},
		"author_link": func() string {
			return "https://github.com/yzimhao"
		},
		"year": func() string {
			return time.Now().Format("2006")
		},
		"unsafe": func(s string) template.HTML {
			return template.HTML(s)
		},
		"site_name": func() string {
			return config.App.Haoadm.SiteName
		},
	}
}
