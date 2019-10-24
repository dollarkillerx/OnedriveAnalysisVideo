/**
 * @Author: DollarKiller
 * @Description: this is main router
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:43 2019-10-24
 */
package router

import (
	"OnedriveAnalysisVideo/web/contriller"
	"github.com/dollarkillerx/erguotou"
)

func AppRouter(app *erguotou.Engine) {
	app.Get("/video/:url", contriller.Video)
}
