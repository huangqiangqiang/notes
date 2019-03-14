# iOS离线测试包

1. 先在项目的ios文件夹内创建`bundle`文件夹，再执行以下命令打包文件到`bundle`文件夹内
```
react-native bundle --entry-file App.js --platform ios --bundle-output ./iOS/bundle/index.jsbundle --assets-dest ./ios/bundle
```

2. 把文件夹拖入到xcode中，文件夹引入的方式，引入后是个蓝色的文件夹图标。

3. `AppDelegate`中`jsCodeLocation`改为以下代码，意为加载本地文件
```
jsCodeLocation = [[NSBundle mainBundle] URLForResource:@"bundle/index" withExtension:@"jsbundle"];
```