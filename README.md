# go-github-subot

![](https://travis-ci.com/Dmitriy-Vas/go-github-subot.svg?token=u26tKHEPz6C6hydytxzK&branch=master)

A Go package to (un)subscribe to developers automatically.

## Table of Content

+ [Description](https://github.com/Dmitriy-Vas/go-github-subot#Description)
+ [Install](https://github.com/Dmitriy-Vas/go-github-subot#Install)
+ [Setup](https://github.com/Dmitriy-Vas/go-github-subot#Setup)
+ [Run](https://github.com/Dmitriy-Vas/go-github-subot#Run)
+ [TODO](https://github.com/Dmitriy-Vas/go-github-subot#TODO)

### Description
Bot uses github [API](https://developer.github.com/v3/users/followers/) and (un)subscribe to developers. Bot supports few options:

+ All GitHub developers or from *user* followers list.
+ On those who have *n* repositories/stars/followers/watchers/gists.
+ Who has an id greater than *n*.
+ By most used language.
+ With recent activity (TODO).

For more info about the parameters, please read the [setup](https://github.com/Dmitriy-Vas/go-github-subot#Setup) section. 

### Install
You can download [precompiled binaries](https://github.com/Dmitriy-Vas/go-github-subot/releases) or compile subot for yourself using Go compiler.

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

**Warning: All options below "id" in `config.json` will increase cooldown time between requests**

```
{
  // Your GitHub personal token.
  "token": "axmjs7eeyfvmpvcv3jjuzqv9eylm6t2cv45sqj4y",
  
  // Your username to send as User-Agent.
  // Please don't use default value.
  // Default is "Subot"
  "username": "Subot",
  
  // The time in milliseconds that the program should wait between requests
  // Default is 1000
  "timer": 1000,
  
  // Supports "follow" and "unfollow" mode
  // If mode is "follow", then will check users from "source" and add them to following list if they meet requirements
  // If mode is "unfollow", then will check users from following list for requirements (followers/stars/repo/etc) and unfollow them
  // Default is "follow"
  "mode": "follow",
  
  // Source to get list of developers. Write username to get user followers.
  // Works only if mode is "follow"
  // Default is "all"
  // In this example subot will get list of followed users from user with username Dmitriy-Vas
  "source": "Dmitriy-Vas",
  
  // If source is username, then starting from this page
  // Works only if mode is "follow"
  // Default is 1
  // In this example subot will start fetching users since 5 page  
  "page": 5,
  
  // If source is "all", then starting from user with this id.
  // Works only if mode is "follow"
  // Defailt is 0
  // In this example, if "source" is "all", then subot will start from a user with id more than 15000
  "id": 15000,
  
  // Check user for amount of followers
  // If mode is "follow", then subscribe to user if he have more followers than this number
  // If mode is "unfollow", then unsubscribe from user if he have less followers than this number
  // Defailt is 0
  "followers": 0,
  
  // Check user repositories and count stars
  // If mode is "follow", then subscribe to user if he have more stars than this number
  // If mode is "unfollow", then unsubscribe from user if he have less stars than this number
  // Defailt is 0
  "stars": 0,
  
  // Check amount of user public repositories
  // If mode is "follow", then subscribe to user if he have more repositories than this number
  // If mode is "unfollow", then ubsubscribe from user if he have less repositories than this number
  // Defailt is 0
  "repos": 0,
  
  // Check user repositories and count watchers
  // If mode is "follow", then subscribe to user if he have more watchers than this number
  // If mode is "unfollow", then unsubscribe from user if he have less watchers than this number
  // Defailt is 0 
  "watchers": 0,
  
  // Check for amount of gists
  // If mode is "follow", then subscribe to user if he have more gists than this number
  // If mode is "unfollow", then unsubscribe from user if he have less gists than this number
  // Default is 0
  "gists": 0,
  
  // Check user repositories and get most used language
  // If mode is "follow", then subscribe to user if he have same most used language
  // If mode is "unfollow", then unsubscribe from user if he have different most used language
  // Default is ""
  // In this example subot will follow developers who uses Java
  "language": "Java"
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

6. Generate your token and save him in config.json.
</details>

### Run

After setting up the `config.json` file, subot is ready to go. To run program, simply use the command `go run main.go` in the console or run precompiled program.
If you have setup your `config.json` properly (and used the correct credentials) you should see an output similar to this

```
asakura doesn't meet requirements! [1]
mthbyd doesn't meet requirements! [1]
koenigc doesn't meet requirements! [1]
lkesslin doesn't meet requirements! [1]
romanzolotarev with ID 29062 now following!
johndoe doesn't meet requirements! [2]
sotarok with ID 29064 now following!
dsieuquay doesn't meet requirements! [1]
polonia doesn't meet requirements! [2]
coldblooded doesn't meet requirements! [1]
Config successfully saved!

Process finished with exit code 0
```

To stop executing program use Ctrl^C.

### TODO

- [x] Add basic net functions.
- [x] Add config/user/project structures.
- [x] Add catching interrupt signal and save config.
- [x] Add filter by repos/stars/followers/watchers/gists.
- [x] Add filter by most used language.
- [x] Add custom cooldown.
- [x] Precompiled binaries.
- [x] Add method to **un**subscribe from developers.
- [ ] Add filter by recent activity.
- [ ] Optimize code. 
