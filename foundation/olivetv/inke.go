package olivetv

import (
	"time"

	"github.com/WLaoDuo/olive/foundation/olivetv/model"
	"github.com/WLaoDuo/olive/foundation/olivetv/util"
)

func init() {
	registerSite("inke", &inke{})
}

type inke struct {
	base
}

func (this *inke) Name() string {
	return "映客"
}

func (this *inke) Snap(tv *TV) error {
	tv.Info = &Info{
		Timestamp: time.Now().Unix(),
	}
	return this.set(tv)
}

func (this *inke) set(tv *TV) error {
	a := new(model.InkeAutoGenerated)
	req := &util.HttpRequest{
		URL:          "https://webapi.busi.inke.cn/web/live_share_pc?uid=" + tv.RoomID,
		Method:       "GET",
		ResponseData: a,
		ContentType:  "application/json",
		Header: map[string]string{
			"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.71 Safari/537.36 Edg/94.0.992.38",
		},
	}
	if err := req.Send(); err != nil {
		return err
	}
	tv.roomName = a.Data.LiveName
	tv.streamerName = a.Data.MediaInfo.Nick
	if len(a.Data.LiveAddr) > 0 {
		tv.streamURL = a.Data.LiveAddr[0].StreamAddr
		tv.roomOn = true
	}

	return nil
}
