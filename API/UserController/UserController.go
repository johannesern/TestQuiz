package User

import (
	database "example/FastTrackTestQuiz/database"
	user "example/FastTrackTestQuiz/packages/user"
)

func Create(name string, score int) {
	database.CreateUser(name, score)
}

func Get() []user.User {
	return database.GetUsers()
}
