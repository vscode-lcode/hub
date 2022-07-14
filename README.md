# 简介

使用 vscode 编辑 ssh 主机上的文件

## 使用场景

```sh
ssh your_host
# on remote_host
lcode ~
vscode://lcode.hub/shy-drone-f0_f0_f0_f0_f0_f0/dav/root
# click the above link will open vscode to edit folder
```

**特性**

- 可在无外网环境的远程主机上使用
- 最小权限原则, 只向本地主机暴露当前要编辑的文件/目录
- **慢** webdav+ssh 两个都是 TCP, 时延不可避免的高

## 预先设置

### 设置本地主机的 `~/.ssh/config` 文件, 为其添加以下内容

```conf
# ~/.ssh/config
# config for lcode
Host *
  # 转发 hub 端口
  RemoteForward 127.0.0.1:4349 127.0.0.1:4349
  # 避免多次端口转发
  ControlMaster auto
  ControlPath /tmp/ssh_control_socket_%lcodeh_%p_%r
  # ignore `connect_to 127.0.0.1 port 4349: failed.`
  LogLevel FATAL
```

### 为远程主机添加 `lcode` 命令

```sh
wget -O /usr/local/bin/lcode https://github.com/vscode-lcode/lcode/releases/download/v0.0.4/lcode && chmod +x /usr/local/bin/lcode
```

## 更多功能

- [x] 添加 ICON
- [ ] 远程主机一键安装脚本
- [ ] 设置: 监听端口选项
- [ ] Windows 远程主机支持并测试
- [ ] 支持 [`vscode.dev`](https://vscode.dev) 编辑. 只要本地主机运行 [`lcode-hub`](https://github.com/vscode-lcode/hub) 服务就行
- [ ] 修改[维基百科常用端口页面](https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers)表明 `4349` 端口已被 `vscode-lcode` 使用

## 如何帮助这个项目

- 提出问题 [Create Issue](https://github.com/vscode-lcode/pack/issues), 让此项目更加完善
- 点击查看 [CONTRIBUTING.md](./CONTRIBUTING.md) 查看技术设计以便对此项目进行改进
- [点击给作者添一根头发](https://afdian.net/item?plan_id=bd853cbc03bd11ed836452540025c377)
