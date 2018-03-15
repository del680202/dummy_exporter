

This project is a simple exmaple of prometheus exporter for learning how to writing a prometheus exporter.

And it also is  a example project for practicing Golang project.

* https://prometheus.io/docs/instrumenting/writing_exporters/


# Prepare

First, install go in your environment

## Setup GO environment

```
$ mkdir -p $HOME/golang

$ export GOROOT=/usr/local/go  # Choose your go installed directory
$ export GOPATH=$HOME/golang
$ export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

# Deploy by manual

```
# Make project folder, the format is $GOPATH/src/github.com/USER_NAME/PROJECT_NAME, using it to bind github
$ USER_NAME=del680202
$ PROJECT_NAME=dummy_exporter
$ mkdir -p $GOPATH/src/github.com/$USER_NAME/$PROJECT_NAME
$ cd $GOPATH/src/github.com/$USER_NAME/$PROJECT_NAME

# Install dep to handle dependency
$ go get -v github.com/golang/dep
$ go install -v github.com/golang/dep/cmd/dep

# Init. it will create vendor,Gopkg.toml,Gopkg.lock on your proejct folder
dep init

# Edit main.go as this site
$ vim main.go

# Install dependency
$ dep ensure
# or
$ dep ensure -v
```

# Deploy from github

# Test Run

```
$ go run main.go
```

# Build & Run

```
$ go build
$ ./dummy_exporter
```

# Push Project to Github(Option)

If you deploy this project by manual, you can push it back to github by

1. Create a PROJECT in github
2. Type command as below

```
$ git init
$ git add .
$ git commit -m "first commit"
$ git remote add origin https://github.com/$USER_NAME/$PROJECT_NAME.git
$ git push -u origin master
```

When you run dummy_exporter, you can access http://127.0.0.1:8081 to see result

And you can see result as below:

```
# HELP dummy_gauge_metric This is a dummy gauge metric
# TYPE dummy_gauge_metric gauge
dummy_gauge_metric 0
# HELP dummy_gauge_vec_metric This is a dummy gauga vece metric
# TYPE dummy_gauge_vec_metric gauge
dummy_gauge_vec_metric{myLabel="hello"} 0
...
```

