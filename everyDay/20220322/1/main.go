package main

import "fmt"

type InviteUserParam struct {
	SubjectIDs map[string]string
	Contacts   []string
}

func initInvParam() *InviteUserParam {
	newInvParam := new(InviteUserParam)
	newInvParam.SubjectIDs = make(map[string]string)
	// newInvParam.Contacts = make([]string, 0)
	return newInvParam
}

func main() {
	sipInvParam := initInvParam()
	sipInvParam.SubjectIDs["a.UserNumber"] = "a.SubjectID"
	sipInvParam.Contacts = append(sipInvParam.Contacts, "a.UserNumber")
	fmt.Printf("%v", sipInvParam)

	v := new(int)
	*v = 2
	println(5 / -+*v)
}
