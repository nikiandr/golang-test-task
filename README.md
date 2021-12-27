# golang-test-task

Deployed version of app could be found here: https://golang-test-task.herokuapp.com/.

## Local deployment guide

Guide was created on Ubuntu 21.10 - based operating system. For this reason there may be some sufficient differences in local deployment process on another OS such as Windows or MacOS.

1. Install Golang

Instruction for Golang installation could be found on this page: https://go.dev/doc/install.

2. Clone this repository on your computer

```bash
git clone git@github.com:nikiandr/golang-test-task.git
```
or 
```bash
git clone https://github.com/nikiandr/golang-test-task.git
```

3. Open project directory 

```bash
cd golang-test-task/
```

4. Set PORT environment variable

In this guide we will set PORT to 8080. You can set it to whatever port you like and have opened.

```bash
export PORT="8080"
```

5. Build project

```bash
go build cmd/main.go
```

6. Run executable file

```bash
./main
```

7. Open application

For this you should open your browser on http://localhost:8080

8. ***FIN***