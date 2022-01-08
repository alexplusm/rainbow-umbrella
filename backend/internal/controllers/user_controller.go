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
	"rainbow-umbrella/internal/objects/dto"
)

type userController struct {
	userService interfaces.IUserService
}

func NewUserController(userService interfaces.IUserService) interfaces.IUserController {
	return &userController{userService: userService}
}

func (c userController) Register(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(1024 * 1024); err != nil {
		log.Printf("[userController.Register][2]: %v", err.Error())
		return
	}

	avatarFile := r.MultipartForm.File["avatar"]
	formValue := r.MultipartForm.Value

	user, err := new(dto.User).BuildFromFormValue(formValue)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err = w.Write([]byte(err.Error())); err != nil {
			log.Printf("[userController.Register][1]: %v", err.Error())
			return
		}
	}

	// TODO: check existence

	fmt.Println("USER: formValue", formValue)
	fmt.Printf("USER: %+v\n err: %v\n", user, err)

	if len(avatarFile) > 0 {
		avatarPath, err := c.saveAvatar(avatarFile[0])
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
