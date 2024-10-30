package server

import (
	"github.com/spf13/viper"

	"github.com/lemoba/ad-analysis-sass/pkg/log"
	"github.com/lemoba/ad-analysis-sass/pkg/server/http"
)

func NewHttpServer(
	logger *log.Logger,
	conf *viper.Viper,
) *http.Server {
	s := http.NewServer(
		logger,
		http.WithServerHost(conf.GetString("http.host")),
		http.WithServerPort(conf.GetInt("http.port")),
	)
	return s
}
