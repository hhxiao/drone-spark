package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"os"
)

var build = "0"

func main() {
	app := cli.NewApp()
	app.Name = "spark plugin"
	app.Usage = "spark plugin"
	app.Action = run
	app.Version = fmt.Sprintf("1.0.%s", build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "auth_token",
			Usage:  "spark bot auth token",
			EnvVar: "PLUGIN_AUTH_TOKEN,SPARK_AUTH_TOKEN",
		},
		cli.StringFlag{
			Name:   "room_id",
			Usage:  "spark room id",
			EnvVar: "PLUGIN_ROOM_ID,SPARK_ROOM_ID",
		},
		cli.StringFlag{
			Name:   "room_name",
			Usage:  "spark room name",
			EnvVar: "PLUGIN_ROOM_NAME,SPARK_ROOM_NAME",
		},
		cli.StringFlag{
			Name:   "person_email",
			Usage:  "spark person email",
			EnvVar: "PLUGIN_PERSON_EMAIL,SPARK_PERSON_EMAIL",
		},
		cli.StringFlag{
			Name:  "template",
			Usage: "spark message template",
			Value: `###Build for [{{repo.owner}}/{{repo.name}}]({{repo.link}}) {{build.status}}!!!
####Build Details
{{#if build.email}}
* **Author:** <@personEmail:{{build.email}}>
{{else}}
* **Author:** {{build.author}}
{{/if}}
* **Branch:** {{build.branch}}
* **Build No:** [{{build.number}}]({{build.link}})
* **Commit:** [{{build.commit}}]({{repo.link}}/commit/{{build.commit}})
* **Event:** {{build.event}}
* **Message:** {{build.message}}
`,
			EnvVar: "PLUGIN_TEMPLATE,SPARK_TEMPLATE",
		},
		cli.StringFlag{
			Name:   "attachment",
			Usage:  "spark message attachment",
			EnvVar: "PLUGIN_ATTACHMENT,SPARK_ATTACHMENT",
		},
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "repo.link",
			Usage:  "repository link",
			EnvVar: "DRONE_REPO_LINK",
		},
		cli.StringFlag{
			Name:   "build.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA",
		},
		cli.StringFlag{
			Name:   "build.ref",
			Value:  "refs/heads/master",
			Usage:  "git commit ref",
			EnvVar: "DRONE_COMMIT_REF",
		},
		cli.StringFlag{
			Name:   "build.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "build.author",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "build.email",
			Usage:  "git author email",
			EnvVar: "DRONE_COMMIT_AUTHOR_EMAIL",
		},
		cli.StringFlag{
			Name:   "build.message",
			Value:  "",
			Usage:  "build message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
		cli.Int64Flag{
			Name:   "build.started",
			Usage:  "build started",
			EnvVar: "DRONE_BUILD_STARTED",
		},
		cli.Int64Flag{
			Name:   "build.created",
			Usage:  "build created",
			EnvVar: "DRONE_BUILD_CREATED",
		},
		cli.StringFlag{
			Name:   "build.tag",
			Usage:  "build tag",
			EnvVar: "DRONE_TAG",
		},
		cli.Int64Flag{
			Name:   "job.started",
			Usage:  "job started",
			EnvVar: "DRONE_JOB_STARTED",
		},
		cli.StringFlag{
			Name:  "env-file",
			Usage: "source env file",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	if c.String("env-file") != "" {
		_ = godotenv.Load(c.String("env-file"))
	}

	plugin := Plugin{
		Repo: Repo{
			Owner: c.String("repo.owner"),
			Name:  c.String("repo.name"),
			Link:  c.String("repo.link"),
		},
		Build: Build{
			Tag:     c.String("build.tag"),
			Number:  c.Int("build.number"),
			Event:   c.String("build.event"),
			Status:  c.String("build.status"),
			Commit:  c.String("build.sha"),
			Ref:     c.String("build.ref"),
			Branch:  c.String("build.branch"),
			Author:  c.String("build.author"),
			Email:   c.String("build.email"),
			Link:    c.String("build.link"),
			Message: c.String("build.message"),
			Started: c.Int64("build.started"),
			Created: c.Int64("build.created"),
		},
		Job: Job{
			Started: c.Int64("job.started"),
		},
		Config: Config{
			AuthToken:   c.String("auth_token"),
			RoomId:      c.String("room_id"),
			RoomName:    c.String("room_name"),
			PersonEmail: c.String("person_email"),
			Template:    c.String("template"),
			Attachment:  c.String("attachment"),
		},
	}
	return plugin.Exec()
}
