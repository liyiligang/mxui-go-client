<div align=center>
<img src="store\image\logo.svg" width="450" height="280" />
</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.16-blue"/>
<img src="https://img.shields.io/badge/protobuf-3.7.0-green"/>
<img src="https://img.shields.io/badge/grpc-1.38.0-brightgreen"/>
<img src="https://img.shields.io/badge/go--echarts-2.2.4-red"/>
</div>
 <br/> 
 
 [English](./README-en.md) | 简体中文


## 简介
mxui-go-client是 [MXUI](https://github.com/liyiligang/mxui) 的golang客户端

## 查看文档
- [中文](https://mxui-doc.liyiligang.com)    
- [English](https://mxui-doc.liyiligang.com)

## 在线预览
- [MXUI](https://mxui.liyiligang.com)    

## 安装
```bash
go get -u github.com/liyiligang/mxui-go-client
```

## 基本用法
```go
//初始化
client, err := mxui.InitMXUIClient(mxui.ClientConfig{
    Addr:"x.x.x.x:302",
    PublicKeyPath:"./cert/ca_cert.pem",
    CertName: "x.test.example.com",
    NodeName: "MyNode",
    ConnectTimeOut: time.Second * 5,
    RequestTimeOut: time.Second * 5,
    KeepaliveTime: time.Second * 1,
    NotifyCall: func (nodeNotify mxui.NodeNotify){
        fmt.Println("receive node notify: ", nodeNotify.Message)
    },
})
if err != nil {
    panic(err)
}



```




## 进入MXUI
###  浏览器访问: http://localhost:806 
能够进入MXUI登录界面, 代表服务端已经部署成功   
<img src="store\image\home.jpg" width="80%"/>


## 联系方式 
### QQ交流群: 757595139
<img src="store\image\group.jpg" width=300/>

## 问题或建议
- 有任何使用问题或者建议都可以提交至 Github issue 或者通过QQ群内联系我
- 在提交 issue 之前，请搜索相关内容是否已被提出

## License
[Apache-2.0](LICENSE)