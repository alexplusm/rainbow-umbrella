package interfaces

type IUserService interface {
	Register()
	GenerateAvatarFileName(originalName string) string
}
