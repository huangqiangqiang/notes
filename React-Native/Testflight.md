# Testflight

1. RN打成Testflight包的话需要生产环境的证书和描述文件，必须是AppStore的生产环境，因为testflight测试通过的话可以直接拿包去发布了，所以和线上环境的包是一样的。

2. 证书和描述文件弄好后确保Scheme为release，点Xcode的`Archive`按钮，包打好后点击`Distribute App -> iOS App Store`，接下去就是一步步根据自己的设置去设置了。

3. 最后点击upload上传后，`App Connect`的`testflight`栏就会出现刚刚提交的包了，不过可能状态还是为`正在处理`，过会就好了。