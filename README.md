# PIPE

web pipe

------------

build linux debian x32
```bash
    GOOS=linux GOARCH=386 go build
```

build linux debian x64
```bash
    GOOS=linux GOARCH=386 go build
```

build linux debian arm6
```bash
    GOOS=linux GOARCH=arm GOARM=6 go build
```

build linux debian arm7
```bash
    GOOS=linux GOARCH=arm GOARM=7 go build
```

build linux debian arm8
```bash
    GOOS=linux GOARCH=arm64 go build
```

build windows x32
```bash
    GOOS=windows GOARCH=386 go build
```

build windows x64
```bash
    GOOS=windows GOARCH=amd64 go build
```

build windows x32 invisible
```bash
    GOOS=windows GOARCH=386 go build -ldflags "-H windowsgui"
```

build windows x64 invisible
```bash
    GOOS=windows GOARCH=amd64 go build -ldflags "-H windowsgui"
```