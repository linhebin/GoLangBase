package main

import "fmt"

func main() {
	subjectIDsMap := make(map[string][]string)
	subjectIDsMap["key"] = append(subjectIDsMap["key"], "acc.SubjectID")
	subjectIDsMap["key"] = append(subjectIDsMap["key"], "acc.SubjectID2")
	fmt.Printf("%v", subjectIDsMap)
}
