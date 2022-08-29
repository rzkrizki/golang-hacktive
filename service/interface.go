package service

type User struct {
	Nama string
}

type UserIface interface {
	Register(u *User) string
}

func NewUserService() UserIface {
	return &User{}
}

func (u *User) Register(user *User) string {
	return user.Nama + " Berhasil didaftarkan"
}
