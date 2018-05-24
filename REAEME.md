# Robosln - robo2025 smart solution

## 使用说明

```bash
./robosln -d database_url
```

## Build

Prepare go environment

```bash
go get -u github.com/gin-gonic/gin
go get -u github.com/gin-contrib/cors
go get -u github.com/sirupsen/logrus
go get -u github.com/fatih/color
go get -u github.com/spf13/cobra
```

Install [UPX](https://upx.github.io/), In mac it is:

```bash
brew install upx
```

Build program

```
make build
```

*Note:* if encounter build error for windows, you may need `go get -u github.com/inconshreveable/mousetrap`
