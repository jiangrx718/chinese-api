package file

import (
	"crm/gopkg/gins"
	"crm/gopkg/minio"
	"crm/gopkg/utils"
	"crm/gopkg/utils/files"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// FileUpload 上传文件到 MinIO
func (h *Handler) FileUpload(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		gins.BadRequest(ctx, fmt.Errorf("文件上传失败: %v", err))
		return
	}

	// 生成对象名: uploads/YYYY/MMDD/shard/uuid.ext
	now := time.Now()
	year := now.Format("2006")
	day := now.Format("0102")
	ext := files.ExtName(fileHeader.Filename)
	uuid := utils.GenUUIDWithoutUnderline()
	objectName := fmt.Sprintf("uploads/%s/%s/%s", year, day, uuid)
	if ext != "" {
		objectName = objectName + "." + ext
	}

	// 保存到临时文件再上传
	tmpPath := filepath.Join(os.TempDir(), objectName)
	if err := os.MkdirAll(filepath.Dir(tmpPath), os.ModePerm); err != nil {
		gins.ServerError(ctx, err)
		return
	}
	if err := ctx.SaveUploadedFile(fileHeader, tmpPath); err != nil {
		gins.ServerError(ctx, err)
		return
	}
	defer files.DeleteFileIfExists(tmpPath)

	// 上传到 MinIO
	if _, err := minio.UploadFile(objectName, tmpPath); err != nil {
		gins.ServerError(ctx, err)
		return
	}

	// 生成预览链接（1小时有效）
	url, err := minio.GeneratePresignedURL(objectName, time.Hour)
	if err != nil {
		// 预览链接失败不影响上传，返回空链接
		url = ""
	}

	// 返回结果
	ctx.JSON(200, gin.H{
		"code":    0,
		"message": "操作成功",
		"data": gin.H{
			"object_name": objectName,
			"preview_url": url,
			"file_size":   fileHeader.Size,
			"file_type":   strings.ToLower(files.ExtName(fileHeader.Filename)),
			"file_name":   fileHeader.Filename,
		},
	})
	return
}
