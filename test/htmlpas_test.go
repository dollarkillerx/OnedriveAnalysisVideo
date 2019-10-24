/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 13:36 2019-10-24
 */
package test

import (
	"github.com/dollarkillerx/easyutils"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPas(t *testing.T) {
	psHtml("https://vipmail-my.sharepoint.cn/:v:/g/personal/dollarkiller_e1_tn/ETX644n2ap1GsZAovwE03pgBJay_mzDDf-75E9bYO3_LsQ?e=UWU8kX")
}

func psHtml(url string) {

	var httpClient *http.Client
	request, e := http.NewRequest("GET", url, nil)
	if e != nil {
		panic(e)
	}
	if request != nil {
		defer request.Body.Close()
	}

	request.Header.Set("User-Agent", easyutils.ReptileGetSpiderAgent())
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Set("Accept-Encoding", "gzip, deflate, br")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	request.Header.Set("Cache-Control", "no-cache")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Host", "vipmail-my.sharepoint.cn")
	request.Header.Set("Pragma", "no-cache")
	request.Header.Set("Upgrade-Insecure-Requests", "1")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:69.0) Gecko/20100101 Firefox/69.0")

	response, e := httpClient.Do(request)
	if response.Body != nil {
		defer response.Body.Close()
	}
	bytes, e := ioutil.ReadAll(response.Body)
	if e != nil {
		panic(e)
	}

	ioutil.WriteFile("test.html", bytes, 00666)
}
