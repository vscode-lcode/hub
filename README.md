# hub README

```txt
vscode://lcode.hub/localhost_0e_0e_0e_0e_0e/home/username/editdir
# will transform to the webdav link
webdav://127.0.0.1:4349/proxy/localhost_0e_0e_0e_0e_0e/home/username/editdir
```

## Features

## Requirements

add port forward to your ssh config file `~/.ssh/config`

```conf
Host *
  RemoteForward 4349 127.0.0.1:4349
  # 避免多次SSH会话连接时端口冲突
  ControlMaster auto
  ControlPath /tmp/ssh_control_socket_%h_%p_%r
```

## Extension Settings

## Known Issues

## Release Notes

### 0.2.0

- change: now opened lcode dir is not appear in recent entry, beacuse we can not open it after the remote host exit lcode
- add: support open file
