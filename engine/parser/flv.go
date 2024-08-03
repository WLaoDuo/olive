package parser

import (
	l "github.com/WLaoDuo/olive/engine/log"
	"github.com/go-olive/flv"
	"github.com/sirupsen/logrus"
)

func init() {
	SharedManager.Register(
		new(customFlv),
	)
}

type customFlv struct {
	*flv.Parser
}

func (this *customFlv) New() Parser {
	return &customFlv{
		Parser: flv.NewParser(),
	}
}

func (this *customFlv) Parse(streamURL string, out string) (err error) {
	l.Logger.WithFields(logrus.Fields{
		// "streamURL": streamURL,
		"out": out,
	}).Debug("flv working")

	return this.Parser.Parse(streamURL, out)
}
