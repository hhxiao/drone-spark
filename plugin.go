package main

import (
	"fmt"
	"github.com/hhxiao/drone-spark/spark"
	"strings"
)

type (
	Repo struct {
		Owner string
		Name  string
		Link  string
	}

	Build struct {
		Tag     string
		Event   string
		Number  int
		Commit  string
		Ref     string
		Branch  string
		Author  string
		Email   string
		Status  string
		Link    string
		Message string
		Started int64
		Created int64
	}

	Config struct {
		AuthToken   string
		RoomId      string
		RoomName    string
		PersonEmail string
		Template    string
		Attachment  string
	}

	Job struct {
		Started int64
	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Config Config
		Job    Job
	}
)

// Exec executes the plugin function
func (p Plugin) Exec() error {
	payload := spark.Message{
		ToPersonEmail: p.Config.PersonEmail,
		Text:          message(p.Repo, p.Build),
		Files:         p.Config.Attachment,
	}

	if p.Config.Template != "" {
		txt, err := RenderTrim(p.Config.Template, p)
		if err != nil {
			return err
		}
		payload.Markdown = txt
	}
	client := spark.NewClient(p.Config.AuthToken)

	if p.Config.RoomName != "" {
		room, err := client.FindRoomByName(p.Config.RoomName)
		if err != nil {
			return err
		}
		payload.RoomId = room.Id
	}

	if p.Config.Attachment == "" || strings.HasPrefix(p.Config.Attachment, "http") {
		return client.PostMessage(&payload)
	} else {
		return client.PostFileMessage(&payload)
	}
}

func message(repo Repo, build Build) string {
	return fmt.Sprintf("Build for %s/commit/%s (%s) %s by %s",
		build.Link,
		build.Commit,
		build.Branch,
		build.Status,
		build.Author,
	)
}
