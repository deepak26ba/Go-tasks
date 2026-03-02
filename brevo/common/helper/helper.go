package helper

import (
	"brevo/common/dto"
	"brevo/internals/config"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Sendrequest(g *gin.Context, req *http.Request) (*http.Response, error) {

	apiKey, err := config.Config()
	if err != nil {
		errorRespone := dto.Error{
			Message:        "Invalid Key",
			Warn:           "Allow Authenticated User",
			Error:          err.Error(),
			HttpStatusCode: http.StatusForbidden}
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": errorRespone})
		return nil, err
	}

	req.Header.Add("api-key", apiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		errorRespone := dto.Error{
			Message:        "Failed to request",
			Warn:           "Invalid path",
			Error:          err.Error(),
			HttpStatusCode: http.StatusNotFound}
		g.JSON(http.StatusServiceUnavailable, gin.H{
			"Error": errorRespone})
		return nil, nil
	}
	return res, nil
}

func GetID(g *gin.Context) (int, error) {

	id := g.Param("id")
	if id == "" {
		g.JSON(http.StatusBadRequest, gin.H{
			"Message": "Invalid Requested id",
			"Error":   "Pass value"})
		return 0, nil
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		errorRespone := dto.Error{
			Message:        "Invalid ID",
			Warn:           "Use valid ID",
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest}
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": errorRespone})
		return 0, err
	}
	return idInt, nil
}

func MarshalRequestBody(g *gin.Context, body any) (*strings.Reader, error) {

	payload, err := json.Marshal(&body)
	if err != nil {
		errorRespone := dto.Error{
			Message:        "Invalid body",
			Warn:           "Use valid body",
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest}
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": errorRespone})
		return nil, err
	}

	return strings.NewReader(string(payload)), nil
}

func ValidateRequestBody(g *gin.Context, res *http.Response) ([]byte, error) {

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		errorRespone := dto.Error{
			Message:        "Failed Reading",
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest}
		g.JSON(http.StatusBadRequest, gin.H{
			"Error": errorRespone})
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode >= 300 {
		g.JSON(res.StatusCode, gin.H{
			"Message": "Error Returned",
			"Error":   string(resBody)})
		return nil, err
	}

	return resBody, nil

}
