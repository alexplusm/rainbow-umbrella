package controllers

import (
	"encoding/json"
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
	userService    interfaces.IUserService
	sessionService interfaces.ISessionService
}

func NewUserController(
	userService interfaces.IUserService,
	sessionService interfaces.ISessionService) interfaces.IUserController {
	return &userController{userService: userService, sessionService: sessionService}
}

func (c userController) Register(w http.ResponseWriter, r *http.Request) {
	// TODO: why we use maxMemory param?
	if err := r.ParseMultipartForm(1024 * 1024); err != nil {
		log.Printf("[userController.Register][1]: %v", err.Error())
		return
	}

	defer func() {
		// TODO: zochem?
		if err := r.MultipartForm.RemoveAll(); err != nil {
			log.Printf("[userController.Register][-1]: %v", err.Error())
		}
	}()

	formValue := r.MultipartForm.Value

	user, err := new(dto.User).BuildFromFormValue(formValue)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err = w.Write([]byte(err.Error())); err != nil {
			log.Printf("[userController.Register][2]: %v", err.Error())
			return
		}
		return
	}

	fmt.Println("[formValue]: check", formValue)
	fmt.Println("User", user)

	ok, err := c.userService.LoginExist(user.Login)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(http.StatusText(http.StatusInternalServerError))); err != nil {
			log.Printf("[userController.Register][3]: %v", err.Error())
		}
		log.Printf("[userController.Register][3.1]: %v", err.Error())
		return
	}
	if !ok {
		w.WriteHeader(http.StatusConflict)
		if _, err := w.Write([]byte("login already exist")); err != nil {
			log.Printf("[userController.Register][4]: %v", err.Error())
		}
		return
	}

	if err := c.userService.Register(user.ToBO()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(http.StatusText(http.StatusInternalServerError))); err != nil {
			log.Printf("[userController.Register][5]: %v", err.Error())
		}
		log.Printf("[userController.Register][5.1]: %v", err.Error())
		return
	}

	// INFO: no need -> remove later
	//avatarFile := r.MultipartForm.File["avatar"]
	//if len(avatarFile) > 0 {
	//	avatarPath, err := c.saveAvatar(avatarFile[0])
	//	fmt.Println(avatarPath, err)
	//}

	w.WriteHeader(http.StatusCreated)
	//w.Header()
	// TODO: write JSON with userID?
	if _, err := w.Write([]byte("user created")); err != nil {
		log.Printf("[userController.Register][10]: %v", err.Error())
		return
	}
}

func (c userController) Login(w http.ResponseWriter, r *http.Request) {
	user := new(dto.User)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(http.StatusText(http.StatusBadRequest))); err != nil {
			log.Printf("[userController.Login][1]: %+v", err)
		}
		return
	}

	if err := json.Unmarshal(body, user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(http.StatusText(http.StatusBadRequest))); err != nil {
			log.Printf("[userController.Login][2]: %+v", err)
		}
		return
	}

	fmt.Printf("[userController.Login][1]: user raw %+v\n", user)

	userBO, err := c.userService.RetrieveByLogin(user.Login)
	if err != nil {
		fmt.Println("err", err)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(http.StatusText(http.StatusInternalServerError))); err != nil {
			log.Printf("[userController.Login][3]: %+v", err)
		}
		return
	}

	if userBO == nil {
		w.WriteHeader(http.StatusNotFound)
		if _, err := w.Write([]byte(http.StatusText(http.StatusNotFound))); err != nil {
			log.Printf("[userController.Login][4]: %+v", err)
		}
		return
	}

	fmt.Printf("[userController.Login][1]: userBO %+v\n", userBO)

	if !userBO.CheckPassword(user.Password) {
		w.WriteHeader(http.StatusNotFound)
		if _, err := w.Write([]byte(http.StatusText(http.StatusNotFound) + ": invalid password")); err != nil {
			log.Printf("[userController.Login][4]: %+v", err)
		}
		return
	}

	sessionID, err := c.sessionService.Create(userBO)
	if err != nil {
		fmt.Println("Error", err)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(http.StatusText(http.StatusInternalServerError))); err != nil {
			log.Printf("[userController.Login][5]: %+v", err)
		}
		return
	}

	responseBody := dto.UserLoginResponse{SessionID: sessionID}
	responseBodyRaw, err := json.Marshal(responseBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(http.StatusText(http.StatusInternalServerError))); err != nil {
			log.Printf("[userController.Login][10]: %+v", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(responseBodyRaw); err != nil {
		log.Printf("[userController.Login][11]: %+v", err)
	}
}

func (c userController) Details(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("kek"))
}

// INFO: unused, may be no need
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
