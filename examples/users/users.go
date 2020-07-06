package main

import (
	"fmt"
	"os"

	"github.com/kenzo0107/backlog"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	baseURL := os.Getenv("BASE_URL")
	c := backlog.New(apiKey, baseURL)

	user, err := c.GetUserMySelf()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("user ID: %d, Name %s, userID: %s\n", *user.ID, *user.Name, *user.UserID)

	// opts := &backlog.GetUserStarCountOptions{
	// 	// Since: backlog.String("2020-07-01"),
	// 	Until: backlog.String("2020-01-01"),
	// }
	// count, err := c.GetUserStarCount(159237, opts)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("count", count)

	// users, err := c.GetUsers()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, user := range users {
	// 	fmt.Printf("id: %v, name: %v\n", *user.ID, *user.Name)
	// }

	// user, err := c.GetUser(159239)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("id: %v, name: %v\n", *user.ID, *user.Name)

	// activities, err := c.GetUserActivities(159237, &backlog.GetUserActivityOptions{
	// 	ActivityTypeIDs: []int{1, 2, 3},
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("activities", activities)

	// NG
	// user, err := c.CreateUser(&backlog.CreateUserInput{
	// 	UserID:      backlog.String("hoge"),
	// 	Password:    backlog.String("aiueo9999"),
	// 	Name:        backlog.String("hoge"),
	// 	MailAddress: backlog.String("kenzo.tanaka0107@gmail.com"),
	// 	RoleType:    backlog.RoleType(backlog.RoleTypeGeneralUser),
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("id: %v, name: %v\n", *user.ID, *user.Name)
}
