# slack
[![Build Status](https://travis-ci.org/joaosoft/slack.svg?branch=master)](https://travis-ci.org/joaosoft/slack) | [![codecov](https://codecov.io/gh/joaosoft/slack/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/slack) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/slack)](https://goreportcard.com/report/github.com/joaosoft/slack) | [![GoDoc](https://godoc.org/github.com/joaosoft/slack?status.svg)](https://godoc.org/github.com/joaosoft/slack)

A slack formatter that allows you to add slack on your output.
The easy way to use the slack:
``` Go
import log github.com/joaosoft/slack

```

###### If i miss something or you have something interesting, please be part of this project. Let me know! My contact is at the end.

## With support for
* text format

## Dependecy Management 
>### Dep

Project dependencies are managed using Dep. Read more about [Dep](https://github.com/golang/dep).
* Install dependencies: `dep ensure`
* Update dependencies: `dep ensure -update`

>### Go
```
go get github.com/joaosoft/slack
```

## Usage 
This examples are available in the project at [slack/examples](https://github.com/joaosoft/slack/tree/master/examples)

```go
func main() {
	fmt.Fprintf(os.Stdout, fmt.Sprintf("%s joao", slack.WithSlack("hello", slack.FormatBold, slack.ForegroundRed, slack.BackgroundCyan)))
}
```

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
