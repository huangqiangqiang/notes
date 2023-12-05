# SteamDeck 必备技能：使用 SSH 传输文件

> hello，大家好，我是xxx，今天我们来了解一下如何使用 ssh 的方式将文件从 PC 传输到 Steam Deck。使用 ssh 的好处是无需安装任何其他的软件，是系统自带的，只是默认没有开启。


#### 1. 设置 deck 用户的密码

因为默认登录的账号 deck 没有设置密码，我们需要手动设置，打开 `konsole` 执行以下命令：

```
passwd
```

> 输入密码时，密码是不会显示出来的

#### 2. 启动 sshd

```
systemctl enable sshd
```

可以使用一下命令查看是否启动成功
```
systemctl status sshd
```

然后我们转移到 PC 端，通过 ssh 登录到 Steam Deck，命令是

```
# PC 必须是和 Steam Deck 连接同一个网络
ssh deck@<Steam Deck IP>

# 例：
ssh deck@192.168.0.10
```

接下来就会让我们输入密码，就是 deck 用户的密码。

#### 2.1 配置 ssh 免密登录（可选）

每次 ssh 登录都要输入密码，很麻烦，所以我们要配置一下免密登录，我们在要登录 Steam Deck 的那台电脑上创建一对 ssh-key，把公钥放到 Steam Deck 里的 `/home/deck/.ssh/authorized_keys` 文件内就行了。

#### 3. 传文件

配置好了 ssh，传文件就很方便了，比如我要传到 Steam Deck 的桌面
```
scp -r test deck@<Steam Deck IP>:<Steam Deck 路径>

# 例
scp -r test deck@192.168.0.10:/home/deck/Desktop
```

#### 3.1 传文件到 sd 卡

首先要知道 sd 卡的目录在哪里？我们可以通过 Steam Deck 的存储可以看到，tf 卡的目录可以在 Steam Deck 中查看






hello，大家好，这次视频主要是补充上期关于 ssh 文件互传的视频，虽然启动 ssh 已经能办到 PC端 和 Steam Deck 文件互传，而且 windows，linux，mac 都通用，但是使用起来不直观。

然后这次我来说说如何使用 FileZilla 实现 PC端 和 Steam Deck 文件互传。

首先，我们要在 Steam Deck 开启 sshd 的功能。

FileZilla 是一个 FTP 工具，是装在我们的 PC 端的。我们打开 filezilla 官网 `https://filezilla-project.org/`。

下载的是 Client 的版本，不是 Server。我们在 Steam Deck 上启用的 sshd 服务就已经充当了 Server 的角色。所以我们只需下载一个 Client 连上去就行了。