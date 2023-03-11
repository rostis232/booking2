package handlers

import (
	"net/http"
	"github.com/rostis232/booking2/internal/service"
)

func Home(w http.ResponseWriter, r *http.Request) {
	service.RenderTamplate(w, "home.page")
}

func About(w http.ResponseWriter, r *http.Request) {
	service.RenderTamplate(w, "about.page")
}