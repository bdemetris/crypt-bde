# Contributing

Welcome! If you're looking to help, this document is a great place to start!

## Finding things that need help

Right now this is a bit of a yolo project just because.

Try checking out `#Crypt` on MacAdmins Slack.

## Building the project

To build Crypt-BDE from source, you will need [Go 1.11](https://golang.org/dl/) or later installed.

```
go get github.com/bdemetris/crypt-bde
```

*or*

```shell
git clone git@github.com:bdemetris/crypt-bde.git && cd crypt-bde
make deps
make
```

## Git workflow
```
username=$GitHubUsername
# add your remote/upstream
git remote add $username git@github.com:$username/micromdm.git

# update from origin/master
git pull --rebase

# create a branch
git checkout -b my_feature

# push changes from my_feature to your fork.
#    -u, --set-upstream    set upstream for git pull/status
git push -u $username
```

## Go Resources

A few helpful resources for getting started with Go:

* [Writing, building, installing, and testing Go code](https://www.youtube.com/watch?v=XCsL89YtqCs)
* [Resources for new Go programmers](http://dave.cheney.net/resources-for-new-go-programmers)
* [How I start](https://howistart.org/posts/go/1)
* [How to write Go code](https://golang.org/doc/code.html)

