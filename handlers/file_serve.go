package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/riyadennis/redis-wrapper"
	"fmt"
)

func FileServe(w http.ResponseWriter, res *http.Request, _ httprouter.Params){
	c := redis_wrapper.Client{}
	client, err := c.Create()
	if err != nil {
		logrus.Error(err.Error())
	}
	var keys []string
	var cursor uint64
	keys, cursor, err = client.RedisClient.Scan(cursor, "", 100).Result()
	detail := "{"
	for _, k := range keys {
		val, err := client.Get(k)
		if err != nil {
			logrus.Error(err.Error())
		}
		detail += fmt.Sprintf("%s \n", val)
	}
	detail += "}"
	api := createResponse(detail, "success", http.StatusOK)
	jsonResponseDecorator(api, w)
}
