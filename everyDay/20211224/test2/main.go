package main

import "fmt"

func main() {
	count := 1212
	users := make([]int, 0, count)
	for i := 0; i < count; i++ {
		users = append(users, i)
	}

	fmt.Printf("%v\n", users)

	const limitPushCount = 50
	userCount := len(users)
	for i := 0; i < userCount; i += limitPushCount {
		j := i + limitPushCount
		if j > userCount {
			j = userCount
		}

		fmt.Printf("%v\n", users[i:j])
	}
}
