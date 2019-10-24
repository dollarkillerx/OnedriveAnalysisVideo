/**
 * @Author: DollarKiller
 * @Description: video
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:44 2019-10-24
 */
package contriller

import "github.com/dollarkillerx/erguotou"

func Video(ctx *erguotou.Context) {
	url, b := ctx.PathValueString("url")
	if !b {
		// 如果用户没有传入参数
		ctx.Redirect(302, "https://www.baidu.com")
		return
	}
	// 获取真实url地址

}
