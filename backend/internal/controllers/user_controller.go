package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
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

	fmt.Printf("USER: %+v\n\n", user)

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
		log.Print(fmt.Errorf("[userController.Login][1]: %w", err))
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(http.StatusText(http.StatusBadRequest))); err != nil {
			log.Print(fmt.Errorf("[userController.Login][1.1]: %w", err))
		}
		return
	}

	if err := json.Unmarshal(body, user); err != nil {
		log.Print(fmt.Errorf("[userController.Login][2]: %w", err))
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(http.StatusText(http.StatusBadRequest))); err != nil {
			log.Print(fmt.Errorf("[userController.Login][2.1]: %w", err))
		}
		return
	}

	fmt.Printf("[userController.Login][1]: user raw %+v\n", user)

	userBO, err := c.userService.RetrieveByLogin(user.Login)
	if err != nil {
		log.Print(fmt.Errorf("[userController.Login][3]: %w", err))
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(http.StatusText(http.StatusInternalServerError))); err != nil {
			log.Print(fmt.Errorf("[userController.Login][3.1]: %w", err))
		}
		return
	}

	if userBO == nil {
		w.WriteHeader(http.StatusNotFound)
		if _, err := w.Write([]byte(http.StatusText(http.StatusNotFound))); err != nil {
			log.Print(fmt.Errorf("[userController.Login][4]: %w", err))
		}
		return
	}

	fmt.Printf("[userController.Login][1]: userBO %+v\n", userBO)

	if !userBO.CheckPassword(user.Password) {
		w.WriteHeader(http.StatusNotFound)
		if _, err := w.Write([]byte(http.StatusText(http.StatusNotFound) + ": invalid password")); err != nil {
			log.Print(fmt.Errorf("[userController.Login][5]: %w", err))
		}
		return
	}

	sessionID, err := c.sessionService.Create(userBO)
	if err != nil {
		log.Print(fmt.Errorf("[userController.Login][6]: %w", err))
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(http.StatusText(http.StatusInternalServerError))); err != nil {
			log.Print(fmt.Errorf("[userController.Login][6.1]: %w", err))
		}
		return
	}

	responseBody := dto.UserLoginResponse{SessionID: sessionID}
	responseBodyRaw, err := json.Marshal(responseBody)
	if err != nil {
		log.Print(fmt.Errorf("[userController.Login][7]: %w", err))
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(http.StatusText(http.StatusInternalServerError))); err != nil {
			log.Print(fmt.Errorf("[userController.Login][7.1]: %w", err))
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(responseBodyRaw); err != nil {
		log.Print(fmt.Errorf("[userController.Login][8]: %w", err))
	}
}

func (c userController) Details(w http.ResponseWriter, r *http.Request) {
	login := chi.URLParam(r, "login")

	user, err := c.userService.RetrieveByLogin(login)
	if err != nil {
		log.Print(fmt.Errorf("[userController.Details][1]: %w", err))
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(http.StatusText(http.StatusInternalServerError))); err != nil {
			log.Print(fmt.Errorf("[userController.Details][1.1]: %w", err))
		}
		return
	}
	if user == nil {
		log.Print(fmt.Errorf("[userController.Details][2]: %w", err))
		w.WriteHeader(http.StatusNotFound)
		if _, err := w.Write([]byte(http.StatusText(http.StatusNotFound) + ": " + login)); err != nil {
			log.Print(fmt.Errorf("[userController.Details][2.1]: %w", err))
		}
		return
	}

	currUserLogin, _ := c.sessionService.GetCurrentUserFromCtx(r.Context())
	friendshipStatus, err := c.userService.GetUsersFriendshipStatus(currUserLogin, login)
	if err != nil {
		log.Print(fmt.Errorf("[userController.Details][3]: %w", err))
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(http.StatusText(http.StatusInternalServerError))); err != nil {
			log.Print(fmt.Errorf("[userController.Details][3.1]: %w", err))
		}
		return
	}

	userDTO := new(dto.User).FromBO(user)

	body := map[string]interface{}{
		"user":             userDTO,
		"friendshipStatus": friendshipStatus,
	}

	// TODO: build body !!!

	fmt.Printf("currUserLogin: %v | friendshipStatus = %v\n\n", currUserLogin, friendshipStatus)

	responseBody, err := json.Marshal(body)
	if err != nil {
		log.Print(fmt.Errorf("[userController.Details][4]: %w", err))
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(http.StatusText(http.StatusInternalServerError))); err != nil {
			log.Print(fmt.Errorf("[userController.Details][4.1]: %w", err))
		}
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(responseBody); err != nil {
		log.Print(fmt.Errorf("[userController.Details][5]: %w", err))
	}
}

func (c userController) List(w http.ResponseWriter, r *http.Request) {
	queryParams, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Printf("[userController.List][01]: %v", err)
	}

	userFilter := new(bo.UserFilter).
		Build().
		SetLimitAndOffset(queryParams.Get("limit"), queryParams.Get("offset")).
		SetSearch(queryParams.Get("search"))

	fmt.Printf("\nuserFilter: %+v\n\n", userFilter)

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
