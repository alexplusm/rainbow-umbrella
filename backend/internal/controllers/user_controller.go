package controllers

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"

	"rainbow-umbrella/internal/consts"
	"rainbow-umbrella/internal/interfaces"
)

type userController struct {
	userService interfaces.IUserService
}

func NewUserController(userService interfaces.IUserService) interfaces.IUserController {
	return &userController{userService: userService}
}

func (c userController) Register(w http.ResponseWriter, r *http.Request) {
	// TODO: check existence

	if err := r.ParseMultipartForm(1024 * 1024); err != nil {
		err = fmt.Errorf("[userController.Register]: %v", err)
		log.Fatal(err)
	}

	file := r.MultipartForm.File["file"]
	formValue := r.MultipartForm.Value

	if len(file) > 0 {
		avatarPath, err := c.saveAvatar(file[0])
		fmt.Println(avatarPath, err)
	}

	fmt.Println("formValue", formValue)

	w.Write([]byte("kekes"))
}

func (c userController) saveAvatar(fileHeader *multipart.FileHeader) (string, error) {
	avatarFilePath := path.Join(consts.MediaRootDir, c.userService.GenerateAvatarFileName(fileHeader.Filename))

	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("[userController.saveAvatar][1]: %+v", err)
	}

	avatarFile, err := os.Create(avatarFilePath)
	if err != nil {
		return "", fmt.Errorf("[userController.saveAvatar][2]: %+v", err)
	}

	if _, err := io.Copy(avatarFile, file); err != nil {
		return "", fmt.Errorf("[userController.saveAvatar][3]: %+v", err)
	}

	if err := avatarFile.Close(); err != nil {
		return "", fmt.Errorf("[userController.saveAvatar][4]: %+v", err)
	}

	if err := file.Close(); err != nil {
		return "", fmt.Errorf("[userController.saveAvatar][5]: %+v", err)
	}

	return avatarFilePath, nil
}
