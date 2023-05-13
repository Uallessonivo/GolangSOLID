package dip

import "fmt"

var inMemoUsers map[string]User

func init() {
	inMemoUsers = make(map[string]User)
}

type User struct {
	id   string
	name string
}

func main() {
	// Wrong way
	name := "Uallesson"
	user_inserted := InsertUserService(name)
	fmt.Printf("User inserted: %#v \n", user_inserted)

	user_returned := GetUserByIDService(user_inserted.id)
	fmt.Printf("User returned from get: %#v \n", user_returned)

	// Right way
	mysqlRepo := &mysqlRepository{}
	userServices := &userServices{mysqlRepo}

	userInserted := userServices.InsertUserService(name)
	fmt.Printf("User inserted: %#v \n", userInserted)

	userReturned := userServices.GetUserByIdService(userInserted.id)
	fmt.Printf("User returned from get: %#v \n", userReturned)
}
