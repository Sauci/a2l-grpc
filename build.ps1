Param($Target = $Env:Target)

switch ($Target)
{
    "build"
    {
        New-Item -ItemType Directory -Force -Path bin
        $Env:GOOS="windows";$Env:GOARCH="amd64";go build -o bin/dbc.exe ./cmd/dbc.go
    }
    "test"
    {
        $env:GOMAXPROCS = 1; $env:GOOS = "windows"; $env:GOARCH = "amd64"; go test -v -coverprofile cover.out -timeout 0s ./pkg/...
        $code = $LASTEXITCODE
        go tool cover -html cover.out -o index.html
        if ($LASTEXITCODE -eq 0)
        {
            $LASTEXITCODE = $code
        }
    }
    "lint"
    {
        golint -set_exit_status ./pkg/... | Out-File -FilePath lint.log
    }
}
exit $LASTEXITCODE