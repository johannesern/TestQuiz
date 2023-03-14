/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	QuizController "example/FastTrackTestQuiz/API/QuizController"
	UserController "example/FastTrackTestQuiz/API/UserController"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "FastTrackTestQuiz",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		run := true
		for run {
			fmt.Println("Welcome to the Quiz!\nWhat's your name?")
			var name string
			fmt.Scan(&name)
			fmt.Println("\nHi " + name + "!\nHere is the quiz:\n")
			userAnswers := Quiz()

			userScore := CheckAnsers(userAnswers)
			fmt.Println("\nDo you wish to save your highscore?[Y]es / [N]o")
			var save string
			fmt.Scan(&save)
			if strings.ToUpper(save) == "Y" {
				UserController.Create(name, userScore)
				fmt.Println("\nSaved successfully! Thank you for playing!")
				run = false
			} else {
				fmt.Println("\nThank you for playing!")
				run = false
			}
		}
	},
}

func Quiz() []string {

	//Calling API
	quiz := QuizController.Get()

	userAnswers := []string{}

	for i := 0; i < len(quiz); i++ {

		fmt.Println(quiz[i].Question)
		var input string
		notAValidChar := true
		for notAValidChar {

			fmt.Println(quiz[i].Answers)
			fmt.Scan(&input)
			input = strings.ToUpper(input)
			if input != "1" && input != "X" && input != "2" {

				fmt.Println("Only 1, X or 2 is a valid choice, try again")

			} else {
				notAValidChar = false
			}
		}
		userAnswers = append(userAnswers, input)
	}
	return userAnswers
}

func CheckAnsers(userAnswers []string) int {

	//Calling API
	correctAnswers := QuizController.GetCorrectAnswers()

	userScore := 0
	for i := 0; i < len(userAnswers); i++ {
		if userAnswers[i] == correctAnswers[i] {
			userScore++
		}
	}

	//Calling API
	allUsers := UserController.Get()

	var userBetterThanOpponent float32 = 0
	for i := 0; i < len(allUsers); i++ {

		oppenentScore := allUsers[i].Highscore

		if oppenentScore < userScore {
			userBetterThanOpponent++
		}
	}

	var numOfUsers = float32(len(allUsers))

	percent := userBetterThanOpponent / numOfUsers
	percent = percent * 100

	fmt.Printf("You had %v correct answers which is better than %0.0f%% the players.\n", userScore, percent)
	return userScore
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.FastTrackTestQuiz.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
