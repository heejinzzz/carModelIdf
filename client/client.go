package client

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/heejinzzz/carModelIdf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io/ioutil"
)

type carModelIdfClient struct {
	client proto.CarModelIdfClient
}

// NewClient 根据指定的 CarModelIdf Server 的地址创建一个 CarModelIdf Client
func NewClient(serverAddress string) *carModelIdfClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := proto.NewCarModelIdfClient(conn)
	return &carModelIdfClient{
		client: client,
	}
}

// PredictByImgUrl 根据图片的url，获取图片，进行预测
func (client *carModelIdfClient) PredictByImgUrl(url string) error {
	req := proto.IdfRequest{
		ImgType:       "url",
		ImgUrlOrBytes: url,
	}
	res, err := client.client.Identify(context.Background(), &req)
	if err != nil {
		return err
	}
	printResult(res)
	return nil
}

// PredictByImgName 根据图片的本地路径，获取图片，进行预测
func (client *carModelIdfClient) PredictByImgName(imageName string) error {
	bytes, err := ioutil.ReadFile(imageName)
	if err != nil {
		return err
	}
	imageString := base64.StdEncoding.EncodeToString(bytes)
	req := proto.IdfRequest{
		ImgType:       "bytes",
		ImgUrlOrBytes: imageString,
	}
	res, err := client.client.Identify(context.Background(), &req)
	if err != nil {
		return err
	}
	printResult(res)
	return nil
}

func printResult(res *proto.IdfResponse) {
	fmt.Println("---- 车型识别结果 ----")
	fmt.Println("编号：", res.Id)
	fmt.Println("车型名称：" + res.Model.Name)
	fmt.Println("开售年份：" + res.Model.Year)
	fmt.Println("车身颜色：" + res.Color)
	fmt.Println("车身尺寸：宽", res.Size.Width, "，高", res.Size.Height, "，左距", res.Size.Left, "，上距", res.Size.Top)
	fmt.Println("识别置信度：", res.Model.Score)
}
