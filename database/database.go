package database

import (
	question "example/FastTrackTestQuiz/packages/question"
	user "example/FastTrackTestQuiz/packages/user"
)

var UsersAndHighscores = []user.User{
	{
		Id:        1,
		Name:      "Thomas",
		Highscore: 2,
	},
	{
		Id:        2,
		Name:      "Susan",
		Highscore: 4,
	},
	{
		Id:        3,
		Name:      "Mary",
		Highscore: 5,
	},
	{
		Id:        4,
		Name:      "Keith",
		Highscore: 5,
	},
	{
		Id:        5,
		Name:      "Robert",
		Highscore: 1,
	},
}

func ReturnAllQuestions() []question.Question {
	sliceOfQuestions := []question.Question{
		{
			Id:       1,
			Question: "When did Golang first appear?",
			Answers:  []string{"1. 2007", "X. 2008", "2. 2009"},
		},
		{
			Id:       2,
			Question: "Which company created Golang?",
			Answers:  []string{"1. Google", "X. Microsoft", "3. IBM"},
		},
		{
			Id:       3,
			Question: "Which year did Fast Track get their \"Great Place to Work\" award?",
			Answers:  []string{"1. 2021", "X. 2022", "2. 2023"},
		},
		{
			Id:       4,
			Question: "Golang has a mascot in their logo, what is it called?",
			Answers:  []string{"1. Gooser", "X. Gopher", "2. Batman"},
		},
		{
			Id:       5,
			Question: "What is the capitol of Sweden?",
			Answers:  []string{"1. Köpenhamn", "X. Stockholm", "2. Oslo"},
		}}
	return sliceOfQuestions
}

func CorrectAnswers() []string {
	correctAnswers := []string{"2", "1", "2", "X", "X"}
	return correctAnswers
}

func GetUsers() []user.User {
	return UsersAndHighscores
}

func CreateUser(name string, score int) {
	user := new(user.User)
	user.Id = len(UsersAndHighscores) + 1
	user.Name = name
	user.Highscore = score
	UsersAndHighscores = append(UsersAndHighscores, *user)
}
