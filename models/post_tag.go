package models

import (
	"github.com/songyzh/blog/configs"
)

type PostTag struct {
	ID     int
	PostID int
	TagID  int
}


func MGetTagIDsByPostID(postID int)[]int{
	tagIDs := []int{}
	configs.DB.Select(&tagIDs, "select tag_id from post_tag where post_id = ? order by id asc", postID)
	return tagIDs
}

func MGetPostIDsByTagID(tagID int)[]int{
	postIDs := []int{}
	configs.DB.Select(&postIDs, "select post_id from post_tag where tag_id = ? order by id asc", tagID)
	return postIDs
}
