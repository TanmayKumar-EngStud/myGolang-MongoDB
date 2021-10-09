package main

func FillUser(NewUser User)(User){
	var user User
	if NewUser.Name != "" {
		user.Name = NewUser.Name
	}
	if NewUser.Email != "" {
		user.Email = NewUser.Email
	}
	if NewUser.Password != "" {
		user.Password = NewUser.Password
	}
	return user
}