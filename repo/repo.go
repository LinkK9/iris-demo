package repo

import (
	"errors"
	"iris-demo/pmodel"
	"iris-demo/rbac"
)

var users = []pmodel.User{
	{
		ID: 1,
		User:  "root",
		Pass:  "1",
		Email: "root@gmail.com",
		BirthYear: "1975",
		Job:   "Developer",
		Roles: pmodel.Roles{rbac.ROOT: true},
	},
	{
		ID: 2,
		User:  "bob",
		Pass:  "1",
		Email: "bob@gmail.com",
		BirthYear: "1994",
		Job:   "Công nhân xây dựng",
		Roles: pmodel.Roles{rbac.GUEST: true},
	},
	{
		ID: 3,
		User:  "long",
		Pass:  "1",
		Email: "long@gmail.com",
		BirthYear: "1995",
		Job:   "Nhân viên bán hàng",
		Roles: pmodel.Roles{rbac.STAFF: true},
	},
	{
		ID: 4,
		User:  "linh",
		Pass:  "1",
		Email: "linh@gmail.com",
		BirthYear: "1992",
		Job:   "Xe ôm",
		Roles: pmodel.Roles{rbac.EDITOR: true},
	},
}

func QueryByEmail(email string) (user *pmodel.User, err error) {
	for _, obj := range users {
		if obj.Email == email {
			user = new(pmodel.User)
			*user = obj
			return user, nil
		}
	}
	return nil, errors.New("User not found")
}

func GetUserList() ([]string) {
	var list []string
	for _, obj := range users {
		list = append(list, obj.User)
	}
	return list
}

func GetUserInfoList() []pmodel.User {
return users
}

func QueryById(id int) (user *pmodel.User, err error) {
	for _, obj := range users {
		if obj.ID == id {
			return &obj, nil
		}
	}
	return nil, errors.New("User not found")
}

func AddNewUser(user pmodel.User) {
users= append(users, user)
}

func ModifyById(user pmodel.User){
	for i, obj := range users {
		if obj.ID == user.ID {
			users[i] = user
		}
	}
}

func DeleteUser(id int) {
	for i, obj := range users {
		if obj.ID == id {
			users = append(users[:i], users[i + 1:]...)
		}
	}
}