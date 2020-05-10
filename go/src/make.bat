set GOOS=windows
set GOARCH=amd64
go build -o loki.exe
move loki.exe ..\bin
