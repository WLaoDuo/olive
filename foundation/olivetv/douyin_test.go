package olivetv_test

import (
	"testing"

	"github.com/WLaoDuo/olive/foundation/olivetv"
)

func TestDouyin_Snap(t *testing.T) {
	u := "https://live.douyin.com/152686547303"
	dy, err := olivetv.NewWithURL(u)
	if err != nil {
		println(err.Error())
		return
	}
	dy.Snap()
	t.Log(dy)
}
