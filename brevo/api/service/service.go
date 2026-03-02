package service

import (
	"brevo/common/dto"
	"brevo/common/helper"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SendCreateRequest(g *gin.Context, method, link string, payload *strings.Reader) (*http.Response, error) {

	url := fmt.Sprintf("https://api.brevo.com/v3/%s", link)
	fmt.Println("URL: ", url)

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		errorRespone := dto.Error{
			Message:        "Failed Creating Request",
			Warn:           "Use correct parameter",
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest}
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": errorRespone})
		return nil, err
	}

	res, err := helper.Sendrequest(g, req)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func SendRequestWithoutPayload(g *gin.Context, method, link string) (*http.Response, error) {

	url := fmt.Sprintf("https://api.brevo.com/v3/%s", link)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		errorRespone := dto.Error{
			Message:        "Failed Creating Request",
			Warn:           "Use correct parameter",
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest}
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": errorRespone})
		return nil, err
	}

	res, err := helper.Sendrequest(g, req)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func SendRequestWithPayload(g *gin.Context, method, link string, payload *strings.Reader, id int) (*http.Response, error) {

	url := fmt.Sprintf("https://api.brevo.com/v3/%s/%d", link, id)

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		errorRespone := dto.Error{
			Message:        "Failed Creating Request",
			Warn:           "Use correct parameter",
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest}
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": errorRespone})
		return nil, err
	}

	res, err := helper.Sendrequest(g, req)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func SendRequestById(g *gin.Context, method, link string, id int) (*http.Response, error) {

	url := fmt.Sprintf("https://api.brevo.com/v3/%s/%d", link, id)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		errorRespone := dto.Error{
			Message:        "Failed Creating Request",
			Warn:           "Use correct parameter",
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest}
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": errorRespone})
		return nil, err
	}

	res, err := helper.Sendrequest(g, req)
	if err != nil {
		return nil, err
	}

	return res, nil

}
