package middlewares

import (
	"github.com/go-chi/render"
	"github.com/songyzh/blog/types"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
)

func Paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pagination := make(map[string]int)
		for _, k := range []string{"page", "size"}{
			v := r.URL.Query().Get(k)
			vInt, err := strconv.Atoi(v)
			if err != nil{
				render.Render(w, r, types.ErrInvalidRequest(err))
				return
			}
			pagination[k] = vInt
		}
		pagination["offset"] = (pagination["page"] - 1) * pagination["size"]
		pagination["limit"] = pagination["size"]
		delete(pagination, "page")
		delete(pagination, "size")
		ctx := context.WithValue(r.Context(), "pagination", pagination)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
