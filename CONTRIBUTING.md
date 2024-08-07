## 技术设计

想法: 通过 webdav 编辑远程主机上的文件

1. 远程主机 Agent 启动 webdav 服务. [`lcode`](https://github.com/vscode-lcode/lcode)
2. 本地主机使用 vscode 编辑 webdav. [`webdav`](https://github.com/vscode-lcode/webdav)
3. 本地主机通过 hub 连接 vscode 和 webdav 服务. [Hub](https://github.com/vscode-lcode/hub)

## 技术实现

### 连接 vscode 和 webdav 服务. Hub

hub 仓库: <https://github.com/vscode-lcode/hub>

- 启动 [`lcode-hub`](https://github.com/vscode-lcode/hub), 等待反弹 shell 的接入

### 启动 webdav 服务的反弹 shell

- 通过`ssh -R 4349:127.0.0.1:4349`将本地主机的 Hub 服务端口 `4349` 转发到远程主机 `127.0.0.1:4349`,
  这样就能建立起与 Hub 服务的连接
- 通过 `lcode -c 127.0.0.1/4349` 连接本地主机的 [`lcode-hub`](https://github.com/vscode-lcode/hub), 提供 webdav 服务便于编辑文件

### 使用 vscode 编辑 webdav.

webdav editor 仓库: <https://github.com/vscode-lcode/web>

这一块其实可以使用现成插件[Remote Workspace](https://marketplace.visualstudio.com/items?itemName=Liveecommerce.vscode-remote-workspace)来完成, 但是该插件默认设置不支持中文出现乱码,
并且因为支持的协议太多所以不支持浏览器, 所以我新创建了一个插件满足项目所需
