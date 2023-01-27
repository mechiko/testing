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

  @rem go build ./cmd/app 

  @rem PUSHD .
  @rem  cd E:\src\Vite\quasar\cleargojs
  @rem   call npx quasar build
  @rem   POPD
  @rem   rmdir E:\src\goproj\go_clean_architech\pkg\vueapp /s /q

  @REM https://ss64.com/nt/robocopy.html
  @REM /E : Copy Subfolders, including Empty Subfolders.
  @rem   robocopy E:\src\Vite\quasar\cleargojs\dist\spa\ E:\src\goproj\go_clean_architech\pkg\vueapp /e
  
  go build -ldflags="-H=windowsgui -s -w" -o app.exe 

  go build -ldflags="-s -w" -o app_console.exe 
endlocal
