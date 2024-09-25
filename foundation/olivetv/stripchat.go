package olivetv

import (
	"fmt"
	"regexp"
	"time"

	"github.com/parnurzeal/gorequest"
	"github.com/tidwall/gjson"
)

func get_modelId(modleName string, daili string) string {

	// 创建一个新的 Request 对象
	request := gorequest.New()

	if daili != "" {
		request = request.Proxy(daili) //代理
	}

	// 添加头部信息
	request.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	request.Set("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	request.Set("Accept-Encoding", "gzip, deflate")
	request.Set("Upgrade-Insecure-Requests", "1")
	request.Set("Sec-Fetch-Dest", "document")
	request.Set("Sec-Fetch-Mode", "navigate")
	request.Set("Sec-Fetch-Site", "none")
	request.Set("Sec-Fetch-User", "?1")
	request.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:124.0) Gecko/20100101 Firefox/124.0")
	// request.Set("If-Modified-Since", "Mon, 29 Jul 2024 08:41:12 GMT")
	request.Set("Te", "trailers")
	request.Set("Connection", "close")

	// 发起 GET 请求
	_, body, errs := request.Get("https://zh.stripchat.com/api/front/v2/models/username/" + modleName + "/chat").End()

	// 处理响应
	if len(errs) > 0 {
		fmt.Println("请求modelID出错:", body, errs)
		return "false"
	} else {
		// 解析 JSON 响应
		if (len(gjson.Get(body, "messages").String())) > 2 {
			modelId := gjson.Get(body, "messages.0.modelId").String()
			return modelId
		} else {
			return "OffLine"
		}
	}
}

func get_M3u8(modelId string, daili string) string {
	if modelId == "false" || modelId == "OffLine" {
		return "false"
	}
	// url := "https://edge-hls.doppiocdn.com/hls/" + modelId + "/master/" + modelId + "_auto.m3u8?playlistType=lowLatency"
	url := "https://edge-hls.doppiocdn.com/hls/" + modelId + "/master/" + modelId + "_auto.m3u8?playlistType=standard"
	// url := "https://edge-hls.doppiocdn.com/hls/" + modelId + "/master/" + modelId + ".m3u8"
	request := gorequest.New()
	if daili != "" {
		request = request.Proxy(daili) //代理
	}
	resp, body, errs := request.Get(url).End()

	if len(errs) > 0 || resp.StatusCode != 200 {
		return "false"
	} else {
		// fmt.Println((body))
		// re := regexp.MustCompile(`(https:\/\/[\w\-\.]+\/hls\/[\d]+\/[\d\_p]+\.m3u8\?playlistType=lowLatency)`)
		re := regexp.MustCompile(`(https:\/\/[\w\-\.]+\/hls\/[\d]+\/[\d\_p]+\.m3u8\?playlistType=standard)`) //等价于\?playlistType=standard
		matches := re.FindString(body)
		return matches
	}
}
func test_m3u8(url string, daili string) bool {
	if url == "false" || url == "" {
		return false
	} else {
		request := gorequest.New()
		if daili != "" {
			request = request.Proxy(daili) //代理
		}
		resp, body, errs := request.Get(url).End()
		if len(errs) > 0 || resp.StatusCode != 200 {
			return false
		}
		if resp.StatusCode == 200 { //403代表开票，普通用户无法查看，只能看大厅表演
			_ = body
			return true
		}
		return false
	}
}
func init() {
	registerSite("stripchat", &stripchat{})
}

type stripchat struct {
	base
}

func (this *stripchat) Name() string {
	return "stripchat"
}

func (this *stripchat) Snap(tv *TV) error {
	tv.Info = &Info{
		Timestamp: time.Now().Unix(),
	}
	return this.set(tv)
}

func (this *stripchat) set(tv *TV) error {
	// fmt.Println("roomID:", tv.RoomID, "\ncookie:", tv.cookie, "\ntv:\n", tv)
	modelName := tv.RoomID
	proxy := "http://127.0.0.1:7890"
	modelID := get_modelId(modelName, proxy)
	m3u8 := get_M3u8(modelID, proxy)
	m3u8Status := test_m3u8(m3u8, proxy)
	if modelID == "false" {
		return nil
	}
	if (modelID == "OffLine") || (m3u8 == "false") {
		tv.roomName = modelName
		tv.streamerName = modelID
		tv.roomOn = false
		// tv.streamURL = ""
		return nil
	}

	if m3u8 != "false" {
		tv.roomName = modelName
		tv.streamerName = modelID
		tv.roomOn = m3u8Status
		tv.streamURL = m3u8

		return nil
	}

	return nil
}
