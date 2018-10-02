package main

import (
	"bytes"
	"encoding/base64"
	"image"
	"log"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"image/jpeg"
	"image/png"
	"net/http"
	"qrudicon/lib"
)

const (
	DefaultSize = 1024
	MaxSize     = 4092
)

type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (Response, error) {
	img := lib.NewSimpleQrudicon(getText(&request), getImageSize(&request))

	b64Encoded, contentType := encode(&request, img)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: true,
		Body:            b64Encoded,
		Headers: map[string]string{
			"Content-Type":     contentType,
			"Content-Encoding": "base64",
		},
	}

	return resp, nil
}

func getImageSize(request *events.APIGatewayProxyRequest) uint {
	rawSize := request.QueryStringParameters["s"]

	if rawSize == "" {
		return DefaultSize
	}

	size, err := strconv.ParseInt(rawSize, 10, 64)
	if err != nil {
		log.Printf("Could not parse '%s' to integer", rawSize)
		return DefaultSize
	}

	if size > MaxSize {
		log.Printf("Given size %d is higher than max size %d", size, MaxSize)
		return DefaultSize
	}

	return uint(size)
}

func getText(request *events.APIGatewayProxyRequest) string {
	switch request.HTTPMethod {
	case http.MethodGet:
		return request.QueryStringParameters["t"]
	case http.MethodPost:
		return request.Body
	}

	return ""
}

func encode(request *events.APIGatewayProxyRequest, img image.Image) (string, string) {
	imageType := request.QueryStringParameters["e"]

	buf := &bytes.Buffer{}
	var contentType string

	switch strings.ToLower(imageType) {
	case "jpeg":
		fallthrough
	case "jpg":
		jpeg.Encode(buf, img, nil)
		contentType = "image/jpeg"
	case "png":
		fallthrough
	default:
		png.Encode(buf, img)
		contentType = "image/png"
	}

	b64 := base64.StdEncoding.EncodeToString(buf.Bytes())
	return b64, contentType
}

func main() {
	lambda.Start(Handler)
}
