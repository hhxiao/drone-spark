---
date: 2016-01-01T00:00:00+00:00
title: Spark
author: hhxiao
tags: [ notifications, chat ]
repo: github.com/hhxiao/drone-spark
image: hhxiao/drone-spark
---

The Spark plugin posts build status messages to your spark account. The below pipeline configuration demonstrates simple usage:

```
pipeline:
  spark:
    image: hhxiao/drone-spark
    auth_token: ####################
    person_email: john.doe@foo.bar
```

Example configuration that posting message to a spark room:
```
pipeline:
  spark:
    image: hhxiao/drone-spark
    auth_token: ####################
    room_name:  "Build Status"
```

Example configuration with file attachment:

```diff
pipeline:
  spark:
    image: hhxiao/drone-spark
    auth_token: ####################
    person_email: john.doe@foo.bar
+   attachment: https://unsplash.it/256/256/?random
```

Example configuration with local file attachment:

```diff
pipeline:
  spark:
    image: hhxiao/drone-spark
    auth_token: ####################
    person_email: john.doe@foo.bar
-   attachment: https://unsplash.it/256/256/?random
+   attachment: ./unit_test_report.html
```

Example configuration with a customized template:

```diff
pipeline:
  spark:
    image: hhxiao/drone-spark
    auth_token: ####################
    person_email: john.doe@foo.bar
+   template: |
+     {{ #success build.status }}
+       build {{ build.number }} succeeded. Good job.
+     {{ else }}
+       build {{ build.number }} failed. Fix me please.
+     {{ /success }}
```

# Secrets

The Spark plugin supports reading credentials from the Drone secret store. This is strongly recommended instead of storing credentials in the pipeline configuration in plain text.

```diff
pipeline:
  spark:
    image: hhxiao/drone-spark
-   auth_token: ####################
    person_email: john.doe@foo.bar
```

# Parameter Reference

**auth_token**
: spark bot token 

**room_id**
: choose the room id this integration will post as

**room_name**
: choose the room name this integration will post as

**person_email**
: choose the user email this integration will post as

**template**
: overwrite the default message template

# Template Reference

**repo.owner**
: repository owner

**repo.name**
: repository name

**repo.link**
: repository link

**build.status**
: build status type enumeration, either `success` or `failure`

**build.event**
: build event type enumeration, one of `push`, `pull_request`, `tag`, `deployment`

**build.number**
: build number

**build.commit**
: git sha for current commit

**build.message**
: git message for current commit

**build.branch**
: git branch for current commit

**build.tag**
: git tag for current commit

**build.ref**
: git ref for current commit

**build.author**
: git author for current commit

**build.email**
: git author email for current commit

**build.link**
: link the the build results in drone

**build.created**
: unix timestamp for build creation

**build.started**
: unix timestamp for build started

# Template Function Reference

**uppercasefirst**
: converts the first letter of a string to uppercase

**uppercase**
: converts a string to uppercase

**lowercase**
: converts a string to lowercase. Example `{{lowercase build.author}}`

**datetime**
: converts a unix timestamp to a date time string. Example `{{datetime build.started}}`

**success**
: returns true if the build is successful

**failure**
: returns true if the build is failed

**truncate**
: returns a truncated string to n characters. Example `{{truncate build.sha 8}}`

**urlencode**
: returns a url encoded string

**since**
: returns a duration string between now and the given timestamp. Example `{{since build.started}}`
