setlocal
  @rem SET PATH=E:\bin\TDM-GCC-64\bin;%PATH%
  SET PATH=E:\bin\mingw\mingw64\bin;%PATH%

  @rem go env -w GOARCH=386
  go env -w GOARCH=amd64
  

  @rem https://golang.org/pkg/cmd/go/internal/clean/
  @rem go clean

  @rem go build -ldflags -H=windowsgui

  @rem go get testingconfig@latest
  @rem go get testinglogger@latest

  set /p Ver=<.version
  set Patch=%date%_%time%

  @rem "-X 'alcogo3/core.NumVersion=%Ver%' -X 'alcogo3/core.Created=%Patch%'"
  go build -ldflags="-H=windowsgui -s -w" -o app_gui64.exe ./gui64/app 
  go build -o app_gui64console.exe ./gui64/app 

  go build -ldflags "-s -w" -o app_cmd64_pack.exe ./cmd64/app 
  go build -o app_cmd64.exe ./cmd64/app 

endlocal
