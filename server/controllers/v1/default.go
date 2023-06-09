package v1

import (
	"bufio"
	"errors"
	"fiber-layout/controllers"
	"fiber-layout/initalize"
	"fiber-layout/models"
	"fiber-layout/pkg/utils"
	"fiber-layout/service"
	"fiber-layout/validator"
	"fiber-layout/validator/form"
	"github.com/gofiber/fiber/v2"
	"io"
	"os"
	"strconv"
	"time"
)

type DefaultController struct {
	controllers.Base
}

func NewDefaultController() *DefaultController {
	return &DefaultController{}
}

// GetList 获取文件夹信息
func (t *DefaultController) GetList(c *fiber.Ctx) error {
	// 初始化参数结构体
	ListRequestForm := form.ListRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &ListRequestForm); err != nil {
		return err
	}
	// 实际业务调用
	api, err := service.NewDefaultService().GetList(ListRequestForm)
	if err != nil {
		initalize.Log.Info(err)
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok(api))
}

// GetFile 获取文件信息
func (t *DefaultController) GetFile(c *fiber.Ctx) error {
	// 初始化参数结构体
	GetRequestForm := form.GetRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &GetRequestForm); err != nil {
		return err
	}
	// 实际业务调用
	api, err := service.NewDefaultService().GetFile(GetRequestForm)
	if err != nil {
		initalize.Log.Info(err)
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok(api))
}

// Download 文件流的形式返回前端
func (t *DefaultController) Download(c *fiber.Ctx) error {
	// 初始化参数结构体
	DownloadRequestForm := form.DownloadRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckQueryParams(c, &DownloadRequestForm); err != nil {
		return err
	}
	pwd, _ := os.Getwd()
	url := pwd + "/static" + DownloadRequestForm.Path
	exists, err := utils.PathExists(url)
	if err != nil {
		return err
	}
	if exists {
		return c.SendFile(url)
	}
	return errors.New("文件错误")
}

// Login 登录
func (t *DefaultController) Login(c *fiber.Ctx) error {
	// 初始化参数结构体
	DownloadRequestForm := form.LoginRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &DownloadRequestForm); err != nil {
		return err
	}
	api, err := service.NewDefaultService().Login(DownloadRequestForm)
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok(api))
}

// MergeFile 合并切片
func (t *DefaultController) MergeFile(ctx *fiber.Ctx) error {
	// hash值（区分当前文件是那个的，也可以用uuid，nanoid，等）
	fileId := ctx.FormValue("fileId")
	fileIndex := ctx.FormValue("fileIndex")
	fileName := ctx.FormValue("fileName")
	// 获取文件后缀
	atom, err := strconv.Atoi(fileIndex)
	if err != nil {
		return ctx.JSON(t.Fail(err))
	}
	p := "static/" + fileName

	// 当文件存在就删除原文件
	exists, _ := utils.PathExists(p)
	if exists {
		if err := os.Remove(p); err != nil {
			return ctx.JSON(t.Fail(err))
		}
	}
	newFile, err := os.OpenFile(p, os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0766)
	if err != nil {
		return ctx.JSON(t.Fail(err))
	}
	index := 0
	for {
		if atom == index {
			break
		}
		// 文件流 存的路径 ：public/blob_1lV4DJWf2qs8MdQojPMwb_1
		filePath := "static/public/" + fileName + "_" + fileId + "_" + strconv.Itoa(index)
		f, _ := os.Open(filePath)
		r := bufio.NewReader(f)
		data := make([]byte, 1024, 1024)
		for {
			total, err := r.Read(data)
			if err == io.EOF {
				f.Close()
				os.Remove(filePath)
				break
			}
			_, err = newFile.Write(data[:total])
		}
		index++
	}
	if err := newFile.Close(); err != nil {
		return ctx.JSON(t.Fail(err))
	}
	Type, Mime, err := initalize.GetFileType(p)
	if err != nil {
		return ctx.JSON(t.Fail(err))
	}
	// 获取文件md5
	md5 := utils.GetFileMd5(p)
	// 保存文件
	fi := models.NewFileInfo()
	fi.CreatedAt = time.Now()
	fi.Name = fileName
	fi.Path = p
	fi.Size = 100
	fi.Type = Type
	fi.Md5 = md5
	fi.MIME = Mime
	if err := fi.Create(); err != nil {
		return ctx.JSON(t.Fail(err))
	}
	return ctx.JSON(t.Ok("合并成功"))
}

// ChunkFile 上传切片
func (t *DefaultController) ChunkFile(ctx *fiber.Ctx) error {
	// 文件名名称
	fileName := ctx.FormValue("fileName")
	// hash值（区分当前文件是那个的，也可以用uuid，nanoid，等）
	fileId := ctx.FormValue("fileId")
	fileIndex := ctx.FormValue("fileIndex")
	// 接收文件的file 分片
	file, err := ctx.FormFile("fileChunk")
	if err != nil {
		return ctx.JSON(t.Fail(err))
	}
	// 文件流 存的路径 ：public/blob_1lV4DJWf2qs8MdQojPMwb_1
	filePath := "static/public/" + fileName + "_" + fileId + "_" + fileIndex
	exists, _ := utils.PathExists("static/public/")
	if !exists {
		if err := os.MkdirAll("static/public/", os.ModePerm); err != nil {
			return ctx.JSON(t.Fail(err))
		}
	}
	if err := ctx.SaveFile(file, filePath); err != nil {
		return ctx.JSON(t.Fail(err))
	}
	return ctx.JSON(t.Ok("ok"))
}

// VerifyFile 检查文件是否存在
func (t *DefaultController) VerifyFile(c *fiber.Ctx) error {
	fileName := c.FormValue("fileName")
	p := "static/" + fileName
	exists, err := utils.PathExists(p)
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	if exists {
		return c.JSON(t.Fail(errors.New("存在重复文件")))
	}
	return c.JSON(t.Ok(exists))
}

// RandomImg 随机图片
func (t *DefaultController) RandomImg(c *fiber.Ctx) error {
	api, err := service.NewDefaultService().RandomImg()
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	return c.SendFile(api.Path)
}
