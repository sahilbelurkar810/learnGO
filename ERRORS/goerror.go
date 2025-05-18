// package main

// import "fmt"

// type User struct {
// 	ID int
// }

// func main() {
// 	user, err := GetUser()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(user)

// 	profile, err := GetUserProfile(user.ID)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	} else {

// 		fmt.Println(profile)
// 	}
// }

// func GetUser() (User, error) {
// 	// Simulate a user not found error
// 	return User{}, fmt.Errorf("user not found")

// }

// func GetUserProfile(userID int) (string, error) {
// 	// Simulate a profile not found error
// 	if userID != 1 {
// 		return "", fmt.Errorf("profile for user with id %d not found", userID)
// 	}
// 	return "Profile of John Doe", nil
// }
