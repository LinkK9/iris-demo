package repo

import (
	"iris-demo/pmodel"
	"iris-demo/rbac"
	"errors"
)

var users = []pmodel.User{
	{
		User:  "root",
		Pass:  "1",
		Email: "root@gmail.com",
		Roles: pmodel.Roles{rbac.ROOT: true},
	},
	{
		User:  "bob",
		Pass:  "1",
		Email: "bob@gmail.com",
		Roles: pmodel.Roles{rbac.GUEST: true},
	},
	{
		User:  "long",
		Pass:  "1",
		Email: "long@gmail.com",
		Roles: pmodel.Roles{rbac.STAFF: true},
	},
	{
		User:  "linh",
		Pass:  "1",
		Email: "linh@gmail.com",
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
