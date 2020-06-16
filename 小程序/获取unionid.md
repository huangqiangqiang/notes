# 小程序端获取用户信息（unionid）

### 客户端提交如下数据到服务端
```
{
  // wx.login() 获取
  code:"021ZcaLF17kwz00lRpOF1LPtLF1ZcaLf",
  // wx.getUserInfo() 获取
  encryptedData:"eKu1e2T7toSICdr6Eqkj2t59xJCJ9vfiZs=",
  // wx.getUserInfo() 获取
  iv:"83e8zy+GcVej6/0zDby9DA==",
}
```
`appid`和`secret`由服务端保管。

### 服务端解密

解密的过程可以参考[官方demo](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html)，解密所需的文件也在官方demo里面
```
const WXBizDataCrypt = require('./WXBizDataCrypt');
const wxObj = new WXBizDataCrypt(appid, session_key);
const data = wxObj.decryptData(encryptedData , iv);
console.log(data);
```

> 如果data里面有可能没有unionid，关于UnionID，微信官方介绍的是：同一用户，对同一个微信开放平台下的不同应用，unionid是相同的。也就是说，unionid是连接同一个账号下的不同的公众号的，所以没有公众号的话就不会有unionid。还有一种方法是把小程序绑定到开发平台，开发平台会给下面的小程序，公众号这些发放unionid。


### issues
- errcode 40029
> 原因是开发者工具中设置的appid和服务器解析的appid不一样。



# web端获取用户信息（unionid）

[参考链接](https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140842)

### 客户端提交code到服务端 

web客户端由微信授权后获取到`code`提交给服务端便可。

### 服务端

服务端要分几步才能获取到用户信息

1. 通过code换取网页授权access_token

> GET https://api.weixin.qq.com/sns/oauth2/access_token?appid=APPID&secret=SECRET&code=CODE&grant_type=authorization_code

只要注意appid和secret和获取code时的保持一致就可以了。

2. 获取用户信息

> GET https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN

# 本地调试

微信开发很多时候都有限制，比如小程序会限制域名，请求必须用https等。这给本地调试带来了难度。

我目前的解决方案是用本地的服务监听80端口，然后用ngrok穿透到外网去。获取到外网的接口地址后配到网关上（公司有统一的网关，支持https），就可以访问本地接口了。