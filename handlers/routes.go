package handlers

import (
	"github.com/julienschmidt/httprouter"
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/sirupsen/logrus"
)

type ApiResponse struct {
	Status int
	Detail string
	Title  string
}

func Run(port string) {
	route := httprouter.New()
	route.GET("/files", FileServe)
	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Listenning to port %s \n", addr)
	logrus.Fatal(http.ListenAndServe(addr, route))
}

func jsonResponseDecorator(response *ApiResponse, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), response.Status)
		return
	}
}
func createResponse(detail, title string, status int) *ApiResponse {
	return &ApiResponse{
		Status: status,
		Detail: detail,
		Title:  title,
	}
}
