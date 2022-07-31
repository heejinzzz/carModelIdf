# carModelIdf
Model recognition microservice implemented by calling Baidu api, based on grpc.  

调用百度api实现的车型识别微服务，基于grpc。

----
## 一、CarModelIdf Server 部署（Only for Linux）

### 1. 常规部署
确保已安装 golang 环境后，直接执行部署脚本：
 
    bash ServerDeploy.sh -a <access_token> -h <ip> -p <port>

参数说明：

access_token 是你在百度云开放平台获取的 access_token，请参考[access_token 获取](https://ai.baidu.com/ai-doc/REFERENCE/Ck3dwjhhu)。无默认值，必须输入。

ip 是你指定的 server 所要部署在的 ip 地址，默认值为 localhost。

port 是你指定的 server 所要部署在的 端口号，默认值为 7180。

也支持长选项：

    bash ServerDeploy.sh --access_token <access_token> --ip <ip> --port <port>
    
### 2. docker容器式部署
[docker: heejinzzz/car-model-idf](https://hub.docker.com/repository/docker/heejinzzz/car-model-idf)

----

## 二、CarModelIdf Client 使用
Server 部署完成后，客户端方在 client/client.go 文件中修改 serverIP、serverPort，并将 req 中的 ImgUrlOrBytes 修改为想要识别的图片的 url 即可（也可以对本地图片进行识别，将其按 base64 编码为string字符串作为 req 中的 ImgUrlOrBytes ，并修改 req 的 ImgType 为 "bytes" ）。然后执行：

    go run client/client.go
    
即可获取 CarModelIdf Server 对车型的预测结果。
