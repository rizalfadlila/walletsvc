package router

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/julo/walletsvc/internal/infrastructure/dto"
	"github.com/rs/zerolog/log"
	"time"
)

const (
	defaultIdleTimeout  = time.Second * 15
	defaultReadTimeout  = time.Second * 15
	defaultWriteTimeout = time.Second * 15
)

type (
	MyRouter struct {
		port   int
		app    *fiber.App
		router fiber.Router
	}

	Config struct {
		Port         int    `env:"PORT"`
		IdleTimeout  string `env:"IDLE_TIMEOUT"`
		ReadTimeout  string `env:"READ_TIMEOUT"`
		WriteTimeout string `env:"WRITE_TIMEOUT"`
		BodyLimit    int    `env:"BODY_LIMIT"`
		Prefix       string `env:"PREFIX"`
	}

	Handler func(ctx *fiber.Ctx) *dto.JSONResponse
)

func New(cfg Config) (*MyRouter, error) {
	fiberCfg, err := parseConfig(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to init router: %v", err)
	}

	app := fiber.New(*fiberCfg)
	myRouter := app.Use(cors.New(), recovery())

	if cfg.Prefix != "" {
		myRouter = myRouter.Group(cfg.Prefix)
	}

	m := &MyRouter{
		port:   cfg.Port,
		app:    app,
		router: myRouter,
	}

	return m, nil
}
func parseConfig(cfg Config) (*fiber.Config, error) {
	cfgDefault := &fiber.Config{
		JSONDecoder:           json.Unmarshal,
		JSONEncoder:           json.Marshal,
		WriteTimeout:          defaultWriteTimeout,
		IdleTimeout:           defaultIdleTimeout,
		ReadTimeout:           defaultReadTimeout,
		DisableStartupMessage: true,
	}

	if cfg.IdleTimeout != "" {
		idleTimeout, err := time.ParseDuration(cfg.IdleTimeout)
		if err != nil {
			return nil, fmt.Errorf("failed parse idle timeout duration: %v", idleTimeout)
		}
		cfgDefault.IdleTimeout = idleTimeout
	}

	if cfg.ReadTimeout != "" {
		readTimeout, err := time.ParseDuration(cfg.ReadTimeout)
		if err != nil {
			return nil, fmt.Errorf("failed parse read timeout duration: %v", readTimeout)
		}
		cfgDefault.ReadTimeout = readTimeout
	}

	if cfg.WriteTimeout != "" {
		writeTimeout, err := time.ParseDuration(cfg.WriteTimeout)
		if err != nil {
			return nil, fmt.Errorf("failed parse write timeout duration: %v", writeTimeout)
		}
		cfgDefault.WriteTimeout = writeTimeout
	}

	if cfg.BodyLimit != 0 {
		cfgDefault.BodyLimit = cfg.BodyLimit * 1024 * 1024
	}

	return cfgDefault, nil
}

func (r *MyRouter) Start() error {
	log.Info().Msg(fmt.Sprintf("API Listening on :%v", r.port))
	return r.app.Listen(fmt.Sprintf(":%v", r.port))
}

func (r *MyRouter) Shutdown(ctx context.Context) error {
	return r.app.ShutdownWithContext(ctx)
}
