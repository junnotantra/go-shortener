package api

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	redirectorApi "github.com/junnotantra/go-shortener/internal/api/redirector"
	shortenerApi "github.com/junnotantra/go-shortener/internal/api/shortener"
	statisticApi "github.com/junnotantra/go-shortener/internal/api/statistic"
	"github.com/junnotantra/go-shortener/internal/config"
	shortenerRsc "github.com/junnotantra/go-shortener/internal/resource/shortener"
	statisticRsc "github.com/junnotantra/go-shortener/internal/resource/statistic"
	shortenerSrv "github.com/junnotantra/go-shortener/internal/service/shortener"
	statisticSrv "github.com/junnotantra/go-shortener/internal/service/statistic"
	bolt "go.etcd.io/bbolt"
)

type (
	// Server base struct
	Server struct {
		server       *http.Server
		shortenerSrv *shortenerSrv.Service
		statisticSrv *statisticSrv.Service
	}
)

// Main will prepare all requires services & resource then start the server
func Main(cfg *config.Config) error {
	// open DB
	db, err := bolt.Open(
		cfg.Database.Main.FileName,
		0600,
		&bolt.Options{
			Timeout: time.Duration(cfg.Database.Main.Timeout) * time.Millisecond,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// open statistic DB
	dbStat, err := bolt.Open(
		cfg.Database.Statistic.FileName,
		0600,
		&bolt.Options{
			Timeout: time.Duration(cfg.Database.Main.Timeout) * time.Millisecond,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	defer dbStat.Close()

	// init all service & resource
	shortenerResource := shortenerRsc.New(db)
	shortenerService := shortenerSrv.New(shortenerResource)

	statisticResource := statisticRsc.New(dbStat)
	statisticService := statisticSrv.New(statisticResource)

	// init server
	s := &Server{
		server: &http.Server{
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			Handler:      handler(),
		},
		shortenerSrv: shortenerService,
		statisticSrv: statisticService,
	}

	if err := s.Serve(cfg.Server.HTTPPort); err != nil {
		log.Fatal(err)
	}

	return nil
}

// Serve will run an HTTP server
func (s *Server) Serve(port string) error {
	// init all API
	shortenerApi.Init(s.shortenerSrv)
	redirectorApi.Init(s.shortenerSrv, s.statisticSrv)
	statisticApi.Init(s.shortenerSrv, s.statisticSrv)

	lis, err := net.Listen("tcp4", port)
	if err != nil {
		return err
	}

	return s.server.Serve(lis)
}

// Shutdown will tear down the server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
