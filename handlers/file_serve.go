package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/riyadennis/redis-wrapper"
	"encoding/json"
)

type Files struct {
	Files []FileResponse `json: "files"`
}
type FileResponse struct {
	Filename string `json: "file_name"`
}

func FileServe(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	c := redis_wrapper.Client{}
	res := make([]FileResponse, 1)
	client, err := c.Create()
	if err != nil {
		logrus.Error(err.Error())
	}
	keys := client.RedisClient.Keys("*")
	for _, k := range keys.Val() {
		val, err := client.Get(k)
		if err != nil {
			logrus.Error(err.Error())
		}
		if val != ""{
			fr := FileResponse{val}
			res = append(res, fr)
		}
	}
	files := Files{Files: res}
	fileJson, err := json.Marshal(files)
	detail := string(fileJson)
	api := createResponse(detail, "success", http.StatusOK)
	jsonResponseDecorator(api, w)
}
