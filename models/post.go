package models

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/songyzh/blog/configs"
	"github.com/songyzh/blog/constants"
	"time"
)


type Post struct {
	Content   string
	Cover     string
	CreatedAt time.Time
	FamilyID  int
	ID        int
	Slug      string
	Status    int
	Title     string
	UpdatedAt time.Time
}

func GetPostBySlug(slug string) (Post, error){
	post := Post{}
	key := fmt.Sprintf("%s%v", constants.CACHE_POST_BY_SLUG, slug)
	cached, err := configs.Cache.Get(key).Result()
	if err == nil{
		err := json.Unmarshal([]byte(cached), &post)
		return post, err
	}
	err = configs.DB.Get(&post, "select * from post where slug = ? and status = ?", slug, constants.ENABLED)
	if err == nil{
		j, _ := json.Marshal(post)
		configs.Cache.Set(key, string(j), 15 * 24 * time.Hour).Err()
	}
	return post, err
}

func MGetPostsByIds(ids []int, offset int, limit int) []Post{
	posts := []Post{}
	if len(ids) == 0{
		return posts
	}
	query, args, _ := sqlx.In(
		"select * from post where id in (?) and status = ? order by id desc limit ?,?",
		ids, constants.ENABLED, offset, limit)
	configs.DB.Select(&posts, query, args...)
	return posts
}

func CountPostsByIds(ids []int)int{
	var cnt int
	if len(ids) == 0{
		return cnt
	}
	query, args, _ := sqlx.In(
		"select count(1) as cnt from post where id in (?) and status = ?",
		ids, constants.ENABLED)
	configs.DB.Get(&cnt, query, args...)
	return cnt
}

func MGetPosts(offset int, limit int) []Post{
	posts := []Post{}
	query, args, _ := sqlx.In(
		"select * from post where status = ? order by id desc limit ?,?",
		constants.ENABLED, offset, limit)
	configs.DB.Select(&posts, query, args...)
	return posts
}

func CountPosts() int{
	var cnt int
	query, args, _ := sqlx.In(
		"select count(1) as cnt from post where status = ?",
		constants.ENABLED)
	configs.DB.Get(&cnt, query, args...)
	return cnt
}
