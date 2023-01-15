# 简介

使用 vscode 编辑 ssh 主机上的文件

## 使用场景

```sh
ssh your_host
# on remote_host
lcode ~
vscode://lcode.hub/3-openwrt.lo.shynome.com:4349/root
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
  # 如果你要修改连接配置的话, 使用-M选项创建新的连接不复用已有的主连接, 示例: ssh -MC user@host.com
  # 复用链接会影响文件传输, 因为流量限制是对每一条tcp连接限制的, 所以传输文件时使用-M新开一个链接就好
  ControlMaster auto
  ControlPath /tmp/ssh_control_socket_%lcodeh_%p_%r
  #  启动 lcode-hub
  LocalCommand $(ls -t ~/.vscode/extensions/lcode.hub-1.*/bin/lcode-hub | head -n 1) -log 0 --hello 'vscode://lcode.hub/{{.host}}.lo.shynome.com:4349{{.path}}' &
  PermitLocalCommand yes
```

### 为远程主机添加 `lcode` 命令

```sh
echo "alias lcode='>/dev/tcp/127.0.0.1/4349 0> >(echo 0) 0>&1  2> >(grep -E ^lo: >&2) bash +o history -i -s -- -x'" >> ~/.bashrc
source ~/.bashrc
```

更多用法的请查看 [lcode](https://github.com/vscode-lcode/lcode)

## 更多功能

- [x] 添加 ICON
- [x] 远程主机一键安装脚本, 使用反弹 shell 无需安装
- [x] Windows 远程主机支持, 只要支持 `bash`, `ls` 和 `dd` 即可使用
- [ ] 支持 [`vscode.dev`](https://vscode.dev) 编辑. 只要本地主机运行 [`lcode-hub`](https://github.com/vscode-lcode/lcode) 服务就行
- [ ] 修改[维基百科常用端口页面](https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers)表明 `4349` 端口已被 `vscode-lcode` 使用. (需要帮助, vps 主机 ip 不可编辑维基百科)

## 如何帮助这个项目

- 提出问题 [Create Issue](https://github.com/vscode-lcode/pack/issues), 让此项目更加完善
- 点击查看 [CONTRIBUTING.md](./CONTRIBUTING.md) 查看技术设计以便对此项目进行改进
- [点击给作者添一根头发](https://afdian.net/item?plan_id=bd853cbc03bd11ed836452540025c377)
