package transport

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (t *Transport) Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("Error when parsing form: %s", err.Error()))
		return
	}

	// Lưu file vào thư mục uploads
	err = ctx.SaveUploadedFile(file, filepath.Join("./uploads", file.Filename))
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error when saving file: %s", err.Error()))
		return
	}

	ctx.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully", file.Filename))
}
