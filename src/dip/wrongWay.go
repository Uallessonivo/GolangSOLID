package dip

// Baixo nivel
func GetUserByIDFomMySQLRepo(user_id string) *User {
	user, ok := inMemoUsers[user_id]
	if !ok {
		return nil
	}

	return &user
}

// Baixo nivel
func InsertUserIntoMySQLRepo(user_id string, name string) *User {
	user := User{user_id, name}
	inMemoUsers[user_id] = user
	return &user
}

// Alto nivel
func GetUserByIDService(user_id string) *User {
	return GetUserByIDFomMySQLRepo(user_id)
}

// Alto nivel
func InsertUserService(name string) *User {
	return InsertUserIntoMySQLRepo("10000", name)
}
