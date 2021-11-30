// Copyright 2021 The Authors. All rights reserved.
// Author: liyiligang
// Date: 2021/11/26 16:23
// Description:

package klee

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/liyiligang/base/commonConst"
	"github.com/liyiligang/base/component/Jrpc"
	"github.com/liyiligang/base/component/Jtool"
	"github.com/liyiligang/klee-client-go/protoFiles/protoManage"
	"io"
	"os"
	"path/filepath"
)

func (client *ManageClient) CheckNodeResourceWithFile(filePath string) (*protoManage.NodeResource, error){
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileName := filepath.Base(filePath)
	md := md5.New()
	if _, err := io.Copy(md, file); err != nil {
		return nil, err
	}
	md5str := hex.EncodeToString(md.Sum(nil))
	req := &protoManage.NodeResource{
		UUID: md5str + "_" + fileName,
		Name: fileName,
		Md5: md5str,
		Sizes: fileInfo.Size(),
		Type: protoManage.NodeResourceType_NodeResourceTypeCache,
	}
	ctx, _ := context.WithTimeout(context.Background(), client.config.RequestTimeOut)
	return client.engine.CheckNodeResource(ctx, req)
}

func (client *ManageClient) CheckNodeResourceWithByte(name string, data []byte) (*protoManage.NodeResource, error){
	md := md5.Sum(data)
	md5str := fmt.Sprintf("%x", md)
	req := &protoManage.NodeResource{
		UUID: md5str + "_" + name,
		Name: name,
		Md5: md5str,
		Sizes: int64(len(data)),
		Type: protoManage.NodeResourceType_NodeResourceTypeCache,
	}
	ctx, _ := context.WithTimeout(context.Background(), client.config.RequestTimeOut)
	return client.engine.CheckNodeResource(ctx, req)
}


func (client *ManageClient) UploadNodeResourceWithFile(filePath string) (*protoManage.NodeResource, error){
	nodeResource ,err := client.CheckNodeResourceWithFile(filePath)
	if err != nil {
		return nil, err
	}
	if !nodeResource.IsExist {
		file, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		err = client.uploadNodeResource(nodeResource, file)
		if err != nil {
			return nil, err
		}
	}
	return nodeResource, nil
}

func (client *ManageClient) UploadNodeResourceWithBytes(name string, data []byte) (*protoManage.NodeResource, error){
	nodeResource ,err := client.CheckNodeResourceWithByte(name, data)
	if !nodeResource.IsExist {
		buffer := bytes.NewBuffer(data)
		err = client.uploadNodeResource(nodeResource, buffer)
		if err != nil {
			return nil, err
		}
	}
	return nodeResource, nil
}

func (client *ManageClient) DownloadNodeResourceWithFile(fileName string, fileDir string) (uErr error) {
	filePath := fileDir + fileName
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(){
		_ = file.Close()
		if uErr != nil {
			_ = os.Remove(filePath)
		}
	}()
	return client.downloadNodeResource(&protoManage.NodeResource{UUID: fileName}, file)
}

func (client *ManageClient) DownloadNodeResourceWithBytes(fileName string) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	err := client.downloadNodeResource(&protoManage.NodeResource{UUID: fileName}, buffer)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (client *ManageClient) uploadNodeResource(nodeResource *protoManage.NodeResource, read io.Reader) error {
	pbData, err := nodeResource.Marshal()
	if err != nil {
		return err
	}
	upload, err := client.engine.UploadNodeResource(Jrpc.SetRpcStreamClientHeader(pbData))
	if err != nil {
		return err
	}
	err = Jtool.ReadIOWithSize(read, commonConst.GrpcMaxMsgSize/2, func(buf []byte) error{
		return upload.Send(&protoManage.ReqNodeResourceUpload{Data: buf})
	})
	if err != nil {
		return err
	}
	_, err = upload.CloseAndRecv()
	return err
}

func (client *ManageClient) downloadNodeResource(nodeResource *protoManage.NodeResource, write io.Writer) error {
	download, err := client.engine.DownloadNodeResource(context.Background(),
		&protoManage.ReqNodeResourceDownload{NodeResource: *nodeResource})
	if err != nil {
		return err
	}
	for {
		req, err := download.Recv()
		if err != nil {
			if errors.Is(err, io.EOF){
				break
			}
			return err
		}
		write.Write(req.Data)
	}
	return nil
}