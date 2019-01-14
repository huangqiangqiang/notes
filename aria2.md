## aria2使用

网上都说`aria2`可以代替迅雷，很牛逼的样子，就花了点时间研究了下

## 安装
[在这里](https://github.com/aria2/aria2/releases)找到对应的操作系统下载

## 使用

> 先单独为aria2建一个文件夹，下文新建的文件都放在此处。

`aria2`的使用一般需要先写一个配置文件，配置文件声明了`aria2`需要开启哪些功能，新建一个配置文件
```
aria2.conf
```

然后把一下的配置复制到配置文件中（每句配置的意思看注释）

- RPC相关的设置是为了和aria2通信，比如用web操作的时候
- session相关设置是为了关闭aria2后保存未下下载完的任务，以便下次启动后继续
- BitTorrent相关设置和BT下载相关


```
########### 基本设置 ###########
# 文件保存路径, 默认为当前启动位置
dir=/Users/xxx/Downloads

########### RPC相关设置 ###########
# 设置加密的密钥
rpc-secret=token
# 允许rpc
enable-rpc=true
# 允许所有来源, web界面跨域权限需要
rpc-allow-origin-all=true
# 允许外部访问，false的话只监听本地端口
rpc-listen-all=true
# RPC端口, 仅当默认端口被占用时修改
#rpc-listen-port=6800
# 最大同时下载数(任务数), 路由建议值: 3
max-concurrent-downloads=5
# 断点续传
continue=true
# 同服务器连接数
max-connection-per-server=5
# 最小文件分片大小, 下载线程数上限取决于能分出多少片, 对于小文件重要
min-split-size=10M
# 单文件最大线程数, 路由建议值: 5
split=10
# 下载速度限制
max-overall-download-limit=0
# 单文件速度限制
max-download-limit=0
# 上传速度限制
max-overall-upload-limit=0
# 单文件速度限制
max-upload-limit=0
# 断开速度过慢的连接
# lowest-speed-limit=0
# 验证用，需要1.16.1之后的release版本
# referer=*
# 文件缓存, 使用内置的文件缓存, 如果你不相信Linux内核文件缓存和磁盘内置缓存时使用, 需要1.16及以上版本
# disk-cache=0
# 另一种Linux文件缓存方式, 使用前确保您使用的内核支持此选项, 需要1.15及以上版本(?)
# enable-mmap=true
# 文件预分配, 能有效降低文件碎片, 提高磁盘性能. 缺点是预分配时间较长
# 所需时间 none < falloc ? trunc << prealloc, falloc和trunc需要文件系统和内核支持
# file-allocation=prealloc

########### session相关设置 ###########
input-file=/Users/xxx/.aria2/aria2.session
save-session=/Users/xxx/.aria2/aria2.session
# 定时保存会话, 0为退出时才保存, 需1.16.1以上版本, 默认:0
# save-session-interval=0

########### BitTorrent相关设置 ###########
# 启用本地节点查找
bt-enable-lpd=true
bt-max-open-files=16
# 单种子最大连接数
bt-max-peers=30
# 强制加密, 防迅雷必备
# bt-require-crypto=true
# 当下载的文件是一个种子(以.torrent结尾)时, 自动下载BT
follow-torrent=true
# BT监听端口, 当端口屏蔽时使用
# listen-port=6881-6999
# bt-tracker 更新，解决Aria2 BT下载速度慢没速度的问题
# bt-tracker=xxx

dht-file-path=/Users/xxx/.aria2/dht.dat
dht-file-path6=/Users/xxx/.aria2/dht6.dat
dht-listen-port=6801
#enable-dht6=true
listen-port=6801
max-overall-upload-limit=0K
seed-ratio=0
enable-dht=true
bt-enable-lpd=true
enable-peer-exchange=true

# 百度盘
# user-agent=netdisk;5.2.6;PC;PC-Windows;6.2.9200;WindowsBaiduYunGuanJia
# referer=http://pan.baidu.com/disk/home

```

复制进配置文件后，需要改两个地方，一个基本设置里，dir表示你需要下载到哪个目录。第二个是文件中的`/Users/xxx/.aria2`这个路径改成之前你新建`aria2`文件夹的路径。

## 在网页端管理aria2任务

打开网页
```
http://aria2c.com/
```

点击设置按钮，JSON-RPC Path改为
```
http://token:token@localhost:6800/jsonrpc
```

正常的话，这时页面就已经和`aria2`连接上了。

## BT下载配置优化

BT下载需要配置`bt-tracker`配置项，可以去[这个网页](http://www.tkser.tk/)找到具体的值。看中哪个就点击获取并生成，生成后点击复制，粘贴在`bt-tracker=`后，如
```
bt-tracker=udp://tracker.coppersurfer.tk:6969/announce,udp://tracker.internetwarriors.net:1337/announce,udp://tracker.opentrackr.org:1337/announce,
...
```

也可以去[github](https://github.com/ngosang/trackerslist)上复制

## 百度网盘下载

下载[chrome插件](https://github.com/acgotaku/BaiduExporter)，解压缩，把`.crx`结尾的文件拖入chrome浏览器自动安装扩展。

然后打开[百度网盘](https://pan.baidu.com)，随便选中一个文件，上面会有一个导出下载的按钮，选择`aria2 rpc`方式，便会自动加入到`aria2`下载。

## 总结

实际上aria2在大部分情况下，下载速度比不过迅雷的，很有可能迅雷速度好几M，aria2却是0，只是aria2的优势是可以离线下载，就是安装在路由器上，启动一个下载任务，等过段时间去看，已经下完了，还有就是迅雷会屏蔽一些资源，限速这些，aria2却不受限制。