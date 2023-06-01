package main

import (
	"fmt"
)

type User struct {
	Id       int
	Username string
	Roles    []Role
}

type Role struct {
	Id          int
	Name        string
	Permissions []Permission
}

type Permission struct {
	Id   int
	Name string
}

func checkPermission(user User, permissionName string) bool {
	for _, role := range user.Roles {
		for _, permission := range role.Permissions {
			if permission.Name == permissionName {
				return true
			}
		}
	}
	return false

}

// assign roles to a user

func assignRole(u *User, role Role) {
	u.Roles = append(u.Roles,role)
}


// remove users roles
func demoteUser(u *User,roleID int){
	for i,role := range u.Roles{
		if role.Id == roleID{
			u.Roles = append(u.Roles[:i],u.Roles[i+1:]... )
			break
		}
	}

}

func main() {
	permission_1 := Permission{Id: 1,Name: "read"}
	permission_2 := Permission{Id: 2,Name: "delete"}
	permission_3 := Permission{Id: 3,Name: "write"}




	role := Role{Id: 1,Name: "Admin",Permissions: []Permission{permission_1,permission_2,permission_3}}
	role2 := Role{Id: 2,Name: "Editor",Permissions: []Permission{permission_1,permission_3}}


	user := User{Id: 1,Username:"kwame233"}
	user2 := User{Id: 2,Username: "adade453"}

	assignRole(&user,role)
	assignRole(&user2,role2)


	// check permissions

	fmt.Printf("User has read permissions", checkPermission(user,"read"))
	fmt.Println()
	fmt.Println()
	fmt.Printf("User has delete permissions", checkPermission(user2,"delete"))
	

}