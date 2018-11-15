# 小程序提交如下数据到服务端
```
{
  appid:"xxx",
  secret:"xxx",
  // wx.login() 获取
  code:"021ZcaLF17kwz00lRpOF1LPtLF1ZcaLf",
  // wx.getUserInfo() 获取
  encryptedData:"eKu1e2T7toSICdr6Eqkj2t59xJCJ9vfiZs=",
  // wx.getUserInfo() 获取
  iv:"83e8zy+GcVej6/0zDby9DA==",
}
```

# 服务端解密

解密的过程可以参考[官方demo](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html)，解密所需的文件也在官方demo里面
```
const WXBizDataCrypt = require('./WXBizDataCrypt');
const wxObj = new WXBizDataCrypt(appid, session_key);
const data = wxObj.decryptData(encryptedData , iv);
console.log(data);
```

> 如果data里面有可能没有unionid，关于UnionID，微信官方介绍的是：同一用户，对同一个微信开放平台下的不同应用，unionid是相同的。也就是说，unionid是连接同一个账号下的不同的公众号的，所以没有公众号的话就不会有unionid。还有一种方法是把小程序绑定到开发平台，开发平台会给下面的小程序，公众号这些发放unionid。


# issues
- errcode 40029
> 原因是开发者工具中设置的appid和服务器解析的appid不一样。v