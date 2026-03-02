package handler

import (
	"brevo/api/service"
	"brevo/common/dto"
	"brevo/common/helper"
	"brevo/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var EmailTemplateBody models.EmailTemplate

var requset dto.Request

func CreateTemplate(g *gin.Context) {

	var err error

	if err := g.ShouldBindJSON(&EmailTemplateBody); err != nil {
		errorRespone := dto.Error{
			Message:        "Invalid Format",
			Warn:           "Use valid format",
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest}
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": errorRespone})
		return
	}

	requset.Body, err = helper.MarshalRequestBody(g, EmailTemplateBody)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": err})
		return
	}
	requset.Method = "POST"
	requset.Url = "smtp/templates"

	res, err := service.SendCreateRequest(g, requset.Method, requset.Url, requset.Body)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": err})
		return
	}

	resBody, err := helper.ValidateRequestBody(g, res)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": err})
		return
	}
	fmt.Println("<=========Successfully Created =========>")
	fmt.Println("Response Body: ", string(resBody))

	g.JSON(http.StatusCreated, string(resBody))

}

func GetTemplate(g *gin.Context) {

	var result models.AllResponse

	requset.Method = "GET"
	requset.Url = "smtp/templates"

	res, err := service.SendRequestWithoutPayload(g, requset.Method, requset.Url)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": err})
		return
	}

	resBody, err := helper.ValidateRequestBody(g, res)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": err})
		return
	}

	fmt.Println("<=========Successfully Receiced =========>")
	fmt.Println("Response Body: ", string(resBody))

	err = json.Unmarshal(resBody, &result)
	if err != nil {
		errorRespone := dto.Error{
			Message:        "Failed UnMarshall",
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest}
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": errorRespone})
		return
	}

	g.JSON(http.StatusOK, result)

}

func GetByIdTemplate(g *gin.Context) {

	var result models.ReceivedData

	requset.Method = "GET"
	requset.Url = "smtp/templates"

	id, err := helper.GetID(g)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": err})
		return
	}

	res, err := service.SendRequestById(g, requset.Method, requset.Url, id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": err})
		return
	}

	resBody, err := helper.ValidateRequestBody(g, res)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": err})
		return
	}

	fmt.Println("<--------Successfully Receiced -------->")
	fmt.Println("Response Body: ", string(resBody))
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		errorRespone := dto.Error{
			Message:        "Failed UnMarshall",
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest}
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": errorRespone})
		return
	}

	g.JSON(http.StatusOK, result)

}

func DeleteTemplate(g *gin.Context) {

	requset.Method = "DELETE"
	requset.Url = "smtp/templates"
	requset.Body = nil

	id, err := helper.GetID(g)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": err})
		return
	}

	res, err := service.SendRequestById(g, requset.Method, requset.Url, id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": err})
		return
	}
	resBody, err := helper.ValidateRequestBody(g, res)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": err})
		return
	}

	fmt.Println("<--------Successfully Deleted -------->")
	fmt.Println("Response Body: ", string(resBody))

	g.JSON(http.StatusOK, gin.H{
		"Message": "Successfully Deleted"})

}

func UpdateTemplate(g *gin.Context) {
	var err error

	if err := g.ShouldBindJSON(&EmailTemplateBody); err != nil {
		errorRespone := dto.Error{
			Message:        "Invalid Format",
			Warn:           "Use valid format",
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest}
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": errorRespone})
		return
	}

	requset.Body, err = helper.MarshalRequestBody(g, EmailTemplateBody)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": err})
		return
	}

	requset.Method = "PUT"
	requset.Url = "smtp/templates"

	id, err := helper.GetID(g)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": err})
		return
	}

	res, err := service.SendRequestWithPayload(g, requset.Method, requset.Url, requset.Body, id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": err})
		return
	}

	resBody, err := helper.ValidateRequestBody(g, res)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": err})
		return
	}

	fmt.Println("<--------Successfully Updated -------->")
	fmt.Println("Response Body: ", string(resBody))

	g.JSON(http.StatusOK, gin.H{
		"Message": "Successfully Updated"})

}
