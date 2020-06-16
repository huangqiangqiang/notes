# 抓http包

- 在 Charles 的菜单栏上选择 “Proxy”->”Proxy Settings”，填入代理端口 8888，并且勾上 “Enable transparent HTTP proxying” 就完成了在 Charles 上的设置。

- 在 iPhone 的 “ 设置 “->” 无线局域网 “ 中，可以看到当前连接的 wifi 名，通过点击右边的详情键，可以看到当前连接上的 wifi 的详细信息，包括 IP 地址，子网掩码等信息。在其最底部有「HTTP 代理」一项，我们将其切换成手动，然后填上 Charles 运行所在的电脑的 IP，以及端口号 8888

- 设置好之后，我们打开 iPhone 上的任意需要网络通讯的程序，就可以看到 Charles 弹出 iPhone 请求连接的确认菜单（如下图所示），点击 “Allow” 即可完成设置。(如果没有弹出allow的框说明代理没有设置成功，试着把ss或其他代理关掉)。

# 抓https包

### 1. 安装证书

如果你需要截取分析 Https 协议相关的内容。那么需要安装 Charles 的 CA 证书。具体步骤如下。

首先我们需要在 Mac 电脑上安装证书。点击“Help” -> “SSL Proxying” -> “Install Charles Root Certificate”，然后输入系统的帐号密码，即可在 KeyChain 看到添加好的证书，提示证书不受信任没关系。

### 2. 添加所有网站https的代理

在`Proxy -> SSL Proxying Settings`中点击`add`，host填*，
port填443。

### 3. 截取移动设备中的 Https 通讯信息

- 点击 Charles 的顶部菜单，选择 “Help” -> “SSL Proxying” -> “Install Charles Root Certificate on a Mobile Device or Remote Browser”，然后就可以看到一个弹出框，
- 在iphone上输入弹框中的网址，比如`chls.pro/ssl`，点击安装证书。
- 在`Setting->General->About->Certificate Trust Settings`中enable刚才安装的证书。

# 将 Charles 设置成系统代理
启动 Charles 后，第一次 Charles 会请求你给它设置系统代理的权限。你可以输入登录密码授予 Charles 该权限。你也可以忽略该请求，然后在需要将 Charles 设置成系统代理时，选择菜单中的 “Proxy” -> “Mac OS X Proxy” 来将 Charles 设置成系统代理。
