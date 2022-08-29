package service

type User struct {
	Nama string
}

type UserService struct {
	db []*User
}

type UserIface interface {
	Register(u *User) string
	GetUser() []*User
}

func NewUserService(db []*User) UserIface {
	return &UserService{
		db: db,
	}
}

func (u *UserService) Register(user *User) string {
	u.db = append(u.db, user)
	return user.Nama + " berhasil didaftarkan"
}

func (u *UserService) GetUser() []*User {
	return u.db
}
