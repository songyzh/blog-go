package handlers

import (
	"github.com/fatih/structs"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/songyzh/blog/models"
	"github.com/songyzh/blog/types"
	"log"
	"net/http"
)

func MGetPosts(w http.ResponseWriter, r *http.Request) {
	pagination := r.Context().Value("pagination").(map[string]int)
	data := map[string]interface{}{}
	posts := []map[string]interface{}{}
	var total int
	if tagText := r.URL.Query().Get("tag"); tagText == ""{
		posts = posts4Response(models.MGetPosts(pagination["offset"], pagination["limit"]))
		total = models.CountPosts();
	}else{
		if tag, err := models.GetTagByText(tagText); err == nil{
			postIds := models.MGetPostIDsByTagID(tag.ID)
			posts = posts4Response(models.MGetPostsByIds(postIds, pagination["offset"], pagination["limit"]))
			total = models.CountPostsByIds(postIds);
		}
	}
	data["posts"] = posts
	data["total"] = total
	totalPage := total / pagination["limit"]
	if total % pagination["limit"] != 0{
		totalPage++
	}
	data["totalPage"] = totalPage
	if err := render.Render(w, r, types.NewResponse(data)); err != nil {
		render.Render(w, r, types.ErrRender(err))
		return
	}
}

func GetPostBySlug(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	post, err := models.GetPostBySlug(slug)
	if err != nil{
		log.Println(err)
		render.Render(w, r, types.ErrNotFound)
		return
	}
	if err := render.Render(w, r, types.NewResponse(post4Response(post))); err != nil {
		render.Render(w, r, types.ErrRender(err))
		return
	}
}

func post4Response(post models.Post)map[string]interface{}{
	ret := structs.Map(post)
	tags := []string{}
	ret["tags"] = &tags
	for _, tagID := range models.MGetTagIDsByPostID(post.ID){
		tag, err := models.GetTagByID(tagID)
		if err != nil{
			log.Println(err)
			continue
		}
		tags = append(tags, tag.Text)
	}
	family, err := models.GetFamilyByID(post.FamilyID)
	if err != nil{
		log.Println(err)
		return ret
	}
	ret["family"] = family
	return ret
}

func posts4Response(posts []models.Post)[]map[string]interface{}{
	ret := []map[string]interface{}{}
	for _, post := range posts{
		ret = append(ret, post4Response(post))
	}
	return ret
}
