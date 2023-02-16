# slack
[![Build Status](https://travis-ci.org/joaosoft/slack.svg?branch=master)](https://travis-ci.org/joaosoft/slack) | [![codecov](https://codecov.io/gh/joaosoft/slack/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/slack) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/slack)](https://goreportcard.com/report/github.com/joaosoft/slack) | [![GoDoc](https://godoc.org/github.com/joaosoft/slack?status.svg)](https://godoc.org/github.com/joaosoft/slack)

A slack formatter that allows you to send messages to slack.
The easy way to use the slack:
``` Go
import log github.com/joaosoft/slack
```

###### If i miss something or you have something interesting, please be part of this project. Let me know! My contact is at the end.

## With support for
* messages
* attachments

## Usage 
This examples are available in the project at [slack/examples](https://github.com/joaosoft/slack/tree/master/examples)

```go
func main() {
    config := &slack.Config{
        Service: "Service",
        Url:     "",
    }
    
    mySlack := slack.New(config)
    
    msg := slack.NewMessage().Title("Title").
    Action("Action").
    Environment("Environment").
    Code("Code")
    
    _ = mySlack.Send(msg)
}
```

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
