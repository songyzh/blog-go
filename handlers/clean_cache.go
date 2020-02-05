package handlers

import (
	"github.com/go-chi/render"
	"github.com/songyzh/blog/configs"
	"github.com/songyzh/blog/types"
	"net/http"
)


func CleanCache(w http.ResponseWriter, r *http.Request) {
	configs.Cache.FlushDB().Err()
	response := types.NewResponse(nil)
	if err := render.Render(w, r, response); err != nil {
		render.Render(w, r, types.ErrRender(err))
		return
	}
}
