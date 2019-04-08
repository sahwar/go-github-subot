# go-github-subot

A Go package to (un)subscribe to developers automatically.

## Table of Content

+ [Description](https://github.com/Dmitriy-Vas/go-github-subot#Description)
+ [Install](https://github.com/Dmitriy-Vas/go-github-subot#Install)
+ [Usage](https://github.com/Dmitriy-Vas/go-github-subot#Usage)

### Description
Bot uses github [API](https://developer.github.com/v3/users/followers/) and subscribe to developers. Bot supports few options:

+ Subscribe to all peoples.
+ On those who have *n* repositories/stars/followers.
+ With recent activity.

### Install
You can use precompiled binaries or compile subot for yourself using Go compiler.

+ Download Go from official [site](https://golang.org/).
+ Unpack Go somewhere.
+ Add Go bin folder to your PATH.
+ Clone the repo to your computer:

```
git clone https://github.com/Dmitriy-Vas/go-github-subot.git
```

+ Change directory into the project folder:

```
cd go-github-subot
```

+ Finally run the program:

```
go run main.go
```

### Usage
