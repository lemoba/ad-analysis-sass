package main

import (
	"context"
	"flag"
	"fmt"

	"go.uber.org/zap"

	"github.com/lemoba/ad-analysis-sass/internal/server"
	"github.com/lemoba/ad-analysis-sass/pkg/app"
	"github.com/lemoba/ad-analysis-sass/pkg/config"
	"github.com/lemoba/ad-analysis-sass/pkg/log"
)

func main() {
	var envConf = flag.String("conf", "config/local.yaml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()

	conf := config.NewConfig(*envConf)
	logger := log.NewLog(conf)

	logger.Info("server start", zap.String("host", fmt.Sprintf("http://%s:%d", conf.GetString("http.host"), conf.GetInt("http.port"))))

	httpServer := server.NewHttpServer(logger, conf)

	app := app.NewApp(app.WithServer(httpServer), app.WithName("ad-sass-server"))
	if err := app.Run(context.Background()); err != nil {
		panic(err)
	}
}
