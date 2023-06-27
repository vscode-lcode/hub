# 简介

使用 vscode 编辑 ssh 主机上的文件, 无需在 ssh 主机上安装任何软件

## 使用场景

```sh
ssh openwrt
# on remote_host
<>/dev/tcp/127.0.0.1/4349 bash -i
webdav://openwrt.lo.shynome.com:4349/root/
# click the above link will open vscode to edit folder
vscode://lcode.hub/openwrt.lo.shynome.com:4349/root/
```

## 视频展示

<https://github.com/vscode-lcode/hub/assets/17316043/b2c7be9e-941a-4c14-b195-7bf2102d6d14>

**特性**

- 可在无外网环境的远程主机上使用
- 最小权限原则, 只向本地主机暴露当前运行的目录
- **慢** webdav+ssh 两个都是 TCP, 时延不可避免的高

## 预先设置

### 安装

插件地址: [lcode.hub](https://marketplace.visualstudio.com/items?itemName=lcode.hub)

vscode 安装命令:

```sh
ext install lcode.hub
```

### 设置本地主机的 `~/.ssh/config` 文件, 为其添加以下内容

```conf
# ~/.ssh/config
# config for lcode
Host *
  # 转发 hub 端口
  RemoteForward 127.0.0.1:4349 127.0.0.1:4349
  # 避免多次端口转发
  # 如果你要修改连接配置的话, 使用-M选项创建新的连接不复用已有的主连接, 示例: ssh -MC user@host.com
  # 复用链接会影响文件传输, 因为流量限制是对每一条tcp连接限制的, 所以传输文件时使用-M新开一个链接就好
  ControlMaster auto
  ControlPath /tmp/ssh_control_socket_%h_%p_%r
  # 启动 lcode-hub. (注: 你也可以在其他地方启动 lcode-hub)
  LocalCommand $(ls -t ~/.vscode/extensions/lcode.hub-1.*/bin/lcode-hub | head -n 1) --hello 'vscode://lcode.hub/{{.host}}.lo.shynome.com:4349{{.path}}' >/dev/null &
  PermitLocalCommand yes
```

### 为远程主机添加 `lcode` 命令

```sh
echo "alias lcode='<>/dev/tcp/127.0.0.1/4349 bash -i -s --'" >> ~/.bashrc
source ~/.bashrc
```

更多用法的请查看 [lcode-hub](https://github.com/vscode-lcode/lcode-hub)

## 更多功能

- [x] 添加 ICON
- [x] 远程主机一行命令运行 webdav server, 基于反弹 shell 无需安装
- [x] Windows 远程主机支持, 只要支持 `bash`, `ls` 和 `dd` 即可使用
- [ ] 支持 [`vscode.dev`](https://vscode.dev) 编辑. 只要本地主机运行 [`lcode-hub`](https://github.com/vscode-lcode/lcode) 服务就行
- [ ] 修改[维基百科常用端口页面](https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers)表明 `4349` 端口已被 `vscode-lcode` 使用. (需要帮助, vps 主机 ip 不可编辑维基百科)

## 如何帮助这个项目

- 提出问题 [Create Issue](https://github.com/vscode-lcode/hub/issues), 让此项目更加完善
- 点击查看 [CONTRIBUTING.md](./CONTRIBUTING.md) 查看技术设计以便对此项目进行改进
- [点击给作者添一根头发](https://afdian.net/item?plan_id=bd853cbc03bd11ed836452540025c377)
