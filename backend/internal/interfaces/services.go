package interfaces

type IUserService interface {
	Register()
	LoginExist(login string) (bool, error)
	GenerateAvatarFileName(originalName string) string
}
