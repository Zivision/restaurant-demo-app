package user

type User struct {
	id   int
	name string
	time string
}

func CreateUser(id int, name string) *User {
	return &User{
		id:   id,
		name: name,
	}

}