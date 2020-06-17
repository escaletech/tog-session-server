package handler

import (
	"context"
	"log"

	"github.com/escaletech/tog-go/sessions"
	"github.com/escaletech/tog-session-server/internal/config"
	"github.com/gofiber/fiber"
)

func Register(app *fiber.App, conf config.Config) error {
	opt := sessions.ClientOptions{
		Addr:    conf.RedisURL,
		Cluster: conf.RedisCluster,
		OnError: func(ctx context.Context, err error) {
			log.Println(err)
		},
	}
	client, err := sessions.NewClient(context.Background(), opt)
	if err != nil {
		return err
	}

	RegisterWithClient(app, conf, client)

	return nil
}

func RegisterWithClient(app *fiber.App, conf config.Config, client sessionCreator) {
	route := "/:namespace/:session"
	if conf.PathPrefix != "" {
		route = conf.PathPrefix + route
	}

	h := &handler{client}
	app.Get(route, h.Session)
}
