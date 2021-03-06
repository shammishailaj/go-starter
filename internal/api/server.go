package api

import (
	"context"
	"database/sql"
	"errors"

	"allaboutapps.dev/aw/go-starter/internal/mailer"
	"allaboutapps.dev/aw/go-starter/internal/mailer/transport"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type Router struct {
	Root       *echo.Group
	APIV1Auth  *echo.Group
	APIV1Users *echo.Group
}

type Server struct {
	Config ServerConfig
	DB     *sql.DB
	Echo   *echo.Echo
	Router *Router
	Mailer *mailer.Mailer
}

func NewServer(config ServerConfig) *Server {
	s := &Server{
		Config: config,
		DB:     nil,
		Echo:   nil,
		Router: nil,
		Mailer: nil,
	}

	return s
}

func (s *Server) InitDB(ctx context.Context) error {
	db, err := sql.Open("postgres", s.Config.Database.ConnectionString())
	if err != nil {
		return err
	}

	if err := db.PingContext(ctx); err != nil {
		return err
	}

	s.DB = db

	return nil
}

func (s *Server) InitMailer(mock ...bool) error {
	if len(mock) > 0 && mock[0] {
		log.Warn().Msg("Initializing mock mailer")
		s.Mailer = mailer.New(s.Config.Mailer, transport.NewMock())
	} else {
		s.Mailer = mailer.New(s.Config.Mailer, transport.NewSMTP(s.Config.SMTP))
	}

	return s.Mailer.ParseTemplates()
}

func (s *Server) Ready() bool {
	return s.DB != nil && s.Echo != nil && s.Router != nil && s.Mailer != nil
}

func (s *Server) Start() error {
	if !s.Ready() {
		return errors.New("server is not ready")
	}

	return s.Echo.Start(s.Config.Echo.ListenAddress)
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Info().Msg("Shutting down server")

	if s.DB != nil {
		log.Debug().Msg("Closing database connection")

		if err := s.DB.Close(); err != nil && err != sql.ErrConnDone {
			log.Error().Err(err).Msg("Failed to close database connection")
		}
	}

	log.Debug().Msg("Shutting down echo server")

	return s.Echo.Shutdown(ctx)
}
