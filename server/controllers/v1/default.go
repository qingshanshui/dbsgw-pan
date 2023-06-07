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
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"mime/multipart"
	"os"
	"path"
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
	fmt.Println("--------------------------------MergeFile--------------------------------")
	// hash值（区分当前文件是那个的，也可以用uuid，nanoid，等）
	fileId := ctx.FormValue("fileId")
	fileIndex := ctx.FormValue("fileIndex")
	fileName := ctx.FormValue("fileName")
	// 获取文件后缀
	extName := path.Ext(fileName)
	atom, err := strconv.Atoi(fileIndex)
	if err != nil {
		return ctx.JSON(t.Fail(err))
	}
	p := "static/upload/" + fileId + extName
	newFile, err := os.OpenFile(p, os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0766)
	if err != nil {
		return ctx.JSON(t.Fail(err))
	}
	index := 0
	for {
		if atom == index {
			break
		}
		fmt.Println("--------------------------------MergeFile--------------------------------")
		// 文件流 存的路径 ：public/blob_1lV4DJWf2qs8MdQojPMwb_1
		filePath := "static/" + "blob" + "_" + fileId + "_" + strconv.Itoa(index)
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
	defer func() {
		newFile.Close()
	}()
	return ctx.JSON(t.Ok(map[string]interface{}{
		"msg":  "合并成功",
		"切片序号": p,
	}))
}

// ChunkFile 上传切片
func (t *DefaultController) ChunkFile(ctx *fiber.Ctx) error {
	// 文件名名称
	//fileName := ctx.FormValue("fileName")
	// hash值（区分当前文件是那个的，也可以用uuid，nanoid，等）
	fileId := ctx.FormValue("fileId")
	fileIndex := ctx.FormValue("fileIndex")
	// 接收文件的file 分片
	file, err := ctx.FormFile("fileChunk")
	if err != nil {
		return ctx.JSON(t.Fail(err))
	}
	// 文件流 存的路径 ：public/blob_1lV4DJWf2qs8MdQojPMwb_1
	filePath := "static/" + file.Filename + "_" + fileId + "_" + fileIndex
	// 转成file
	upFile, _ := file.Open()
	// 创建文件
	fileBool, err := createFile(filePath, upFile)
	if !fileBool {
		return ctx.JSON(t.Fail(err))
	}
	return ctx.JSON(t.Ok(map[string]interface{}{
		"msg":  "上传成功",
		"切片序号": fileIndex,
	}))
}

// 创建文件
func createFile(filePath string, upFile multipart.File) (bool, error) {
	fileBool, err := utils.PathExists(filePath)
	if fileBool && err == nil {
		return true, errors.New("文件以存在")
	} else {
		newFile, err := os.Create(filePath)
		data := make([]byte, 1024, 1024)
		for {
			total, err := upFile.Read(data)
			if err == io.EOF {
				break
			}
			_, err = newFile.Write(data[:total])
			if err != nil {
				return false, errors.New("文件上传失败")
			}
		}
		defer newFile.Close()
		if err != nil {
			return false, errors.New("创建文件失败")
		}
	}
	return true, nil
}

// Upload 文件上传
func (t *DefaultController) Upload(c *fiber.Ctx) error {
	//f_path := c.FormValue("f_path")   // 保存路径
	fName := c.FormValue("f_name") // 文件名称
	fSize := c.FormValue("f_size") // 文件大小
	//f_start := c.FormValue("f_start") // 文件传输大小
	file, _ := c.FormFile("blob") // 文件流

	openFile, _ := file.Open()
	p := "static/" + fName
	newFile, _ := os.OpenFile(p, os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0766)

	data := make([]byte, 1024, 1024)
	for {
		total, err := openFile.Read(data)
		if err == io.EOF {
			openFile.Close()
			break
		}
		_, err = newFile.Write(data[:total])
	}
	defer func() {
		newFile.Close()
	}()
	stat, _ := newFile.Stat()
	atopSize, _ := strconv.Atoi(fSize)
	if stat.Size() == int64(atopSize) {

		Type, Mime, err := initalize.GetFileType(file)
		if err != nil {
			return c.JSON(t.Fail(err))
		}

		// 保存文件
		fi := models.NewFileInfo()
		fi.CreatedAt = time.Now()
		fi.Name = fName
		fi.Path = fName
		fi.Size = atopSize
		fi.Type = Type
		fi.MIME = Mime
		if err := fi.Create(); err != nil {
			return c.JSON(t.Fail(err))
		}
		return c.JSON(t.Ok("文件上传成功"))
	}
	return c.JSON(t.Ok(stat.Size()))
}

// RandomImg 随机图片
func (t *DefaultController) RandomImg(c *fiber.Ctx) error {
	api, err := service.NewDefaultService().RandomImg()
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	fmt.Println(api)
	return c.SendFile(api.Path)
}
