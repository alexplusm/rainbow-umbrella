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

	"github.com/go-chi/chi/v5"

	"rainbow-umbrella/internal/consts"
	"rainbow-umbrella/internal/interfaces"
	"rainbow-umbrella/internal/objects/bo"
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

	ok, err := c.userService.LoginExist(user.Login)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(http.StatusText(http.StatusInternalServerError))); err != nil {
			log.Printf("[userController.Register][3]: %v", err.Error())
		}
		log.Printf("[userController.Register][3.1]: %v", err.Error())
		return
	}
	if ok {
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
	if _, err := w.Write([]byte("user created")); err != nil {
		log.Printf("[userController.Register][10]: %v", err.Error())
		return
	}
}

func (c userController) Login(w http.ResponseWriter, r *http.Request) {
	user := new(dto.User)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error[2]", err)
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(http.StatusText(http.StatusBadRequest))); err != nil {
			log.Printf("[userController.Login][1]: %+v", err)
		}
		return
	}

	if err := json.Unmarshal(body, user); err != nil {
		fmt.Println("Error[1]", err)
		fmt.Println("body=", string(body))
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(http.StatusText(http.StatusBadRequest))); err != nil {
			log.Printf("[userController.Login][2]: %+v", err)
		}
		return
	}

	fmt.Printf("[userController.Login][1]: user raw %+v\n", user)

	userBO, err := c.userService.RetrieveByLogin(user.Login)
	if err != nil {
		fmt.Println("Error", err)
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
	login := chi.URLParam(r, "login")
	fmt.Println("YEP", login)

	user, err := c.userService.RetrieveByLogin(login)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(http.StatusText(http.StatusInternalServerError))); err != nil {
			log.Printf("[userController.Details][1]: %+v", err)
		}
		return
	}

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		if _, err := w.Write([]byte(http.StatusText(http.StatusNotFound) + ": " + login)); err != nil {
			log.Printf("[userController.Details][1]: %+v", err)
		}
		return
	}

	userDTO := new(dto.User).FromBO(user)

	responseBody, err := json.Marshal(userDTO)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(http.StatusText(http.StatusInternalServerError))); err != nil {
			log.Printf("[userController.Details][10]: %+v", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json") // TODO: why don't work
	if _, err := w.Write(responseBody); err != nil {
		log.Printf("[userController.Details][1]")
	}
}

func (c userController) List(w http.ResponseWriter, r *http.Request) {
	userFilter := new(bo.UserFilter)

	if valueRaw, ok := r.Context().Value("currentUserLogin").(string); ok {
		userFilter.ExcludeLogin = valueRaw
	}

	users, err := c.userService.List(userFilter)
	if err != nil {
		processError(w, http.StatusInternalServerError, nil)
		log.Println(fmt.Errorf("[userController.List][1]: %+v", err))
		return
	}

	usersDTO := make([]dto.User, 0, len(users))

	for _, item := range users {
		userDTO := new(dto.User).FromBO(&item)
		usersDTO = append(usersDTO, *userDTO)
	}

	bodyRaw := map[string][]dto.User{"users": usersDTO}

	body, err := json.Marshal(bodyRaw)
	if err != nil {
		processError(w, http.StatusInternalServerError, nil)
		log.Println(fmt.Errorf("[userController.List][2]: %+v", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
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
