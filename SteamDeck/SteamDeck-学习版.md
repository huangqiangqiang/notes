# proton

hello，大家好，这期我们来聊聊如何在 Steam Deck 上面玩学习版游戏。说到玩学习版，大家肯定听过 Proton，网络上的其他教程也会让我们安装一个叫 ProtonUp-QT 的软件，那这个软件是不是必须要装的呢，不装可不可以呢，下面我来为大家解答一下。

不知道大家有没有一个疑问，windows 的游戏怎么在 linux 上跑呢？这不得不提到 Wine，Wine 是一个开源项目，他的作用是在 linux 上运行 windows 应用程序和游戏，Wine 这种工具就是一种兼容层。valve 基于 Wine 也开发了自己的兼容层，叫做 Proton。但是 Proton 通常会落后和与 Wine 数个版本，所以也有一些非官方的分支版本，Proton-GE 就是其中一个。Proton-GE 相对于 Proton，他的版本更新一点，可能会改善对某些游戏的兼容性。

好，了解了这些兼容层的关系之后呢，我们来看一下怎么安装第三方的兼容层，我们可以在系统自带的应用商店里面搜索 ProtonUp-QT 这个软件，这个软件其实只是一个下载器，用来下载第三方的兼容层的，打开之后呢我们点击 Add Version，添加一个版本，然后我们可以选择一个兼容层，我们点开可以看到这里除了Proton-GE，还有别的兼容层，比如 Boxtron，他是用来在 linux 系统上运行 DOS 游戏的，我们常用的还是 GE 这个版本，我们这里选择 GE，然后选择一个版本安装，一般都是选最新的版本就可以了，没有必要每个都装。安装完之后，这后面有一个 unused 标记，意思是我们现在还没有哪个游戏用这个版本的兼容层打开的，现在我们不用管。然后我们重启steam。只有重启后，这个版本才会在 Steam 的兼容性列表中出现。


# 怎样玩学习版游戏

好，我们接下来说说学习版的游戏怎么运行。

主要分三步：
1. 拷贝游戏到 Steam Deck。这个具体怎么操作我这里就不多说了。我之前讲过一个 SSH 传文件的方式，那个方式的好处就是不用装任何别的软件，只要启动系统自带的服务就行了。感兴趣的同学可以查看我之前的视频，也很简单。

2. 然后是把游戏添加到 Steam 库中，点击 添加游戏 -> 添加非 Steam 游戏 -> 选择游戏的 exe 文件 -> 添加所选程序，就行了

3. 选中添加的程序，点击设置 -》属性-》兼容性-》强制使用特定兼容性工具，然后在下面的兼容性工具中选择一个，GE 开头的是我们刚才安装的第三方兼容层，下面是官方的 Proton 兼容层，其中有一个叫 Proton Experimental，这是 proton 的永久测试分支，他会比一般版本更快包含一些新功能和 bug 修复。


2020年12月，valve 官方发布了一个 Proton Experimental，是 proton 的永久测试分支，他会比一般版本更快包含一些新功能和 bug 修复。

然后就能打开了。我自己测试了10来款游戏，都是能打开的。


# 游戏如何改中文

然后下面来聊聊学习版游戏如何改中文，我测试的10来个游戏里，有一些是一打开就是中文的，还有一些是游戏内的设置选项可以改中文，这个比较简单。还有一些是游戏打开是英文，然后设置项里面没有语言的选项的，那么可以试试我这种方法。

1. 第一种是找有没有 language.txt 文件，这个是很多游戏改语言的常见的配置文件，
2. 在游戏目录中找 steam_emu.ini 或者是其他 *.ini 文件，打开看看里面有没有 language 相关的配置



