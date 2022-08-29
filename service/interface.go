package service

type User struct {
	ID   int
	Nama string
}

type UserService struct {
	db []*User
}

type UserIface interface {
	Register(u *User) string
	GetUser() []*User
	GetUserbyID(number int) *User
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

func (u *UserService) GetUserbyID(number int) *User {
	for _, v := range u.db {
		if v.ID == number {
			return v
		}
	}
	return nil
}

// func main() {
// 	var db []*service.User
// 	userSrv := service.NewUserService(db)
// 	user1 := userSrv.Register(&service.User{ID: 1, Nama: "rizki"})
// 	fmt.Println(user1)
// 	user2 := userSrv.Register(&service.User{ID: 2, Nama: "zaka"})
// 	fmt.Println(user2)
// 	user3 := userSrv.Register(&service.User{ID: 3, Nama: "fajar"})
// 	fmt.Println(user3)
// 	fmt.Println("-----------Hasil get user-------------")
// 	resGet := userSrv.GetUser()
// 	for _, v := range resGet {
// 		fmt.Println(v.ID, v.Nama)
// 	}

// 	fmt.Println("-----------Hasil get user by id-------------")
// 	getusr1 := userSrv.GetUserbyID(&service.User{ID: 1})
// 	fmt.Println(getusr1)
// }
