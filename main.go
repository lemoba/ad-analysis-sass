package main

import (
	"flag"
	"fmt"
	"github.com/lemoba/ad-analysis-sass/pkg/config"
	"github.com/lemoba/ad-analysis-sass/pkg/log"
	"go.uber.org/zap"
)

func main() {
	var envConf = flag.String("conf", "config/local.yaml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()

	conf := config.NewConfig(*envConf)
	logger := log.NewLog(conf)

	logger.Info("server start", zap.String("host", fmt.Sprintf("http://%s:%d", conf.GetString("http.host"), conf.GetInt("http.port"))))
}
