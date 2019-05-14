
# for mac os
echo building binary for mac os ...
go build -o osmdownloader main.go
zip -m osmdownloader-mac-amd64.zip osmdownloader

# for linux
echo building binary for linux ...
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o osmdownloader main.go
zip -m osmdownloader-linux-amd64.zip osmdownloader

# for windows
echo building binary for windows ...
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o osmdownloader.exe main.go
zip -m osmdownloader-windows-amd64.zip osmdownloader.exe

echo done!
