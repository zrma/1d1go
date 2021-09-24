package main

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*20)
	defer cancel()

	logger, _ := zap.NewProduction()
	defer func() { _ = logger.Sync() }()

	for ctx.Err() == nil {
		const targetUrl = "https://{target endpoint}"

		// GET 호출
		_, err := http.Get(targetUrl)
		if err != nil {
			logger.Error("fail",
				zap.Error(err),
			)
		}
	}
}
