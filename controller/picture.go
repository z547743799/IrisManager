package controller

import (
	"io"
	"path"

	"github.com/kataras/iris"
	"github.com/weilaihui/fdfs_client"
)

type PictureController struct {
	Ctx iris.Context
}

func (c *PictureController) PostUpload() {
	uploadFile, Handle, _ := c.Ctx.FormFile("uploadFile")

	resp := make(map[string]interface{})
	suffix := path.Ext(Handle.Filename) //获取到的是.jpg，有点。下面把.去掉
	//判断是否为图片
	if suffix != ".jpg" && suffix != ".png" && suffix != ".gif" && suffix != ".jpeg" {
		resp["errno"] = "--"
		resp["errmsg"] = "--"
		return
	}
	//去掉.jpg前面的.，变成jpg
	next := suffix[1:]
	//创建hd.Size大小的[]byte数组用来存放fileData.Read读出来的[]byte数据

	ns := make(map[string]interface{})
	Fileb := make([]byte, 0)
	b := make([]byte, 1024)
	for {
		n, err := uploadFile.Read(b)

		if err == io.EOF {
			break
		}
		Fileb = append(Fileb, b[:n]...)
	}
	//fast, err :=fdfs_client.NewClientWithConfig("/home/lzw/DarkShell/src/gitlab.com/z547743799/IrisManager/config/client.conf")

	fast, err := fdfs_client.NewFdfsClient("/home/lzw/DarkShell/src/gitlab.com/z547743799/IrisManager/config/client.conf")
	if err != nil {
		ns["err"] = 1
		ns["url"] = err

	}
	response, err := fast.UploadByBuffer(Fileb, next)
	//response, err := fast.UploadByBuffer(Fileb, Handle.Filename)
	if err == nil {
		ns["error"] = 0
		ns["url"] = "192.168.25.133/" + response.RemoteFileId
	} else {
		ns["error"] = 1
		ns["massge"] = err.Error()
	}
	c.Ctx.JSON(ns)

}
