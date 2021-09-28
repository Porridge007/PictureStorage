package controllers

import (
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type PictureController struct {
	beego.Controller
}

// @router / [get]
func (c *PictureController) Get() {
	c.TplName = "upload.html"
}

func (c *PictureController) Post() {
	f, h, _ := c.GetFile("myfile") //获取上传的文件
	defer f.Close()                //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	pathPrefix, error := beego.AppConfig.String("picturepath")
	if error != nil {
		fmt.Printf("config error: %s", pathPrefix)
		pathPrefix = "."
	}
	path := pathPrefix + h.Filename //保存路径
	fmt.Println(path)
	if error :=c.SaveToFile("myfile", path); error != nil{ //存文件
		fmt.Println(error)
	}
}
