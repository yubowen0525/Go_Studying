```shell script

go get -u github.com/go-delve/delve/cmd/dlv
sudo find / -name dlv
ln -s /root/go/bin/dlv $GOROOT/bin

dlv debug --headless --listen=:80 --api-version=2
```

