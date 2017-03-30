# wx-devdemo

# Prerequisites

Docker 1.12+

Golang 1.8+

make

git

# Clone

```sh
git clone https://github.com/zyfdegh/wx-devdemo
```

# Build

```sh
make dep-init
make dep-update

make build
```

# Run

```sh
export TOKEN=some-wechat-token
export APPID=some-wechat-appid
export SECRET=some-wechat-appsecret
# optinal, token refresh period in seconds
# export POLLING_SEC=7100

make run
```
