package repo

import (
	"errors"

	"github.com/TechMaster/core/pmodel"
	"github.com/TechMaster/core/rbac"

	"github.com/segmentio/ksuid"
)

type MyUser struct {
	ID        string
	User      string
	Pass      string
	Email     string
	BirthDay string
	Job       string
	Avatar    string
	Roles     pmodel.Roles
}

func (u *MyUser) createID() {
	u.ID = ksuid.New().String()
}

func (u *MyUser) writeAvatar(fileName string) {
	u.Avatar = fileName
}

var users = []MyUser{
	{
		ID:        "1554646",
		User:      "root",
		Pass:      "1",
		Email:     "root@gmail.com",
		BirthDay: "1975-08-11",
		Job:       "Developer",
		Roles:     pmodel.Roles{rbac.ROOT: true},
	},
	{
		ID:        "24546",
		User:      "bob",
		Pass:      "1",
		Email:     "bob@gmail.com",
		BirthDay: "1994-05-25",
		Job:       "Công nhân xây dựng",
		Roles:     pmodel.Roles{rbac.STUDENT: true},
	},
	{
		ID:        "3789",
		User:      "long",
		Pass:      "1",
		Email:     "long@gmail.com",
		BirthDay: "1995-02-27",
		Job:       "Nhân viên bán hàng",
		Roles:     pmodel.Roles{rbac.SALE: true},
	},
	{
		ID:        "4121",
		User:      "linh",
		Pass:      "1",
		Email:     "linh@gmail.com",
		BirthDay: "1992-10-26",
		Job:       "Xe ôm",
		Roles:     pmodel.Roles{rbac.EDITOR: true},
	},
}

func QueryByEmail(email string) (user *MyUser, err error) {
	for _, obj := range users {
		if obj.Email == email {
			user = new(MyUser)
			*user = obj
			return user, nil
		}
	}
	return nil, errors.New("User not found")
}

func GetUserList() []string {
	var list []string
	for _, obj := range users {
		list = append(list, obj.User)
	}
	return list
}

func GetUserInfoList() []MyUser {
	return users
}

func QueryById(id string) (user *MyUser, err error) {
	for _, obj := range users {
		if obj.ID == id {
			return &obj, nil
		}
	}
	return nil, errors.New("User not found")
}

func AddNewUser(user MyUser) {
	newUser := user
	newUser.createID()

	users = append(users, newUser)
}

func ModifyUser(user MyUser) {
	for i, obj := range users {
		if obj.ID == user.ID {
			users[i].BirthDay = user.BirthDay
			users[i].User  = user.User
			users[i].Email = user.Email
			users[i].Job = user.Job
		}
	}
}

func DeleteUser(id string) {
	for i, obj := range users {
		if obj.ID == id {
			users = append(users[:i], users[i+1:]...)
		}
	}
}

func GetIdByEmail(email string) (string, error) {
	for _, obj := range users {
		if obj.Email == email {
			return obj.ID, nil
		}
	}
	return "", errors.New("id not found")
}

func WriteAvatarFileName(id, fileName string) error {
	for i, obj := range users {
		if obj.ID == id {
			users[i].writeAvatar(fileName)
			return nil
		}
	}
	return errors.New("id not found")
}
