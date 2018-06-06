package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/riyadennis/redis-wrapper"
	"encoding/json"
	"regexp"
	"strings"
)

type Files struct {
	Files []FileResponse `json: "files"`
}
type FileResponse struct {
	Filename string `json: "filename"`
}

func FileServe(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var keys []string
	var cursor uint64
	c := redis_wrapper.Client{}
	res := make([]FileResponse, 1)
	client, err := c.Create()
	if err != nil {
		logrus.Error(err.Error())
	}
	keys, cursor, err = client.RedisClient.Scan(cursor, "", 100).Result()
	for _, k := range keys {
		val, err := client.Get(k)
		if err != nil {
			logrus.Error(err.Error())
		}
		fr := FileResponse{val}
		res = append(res, fr)
	}
	files := Files{Files: res}
	fileJson, err := json.MarshalIndent(files, " ", " ")
	detail := cleanResponse(string(fileJson))
	api := createResponse(detail, "success", http.StatusOK)
	jsonResponseDecorator(api, w)
}
func cleanResponse(jsonString string) string {
	reg := regexp.MustCompile("/(\r\n)+|\r+|\n+|\t+/i")
	return strings.Replace(reg.ReplaceAllString(jsonString, " "), "\\", " ", -1)
}
