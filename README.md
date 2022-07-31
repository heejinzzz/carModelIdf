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
先获取 CarModelIdf Client ：

    go get github.com/heejinzzz/carModelIdf/client

函数调用示例：

    import "github.com/heejinzzz/carModelIdf/client"
    
    
    // 根据指定的 CarModelIdf Server 地址，创建一个 CarModelIdf Client
    c := client.NewClient("127.0.0.1:7180")
    
    // 根据图片的url获取图片，进行预测
    c.PredictByImgUrl("https://www.ssfiction.com/wp-content/uploads/2020/08/20200806_5f2c89cba3144.jpg")
    // 根据图片的本地路径获取图片，进行预测
    c.PredictByImgName("./car.jpg")
