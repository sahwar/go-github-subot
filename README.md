# go-github-subot

![](https://travis-ci.com/Dmitriy-Vas/go-github-subot.svg?token=u26tKHEPz6C6hydytxzK&branch=master)

A Go package to (un)subscribe to developers automatically.

## Table of Content

+ [Description](https://github.com/Dmitriy-Vas/go-github-subot#Description)
+ [Install](https://github.com/Dmitriy-Vas/go-github-subot#Install)
+ [Setup](https://github.com/Dmitriy-Vas/go-github-subot#Setup)
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

After these manipulations, you can start configuring the program.

### Setup

Now that subot is installed, you will need to setup your `config.json` file. This can be done in few steps:

1. Open the project folder in file explorer.
2. Rename the file `config-sample.json` to `config.json`. (Note: Depending on your computer's settings you might not see the `.json` part of the file name)
3. Change the bot settings with your own settings.

```
{
  "token": "axmjs7eeyfvmpvcv3jjuzqv9eylm6t2cv45sqj4y", // Your GitHub personal token
  "username": "Subot"                                  // Your username to send as User-Agent
}
```

<details>
<summary>Click here If you don't know how to get your personal token</summary>

1. Open the menu in the upper right corner and click to the settings.

![](https://i.imgur.com/UdUNv2r.png)

2. Open the developer settings.

![](https://i.imgur.com/1RKyeSZ.png)

3. Navigate to personal access tokens.

![](https://i.imgur.com/U4TnHIN.png)

4. Click to the "Generate new token" button.

![](https://i.imgur.com/zFhZdXN.png)

5. Add a description and tick the "user:follow" scope.

6. Generate your token and save somewhere.
</details>

### Usage
