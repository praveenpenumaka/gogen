# Gogen

GoGen is a scaffolding tool for go based services

# Installation

```
git clone https://github.com/praveenpenumaka/gogen
cd gogen
make install
```

# Quick start

1. Go to the directory where you want to create the project

```sh
cd $GOPATH/src/github.com/praveenpenumaka/
```

2. Create a new project 

```sh
gogen new testproject
```

This will create a project folder with project manifest file

3. change to project directory

```sh
cd testproject
```


4. Generate project folders/files

```sh
gogen generate
```

5. Go modules fetch

```sh
make once
```

## Adding a new crud route

```sh
gogen generate crud Product
```

## Running migration

```sh
make migrate
```

## Running the server
```sh
make run-api
```

# TODO

* ~~Config~~
    - ~~send basepath as argument~~
* ~~AppContext~~
* ~~Router~~
    - ~~CRUD~~
    - Generic routes
* ~~Makefile~~
* ~~Models~~
    - gorm types for models
* ~~Controllers~~
* ~~Main~~
* ~~Command line~~
* Authentication
    * Basic authentication
        * ~~Login~~
        * ~~Auth middleware~~
        * register
    * Phone OTP authentication
    * Slack authentication
    * Google authentication
    * Facebook authentication
* Swagger docs
* FileUpload
* Static file server
* Logging
* Metric tracer
* Cache
* Worker
* EventHandler
* Readme file
* Dockerfile