package handlers

import (
	"database/sql"
	"net/http"

	"github.com/yoeritjuu/Computer-Store/pkg/config"
	"github.com/yoeritjuu/Computer-Store/pkg/parts"
	"github.com/yoeritjuu/Computer-Store/pkg/repository"
)

type Server struct {
	Cfg config.MySQLConfig

	Sql *sql.DB

	logic parts.PartsService
}

func NewServer(cfg config.MySQLConfig, sql *sql.DB) *Server {
	s := &Server{
		Cfg: cfg,
		Sql: sql,
	}

	repository := repository.NewPartsRepository(s.Sql)
	s.logic = parts.NewPartService(repository)

	return s
}

func (s *Server) Run() {
	http.HandleFunc("/", GetPartsHandler)
	http.ListenAndServe("localhost:8000", nil)
}
