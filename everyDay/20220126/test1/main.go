package main

import (
	"encoding/json"
	"fmt"
)

type LiveStreamPlatform string

const (
	LiveStreamPlatformYoutube    LiveStreamPlatform = "youtube"
	LiveStreamPlatformYealink    LiveStreamPlatform = "yealink"
	LiveStreamPlatformAliyun     LiveStreamPlatform = "aliyun"
	LiveStreamPlatformThirdParty LiveStreamPlatform = "thirdParty"
)

type InviteUserEventBody struct {
	MeetingNum         string             `json:"meetingNum"`
	LiveStreamPlatform LiveStreamPlatform `json:"livePlatform,omitempty"`
	Count              int                `json:"count,omitempty"`
}

func main() {
	i := &InviteUserEventBody{
		MeetingNum:         "123424",
		LiveStreamPlatform: LiveStreamPlatformYealink,
		Count:              10,
	}

	i2 := &InviteUserEventBody{
		MeetingNum:         "123424",
		LiveStreamPlatform: "",
		Count:              0,
	}

	stuByts1, _ := json.Marshal(i)

	stuByts2, _ := json.Marshal(i2)

	fmt.Println("stu1:", string(stuByts1))
	fmt.Println("stu2:", string(stuByts2))

}
