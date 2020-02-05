package models

import (
	"encoding/json"
	"fmt"
	"github.com/songyzh/blog/configs"
	"github.com/songyzh/blog/constants"
	"time"
)

type Family struct {
	Avatar string
	ID     int
	Name   string
}

func GetFamilyByID(id int) (Family, error){
	family := Family{}
	key := fmt.Sprintf("%s%v", constants.CACHE_FAMILY_BY_ID, id)
	cached, err := configs.Cache.Get(key).Result()
	if err == nil{
		err := json.Unmarshal([]byte(cached), &family)
		return family, err
	}
	err = configs.DB.Get(&family, "select * from family where id = ?", id)
	if err == nil{
		j, _ := json.Marshal(family)
		configs.Cache.Set(key, string(j), 15 * 24 * time.Hour).Err()
	}
	return family, err
}
