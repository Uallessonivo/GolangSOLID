package dip

type UserRepository interface {
	InsertUser(user_id string, name string) *User
	GetById(user_id string) *User
}

type mysqlRepository struct{}

type userServices struct {
	repo UserRepository
}

// Baixo nivel
func (m *mysqlRepository) InsertUser(user_id string, name string) *User {
	user := User{user_id, name}
	inMemoUsers[user_id] = user
	return &user
}

// Baixo nivel
func (m *mysqlRepository) GetById(user_id string) *User {
	user, ok := inMemoUsers[user_id]
	if !ok {
		return nil
	}

	return &user
}

// Alto nivel
func (s *userServices) GetUserByIdService(user_id string) *User {
	return s.repo.GetById(user_id)
}

// Alto nivel
func (s *userServices) InsertUserService(name string) *User {
	return s.repo.InsertUser("1000000", name)
}
