/**
 * @Author: DollarKiller
 * @Description: Main
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:41 2019-10-24
 */
package main

import (
	"OnedriveAnalysisVideo/config"
	"OnedriveAnalysisVideo/web/router"
	"github.com/dollarkillerx/erguotou"
)

func main() {
	app := erguotou.New()

	// 注册主路由
	router.AppRouter(app)

	err := app.Run(
		erguotou.SetHost(config.MyConfig.App.Host),
		erguotou.SetDebug(config.MyConfig.App.Debug),
	)
	if err != nil {
		panic(err)
	}

}
