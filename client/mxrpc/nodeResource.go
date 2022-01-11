/*
 * Copyright 2021 liyiligang.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mxrpc

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/liyiligang/base/component/Jrpc"
	"github.com/liyiligang/base/component/Jtool"
	"github.com/liyiligang/mxrpc-go-client/protoFiles/protoManage"
	"github.com/liyiligang/mxrpc-go-client/typedef/constant"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func (client *Client) CheckNodeResourceWithFile(filePath string) (*protoManage.NodeResource, error){
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(){
		err := file.Close()
		if err != nil {
			client.RpcStreamError("file close error", err)
		}
	}()
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
	return client.checkNodeResource(fileName, md5str, fileInfo.Size())
}

func (client *Client) CheckNodeResourceWithByte(name string, data []byte) (*protoManage.NodeResource, error){
	md := md5.Sum(data)
	md5str := fmt.Sprintf("%x", md)
	return client.checkNodeResource(name, md5str, int64(len(data)))
}


func (client *Client) UploadNodeResourceWithFile(filePath string) (*protoManage.NodeResource, error){
	nodeResource ,err := client.CheckNodeResourceWithFile(filePath)
	if err != nil {
		return nil, err
	}
	if nodeResource.Base.ID == 0 {
		file, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}
		defer func(){
			err := file.Close()
			if err != nil {
				client.RpcStreamError("file close error", err)
			}
		}()
		err = client.uploadNodeResource(nodeResource, file)
		if err != nil {
			return nil, err
		}
	}
	return nodeResource, nil
}

func (client *Client) UploadNodeResourceWithBytes(name string, data []byte) (*protoManage.NodeResource, error){
	nodeResource ,err := client.CheckNodeResourceWithByte(name, data)
	if err != nil {
		return nil, err
	}
	if nodeResource.Base.ID == 0 {
		buffer := bytes.NewBuffer(data)
		err = client.uploadNodeResource(nodeResource, buffer)
		if err != nil {
			return nil, err
		}
	}
	return nodeResource, nil
}

func (client *Client) DownloadNodeResourceWithFile(url string, filePath string) (uErr error) {
	idStr := strings.Split(url, "_")
	id := Jtool.StringToInt64(idStr[0])
	name := idStr[1]
	file, err := os.Create(filePath + name)
	if err != nil {
		return err
	}
	defer func(){
		_ = file.Close()
		if uErr != nil {
			_ = os.Remove(filePath)
		}
	}()
	return client.downloadNodeResource(id, file)
}

func (client *Client) DownloadNodeResourceWithBytes(url string) (string, []byte, error) {
	idStr := strings.Split(url, "_")
	id := Jtool.StringToInt64(idStr[0])
	name := idStr[1]
	buffer := bytes.NewBuffer(nil)
	err := client.downloadNodeResource(id, buffer)
	if err != nil {
		return "", nil, err
	}
	return name, buffer.Bytes(), nil
}

func (client *Client) checkNodeResource(name string, md5 string, size int64) (*protoManage.NodeResource, error) {
	node, err := client.getNode()
	if err != nil {
		return nil, err
	}
	req := &protoManage.NodeResource{
		UploaderID:node.Base.ID,
		UploaderType: protoManage.NotifySenderType_NotifySenderTypeNode,
		Name: name,
		Md5: md5,
		Sizes: size,
		Type: protoManage.NodeResourceType_NodeResourceTypeCache,
	}
	ctx, _ := context.WithTimeout(context.Background(), client.config.RequestTimeOut)
	return client.engine.CheckNodeResource(ctx, req)
}

func (client *Client) uploadNodeResource(nodeResource *protoManage.NodeResource, read io.Reader) error {
	pbData, err := nodeResource.Marshal()
	if err != nil {
		return err
	}
	upload, err := client.engine.UploadNodeResource(Jrpc.SetRpcStreamClientHeader(pbData))
	if err != nil {
		return err
	}
	err = Jtool.ReadIOWithSize(read, constant.ConstRpcClientMaxMsgSize/2, func(buf []byte) error{
		return upload.Send(&protoManage.ReqNodeResourceUpload{Data: buf})
	})
	if err != nil {
		return err
	}
	_, err = upload.CloseAndRecv()
	return err
}

func (client *Client) downloadNodeResource(id int64, write io.Writer) error {

	download, err := client.engine.DownloadNodeResource(context.Background(),
		&protoManage.ReqNodeResourceDownload{NodeResource: protoManage.NodeResource{
			Base: protoManage.Base{
				ID:  id,
			},
		}})
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
		_, err = write.Write(req.Data)
		if err != nil {
			return err
		}
	}
	return nil
}