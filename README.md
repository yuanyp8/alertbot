# alertbot
receive alertmanaer data and push

## Installation
```shell
# 依赖golang环境
# TODO release github packages

git clone https://github.com/yuanyp8/alertbot
cd alertbot
go build -o alertbot main.go
cp alertbot /usr/bin
```

## Start with Systemd Service
```shell
cat  > /etc/systemd/system/alertbot.service << 'EOF'
[Unit]
Description=https://github.com/yuanyp8/alertbot

[Service]
Restart=on-failure
ExecStart=/usr/bin/alertbot start -f {{ $YOUR_PATHNAME }}/alertbot/etc/config.yaml
ExecReload=/bin/kill -HUP $MAINPID
Restart=always
[Install]
WantedBy=multi-user.target
EOF
```
enable 
```shell
systemctl enable alertbot
```

start 
```shell
systemctl start alertbot
```
stop 
```shell
systemctl stop alertbot
```

## configuration
your can modify config across the etc/*.yaml file to reset your receiver and publisher
