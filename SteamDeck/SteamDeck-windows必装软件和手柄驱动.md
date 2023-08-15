# Steam Deck 安装 windows 后必装软件和手柄驱动

<!-- #### windows 装到 tf 卡的弊端
如果经常使用 windows 的话，还是建议装在 ssd 硬盘上，因为操作系统会有比较多的随机读写，tf 卡的小文件的读写性能不太好，会有卡顿，而且频繁读写也容易造成 tf 卡损坏。 -->

<!-- hello，大家好，今天我们来聊聊给 Steamdeck 装好 windows 之后，那我们可以做哪些配置，可以让他更接近一台完美的 windows 掌机呢？毕竟折腾这些可比玩游戏有意思多了。对吧。话不多说，直接开始吧。 -->

### 设置自动登录

<!-- 如果你有给账号设置过密码的话，那么每次开机或者睡眠模式唤醒后都需要登录，这个我们可以把它设置为自动登录，更符合掌机的使用习惯。 -->
**开机时的自动登录**

[下载 AutoLogin](https://learn.microsoft.com/en-us/sysinternals/downloads/autologon) ，打开 `Autologon.exe`，输入密码，点击 enable。

> 注意：这里输入密码是不会进行校验正确性的。

**取消睡眠唤醒后的自动登录**

打开 `设置 - 搜索 ”登录“ - 打开 ”登录选项“ - 需要登录改为“从不”`

### 打开触摸键盘

设置 - 设备 - 输入 - 开启 “不处于平板电脑模式且为连接键盘时显示触摸键盘”

### 关闭休眠
windows 有睡眠和休眠两种
睡眠是类似 SteamOS 系统待机的效果，按下电源键不会关机，是吧数据保存在内存中，下次再按电源键开启的时候就能很快恢复。
休眠是要断电的，是把数据保存到硬盘后就关机了，唤醒的时候要走一遍开机流程，这就不是我们想要的了，所以要把它禁用掉。

cmd 执行（需管理员权限）
```
Powercfg /h off
```

### 禁用Windows系统更新

因为之前有出现过更新系统后破坏了驱动的情况，所以这里把它禁用掉。

我们下载一个叫 [inControl](https://www.softpedia.com/get/Tweak/System-Tweak/Gibson-InControl.shtml)的软件。打开 exe，点击 take control，就表示禁用更新，再次点击 release control，就表示启用更新。

### windows 时间校准
在 windows 下校准过时间之后，如果打开 SteamOS 系统，SteamOS 又会改回去，切换回 windows 又是有问题的。
原因在于 SteamOS 和 windows 他们的时间都是从主板上读取的， SteamOS 和 windows 对于主板上对时间的解析不一致，SteamOS 属于 linux 操作系统，linux 普遍都是保存 utc 的时间，就是 0 时区的时间。系统展示的时候会根据时区转换为本地时间。但是 windows 不会转换时区，而是直接作为本地时间展示。这就造成了时间展示的差异。我们修改的方法也很简单，告诉 windows 主板上的时间是 utc 的时间，不要直接作为本地时间展示。

cmd 执行：
```
reg add "HKEY_LOCAL_MACHINE\System\CurrentControlSet\Control\TimeZoneInformation" /v RealTimeIsUniversal /d 1 /t REG_DWORD /f
```

然后重启，就发现时间正确显示了

### Edge 浏览器禁用后台驻留
edge - 设置 - 左上角更多 - 系统和性能 - 关闭 “在 Microsoft Edge 关闭后继续运行后台扩展和应用”

### 关机时强制关闭无响应的程序

有时候，我们点了关机，会提示程序无响应，然后系统就卡在那里，需要手动点击强制关闭才行，这是因为操作系统不能让程序正常退出了，而对于一台游戏机来说，这个功能没什么必要。
打开网址：https://www.tenforums.com/tutorials/97821-turn-autoendtasks-restart-shut-down-sign-out-windows-10-a.html
下载 AutoEndTasks 两个文件
打开看看，就是设置一个值为1，另一个是关闭


### 如何修改默认启动盘
我们在装好 windows 后，系统开机就默认进入 windows 了，但是如果想要把默认系统改成 SteamOS，要怎么做呢。这个的话我在 bios 里没发现修改的地方，目前就只发现一种傻瓜式的方式修改，拔出 tf 卡启动，就会默认启动 SteamOS，而且后续都会以 SteamOS 作为默认系统启动，只要启动一次 windows 后，后续又会改为默认 windows 启动


## 按键驱动

> 默认情况下，windows 对 steamdeck 的按键识别的不太好，我比较了目前主流的 3 种控制器驱动软件，SWICD，steam deck tools，handheld companion。接下来我们来聊聊控制器驱动装哪个比较好。


### SWICD
先来聊聊 swicd。他会虚拟出一个 xbox 360 的手柄，将 steam deck 的按键映射到这个虚拟的手柄上。在游戏中的话就是会显示连了一个 xbox360 的手柄。

安装的话我们需要打开 swicd 的官网。https://github.com/mKenfenheuer/steam-deck-windows-usermode-driver/wiki/Installation

但是这个的话我觉得已经不是很好用了，因为把游戏以非 steam 游戏的方式添加到 steam 库中，也能达到一样的效果。


根据安装教程一步步执行。


### Steam deck tools

steam deck tools 除了能实现 swicd 一样的功能，而且还能实现类似 SteamOS 系统对风扇转速，游戏 fps，功耗这些的控制。

github 上下载绿色包，解压就能用的。文件夹内有四个 exe 文件，分别对应风扇控制，xxx，xxx，手柄控制。但是第一次启动的时候会提示缺少依赖库，但是别慌，他问我们是否安装这个缺失的依赖库，点击确定。

装完之后，我们运行这四个 exe 文件。启动后，在小图标这边会显示红色，鼠标移上去提示缺少 RTSS。右键有一个 install missing RTSS 的选项，我们点击进去。页面拉到最后，点击 Final 版本的下载。

下载好后双击安装，安装完后启动 RTSS，

use kernel drivers：作用是查看cpu使用率和风扇转速，但是有可能会被判断成外挂。如果是联机游戏的话要考虑一下。


Steam 控制器
第一项是只使用 360，ps4 的控制器
第二个是如果打开steam 就使用 steam 控制器，但是必须禁用 steam 的桌面模式
第三个是忽略 steam 的影响，就是要自己处理按键冲突的问题。

### handheldCompanion

这个工具和 Steam deck tools 差不多，像功耗控制，fps，控制器按键映射，该有的功能也都有，但是他还多了一个漂亮的 GUI 界面，支持组合键等等。但是相比于 Steam deck tools，这个要花一些学习成本，而且他每个游戏都需要单独配置按键映射，我装完 handheldCompanion 之后也是花了两个小时才能在游戏中正确识别手柄。感觉就是需要手动配置的地方有点多，不像 steam deck tools 那样开箱即用。但是他的功能和自由度方面确实比 steam deck tools 要高一些。哦，还有一个原因，handheldCompanion 的安装可能需要网络好点才行，他在 0.16 版本之后不提供离线安装的方式了。安装过程中会去 github 还有一些国外网站下载文件的，所以对网络有点要求。

