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
 
 [简体中文](./README.md) | English


## About
mxui-go-client is [MXUI](https://github.com/liyiligang/mxui) Golang client, which can use [MXUI](https://github.com/liyiligang/mxui) The platform quickly uses the API to generate UI components such as forms, tables, charts, file upload and download, audio and video playback and so on

## Document
- [中文](https://mxui-doc.liyiligang.com)    
- [English](https://mxui-doc.liyiligang.com/en)

## Preview
- [MXUI](https://mxui.liyiligang.com)    

## Install
```bash
go get -u github.com/liyiligang/mxui-go-client
```

## Basic Usage
```go
//initialization
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

//Define API
type resume struct {
    Name      	string
    Age       	int
    Boy		  	bool
    Occupation  string    `schema:"enum=teacher;sales;doctor"`
}

callFunc := func (form *resume) string {
    sex := "boy"
    if !form.Boy {
        sex = "girl"
    }
    return "Hello World, My Name is " + form.Name+ ", " + 
    strconv.Itoa(form.Age)+ " years old, "+ sex + ". I'm a " + 
    form.Occupation + "."
}

//Create API visual UI
err := client.RegisterNodeFunc(mxui.NodeFuncRegister{
    Name:     "Hello World",
    CallFunc: callFunc,
})
if err != nil {
    fmt.Println(err)
}
```

## Generated UI
### Request parameters
<img src="store\image\request_en.jpg" width="60%"/>

### Return value
<img src="store\image\response.jpg" width="60%"/>


## Contact information
### QQ Group: 757595139
<img src="store\image\group.jpg" width=300/>

## Questions or suggestions
- Any use questions or suggestions can be submitted to GitHub issue or contact me through QQ group
- Before submitting the issue, please search whether the relevant content has been proposed

## License
[Apache-2.0](LICENSE)