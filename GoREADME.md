### This is a fun directory of some Go projects I found on youtube

#### Notes
For a new Go environment: 
    1. create a go.mod file in your directory
        go mod init go-server

    2. Tidy up dependencies:
        go mod tidy

    3. Build command 
    go build

    4. Run command
    go run main.go


Cool way to get packages from a github (example github.com/gorilla/mux) go-movies-crud project
    1. create go.mod file 
        go mod init go-movies-crud
    2. go get "github.com/gorilla/mux"   
        result: **stored in a go.sum file**
            go: downloading github.com/gorilla/mux v1.8.1
            go: added github.com/gorilla/mux v1.8.1
