package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/heejinzzz/carModelIdf/proto"
	"google.golang.org/grpc"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
)

var BaiduUrl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/car?access_token="
var accessToken = "**********************************************************************"
var ip = ""
var port = "7180"

type response struct {
	Id     int64  `json:"log_id"`
	Color  string `json:"color_result"`
	Size   size   `json:"location_result"`
	Result []car  `json:"result"`
}

type size struct {
	Width  int32 `json:"width"`
	Height int32 `json:"height"`
	Left   int32 `json:"left"`
	Top    int32 `json:"top"`
}

type car struct {
	Name  string  `json:"name"`
	Year  string  `json:"year"`
	Score float64 `json:"score"`
}

type carModelIdfServer struct {
}

func (server *carModelIdfServer) Identify(ctx context.Context, request *proto.IdfRequest) (*proto.IdfResponse, error) {
	var data url.Values
	if request.ImgType == "url" || request.ImgType == "URL" {
		data = url.Values{
			"url": []string{request.ImgUrlOrBytes},
		}
	} else if request.ImgType == "bytes" {
		data = url.Values{
			"image": []string{request.ImgUrlOrBytes},
		}
	} else {
		return nil, errors.New("Error: ImgType must be 'url' or 'bytes' ")
	}
	body := strings.NewReader(data.Encode())
	res, err := http.Post(BaiduUrl+accessToken, "application/x-www-form-urlencoded", body)
	if err != nil {
		return nil, err
	}
	resMsg, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result response
	err = json.Unmarshal(resMsg, &result)
	if err != nil {
		return nil, err
	}
	carSize := proto.CarSize{
		Height: result.Size.Height,
		Left:   result.Size.Left,
		Top:    result.Size.Top,
		Width:  result.Size.Width,
	}
	carModel := proto.CarModel{
		Name:  result.Result[0].Name,
		Year:  result.Result[0].Year,
		Score: result.Result[0].Score,
	}
	idfResponse := proto.IdfResponse{
		Id:    result.Id,
		Color: result.Color,
		Size:  &carSize,
		Model: &carModel,
	}
	return &idfResponse, nil
}

func main() {
	lis, err := net.Listen("tcp", ip + ":" + port)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	var server carModelIdfServer
	proto.RegisterCarModelIdfServer(grpcServer, &server)
	grpcServer.Serve(lis)
}
