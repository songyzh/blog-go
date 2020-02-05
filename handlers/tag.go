package handlers

import (
	"github.com/go-chi/render"
	"github.com/songyzh/blog/models"
	"github.com/songyzh/blog/types"
	"log"
	"net/http"
)

func GetAvailableTags(w http.ResponseWriter, r *http.Request) {
	tagTexts, err := models.GetAvailableTags()
	if err != nil{
		log.Println(err)
		render.Render(w, r, types.NewResponse([]string{}))
		return
	}
	if err := render.Render(w, r, types.NewResponse(tagTexts)); err != nil {
		render.Render(w, r, types.ErrRender(err))
		return
	}
}

func Benchmark(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, types.NewResponse([]string{}))
	return
}
