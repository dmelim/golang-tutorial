package main

import (
	"testing"

	"github.com/dmelim/bookings/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// nothing
	default:
		t.Errorf("type is not *chi.Mux, type is %T", v)
	}
}
