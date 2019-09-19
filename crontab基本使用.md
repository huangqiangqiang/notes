首先准备好要执行的脚本。

写入到 crontab 中，输入 `crontab -e`

格式为：
```
分      时      日      月      星期几      command
0-59    0-23   1-31    1-12   0-7
```

- * 代表所有可能的值
- / 代表间隔频率，比如 0-23/2 表示每两小时执行一次， */10 表示每十分钟执行一次

例：

```
*/1 * * * * sh ~/Desktop/mongo_bak.sh 2>~/Desktop/error.txt
```

以上语句的意思为 `每分钟执行一次 ~/Desktop/mongo_bak.sh 脚本，并且把错误输出到 ~/Desktop/error.txt 文件中`

重启 crontab

```
# mac
sudo /usr/sbin/cron restart

# linux

```