/**
 * @Author: DollarKiller
 * @Description: video service
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 13:04 2019-10-24
 */
package service

import (
	"OnedriveAnalysisVideo/datasource/cache"
	"github.com/dollarkillerx/easyutils"
)

type VideoService struct {

}

func (v *VideoService) Get(url string) {
	urlSh256 := easyutils.Sha256Encode(url)
	dataInterface, e := cache.Get(urlSh256)

	dataInterface = dataInterface
	if e != nil {
		// 如果没有被缓存
		// 进行查询

	}

}

// 解析获取url
func (v *VideoService) paUrl(url string) string {
	return ""
}