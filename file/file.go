// Package file 文件操作辅助函数
package file

import (
	"fmt"
	"github.com/curatorc/cngf/app"
	"github.com/curatorc/cngf/auth"
	"github.com/curatorc/cngf/helpers"
	"github.com/curatorc/cngf/logger"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"

	"github.com/gin-gonic/gin"
)

// Get 获取文件数据
func Get(path string, defaultValue string) (content []byte) {
	if !Exists(path) {
		content = []byte(defaultValue)
		err := Put(content, path)
		logger.LogIf(err)
		return
	}
	content, err := ioutil.ReadFile(path)
	logger.LogIf(err)
	return
}

// Put 将数据存入文件
func Put(data []byte, to string) error {
	err := ioutil.WriteFile(to, data, 0644)

	if err != nil {
		return err
	}
	return nil
}

// Exists 判断文件是否存在
func Exists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}
	return true
}

// NameWithoutExtension 去除文件后缀
func NameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

// SaveUploadAvatar 保存上传图片
func SaveUploadAvatar(c *gin.Context, file *multipart.FileHeader) (string, error) {
	var avatar string
	// 确保目录存在，不存在创建
	publicPath := "public"
	dirName := fmt.Sprintf("/uploads/avatars/%s/%s/", app.TimenowInTimezone().Format("2006/01/02"), auth.CurrentUID(c))

	err := os.MkdirAll(publicPath+dirName, 0755)
	if err != nil {
		logger.LogIf(err)
		return "", err
	}

	// 保存文件
	fileName := randomNameFromUploadFile(file)
	// public/uploads/avatars/2022/01/22/1/nfDaCgaWKpWWOmOt.png
	avatarPath := publicPath + dirName + fileName
	if err := c.SaveUploadedFile(file, avatarPath); err != nil {
		return avatar, err
	}

	// 裁切图片
	img, err := imaging.Open(avatarPath, imaging.AutoOrientation(true))
	if err != nil {
		return avatar, err
	}

	resizeAvatar := imaging.Thumbnail(img, 256, 256, imaging.Lanczos)
	resizeAvatarName := randomNameFromUploadFile(file)

	resizeAvatarPath := publicPath + dirName + resizeAvatarName

	err = imaging.Save(resizeAvatar, resizeAvatarPath)
	if err != nil {
		return avatar, err
	}

	// 删除老文件
	err = os.Remove(avatarPath)
	if err != nil {
		return avatar, err
	}

	return dirName + resizeAvatarName, nil
}

func randomNameFromUploadFile(file *multipart.FileHeader) string {
	return helpers.RandomString(16) + filepath.Ext(file.Filename)
}
