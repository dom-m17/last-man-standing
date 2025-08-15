package footballdata

import (
	"log"
	"net/http"

	"github.com/dom-m17/lms/backend/internal/db"
)

type Service struct {
	config     Config
	httpClient HTTPClientInterface
	Querier    db.Querier
}

type HTTPClientInterface interface {
	Do(req *http.Request) (*http.Response, error)
}

func New(db db.Querier) *Service {
	cfg, err := ReadConfig()
	if err != nil {
		log.Fatal("reading config")
	}

	return &Service{
		config:     cfg,
		httpClient: &http.Client{},
		Querier:    db,
	}
}
