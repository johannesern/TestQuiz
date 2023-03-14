package quiz

import (
	database "example/FastTrackTestQuiz/database"
	question "example/FastTrackTestQuiz/packages/question"
)

func Get() []question.Question {
	var allquestions = database.ReturnAllQuestions()
	return allquestions
}

func GetCorrectAnswers() []string {
	return database.CorrectAnswers()
}
