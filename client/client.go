package main

import (
	"carModelIdf/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverIP = "127.0.0.1"
var serverPort = "7180"

func main() {
	conn, err := grpc.Dial(serverIP + ":" + serverPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewCarModelIdfClient(conn)
	req := proto.IdfRequest{
		ImgType:       "url",
		ImgUrlOrBytes: "https://club2.autoimg.cn/album/g20/M0F/5C/39/userphotos/2021/04/23/13/820_ChwEpmCCV4aAQGlcAAPT2PwwsJo731.jpg",
	}
	res, err := client.Identify(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	fmt.Println("---- 车型识别结果 ----")
	fmt.Println("编号：", res.Id)
	fmt.Println("车型名称：" + res.Model.Name)
	fmt.Println("开售年份：" + res.Model.Year)
	fmt.Println("车身颜色：" + res.Color)
	fmt.Println("车身尺寸：宽", res.Size.Width, "，高", res.Size.Height, "，左距", res.Size.Left, "，上距", res.Size.Top)
	fmt.Println("识别置信度：", res.Model.Score)
}
