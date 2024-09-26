package config

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/pelletier/go-toml/v2"
)

//	type Config struct {
//		PortalUsername    string
//		LogDir            string
//		OutTmpl           string
//		BiliupEnable      bool
//		Threads           int
//		MaxBytesPerSecond float64
//	}

type Show struct {
	ID           string
	Enable       bool
	Platform     string `toml:"Platform"`
	StreamerName string `toml:"StreamerName"`
	RoomID       string
	OutTmpl      string
	Parser       string
	SaveDir      string
	PostCmds     string
	SplitRule    string
	DateCreated  time.Time
	DateUpdated  time.Time
}
type AppConfig struct {
	Config Config `toml:"Config"`
	Shows  []Show `toml:"Shows"`
}

func ReadConfigWithBytes(b []byte) (*AppConfig, error) {
	var config AppConfig
	if err := toml.Unmarshal(b, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func ReadConfigWithFile(file string) (*AppConfig, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("can`t open file: %s", file)
	}
	config, err := ReadConfigWithBytes(b)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return config, nil
}

func ReadToml(filepath string) (*AppConfig, error) {
	// test, err := ReadConfigWithFile("D:/study/剪辑-录屏-转码-OCR/go-olive/Lucky-uu.toml")
	test, err := ReadConfigWithFile(filepath)
	if err == nil {
		// fmt.Println(test)
		return test, nil
	} else {
		return nil, err
	}

}
