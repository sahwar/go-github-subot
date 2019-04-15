set GOARCH=386
go build -o WindowsSubot.exe main.go
set GOOS=darwin
go build -o DarwinSubot main.go
set GOOS=linux
go build -o LinuxSubot main.go