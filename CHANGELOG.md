# Change Log

All notable changes to the "hub" extension will be documented in this file.

Check [Keep a Changelog](http://keepachangelog.com/) for recommendations on how to structure this file.

## [Release]

### 2.0.5

- 同步 go mod 版本分支

### 2.0.4

- 修复错误发布的问题

### 2.0.1

- 修复服务二次启动导致的输出霸屏

### 2.0.0

替换掉 [lcode-hub](https://github.com/vscode-lcode/lcode-hub), 基于 `bash` 的 `webdav` 终究是不够可靠, 还是使用反向代理的模式稳定

### 1.0.6

- 更新到 lcode-hub@v2.1.8. 修复编辑目标是文件夹是以/结尾时无法通过路径结尾不带/的路径访问 (因为 vscode 第一次访问是路径末尾不带/访问)

### 1.0.5

- 更新到 lcode-hub@v2.1.7. 修复: namespace 含大写字母时无法访问的问题

### 1.0.4

- 修复 Readme `ssh config` 错误配置说明

### 1.0.3

- 更新到 lcode-hub@v2.1.6, 避免服务器上的恶意程序通过网络权限探测&获取权限外的文件

### 1.0.0

### 0.3.0

- 新增: 支持浏览器打开
-

### 0.2.0

- change: now opened lcode dir is not appear in recent entry, beacuse we can not open it after the remote host exit lcode
- add: support open file
