package models

import (
	"encoding/json"
	"fmt"
	"github.com/songyzh/blog/configs"
	"github.com/songyzh/blog/constants"
	"time"
)

type Tag struct {
	ID   int
	Text string
}


func GetTagByID(id int)(Tag, error){
	tag := Tag{}
	key := fmt.Sprintf("%s%v", constants.CACHE_TAG_BY_ID, id)
	cached, err := configs.Cache.Get(key).Result()
	if err == nil{
		err := json.Unmarshal([]byte(cached), &tag)
		return tag, err
	}
	err = configs.DB.Get(&tag, "select * from tag where id = ?", id)
	if err == nil{
		j, _ := json.Marshal(tag)
		err = configs.Cache.Set(key, string(j), 15 * 24 * time.Hour).Err()
	}
	return tag, err
}

func GetTagByText(text string)(Tag, error){
	tag := Tag{}
	key := fmt.Sprintf("%s%v", constants.CACHE_TAG_BY_TEXT, text)
	cached, err := configs.Cache.Get(key).Result()
	if err == nil{
		err := json.Unmarshal([]byte(cached), &tag)
		return tag, err
	}
	err = configs.DB.Get(&tag, "select * from tag where text = ?", text)
	if err == nil{
		j, _ := json.Marshal(tag)
		configs.Cache.Set(key, string(j), 15 * 24 * time.Hour).Err()
	}
	return tag, err
}

func GetAvailableTags()([]string, error){
	tagTexts := []string{}
	err := configs.DB.Select(&tagTexts, "select t.text from tag t join post_tag pt on t.id = pt.tag_id group by pt.tag_id order by count(pt.tag_id) desc")
	return tagTexts, err
}
