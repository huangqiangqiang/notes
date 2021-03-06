先讲一下为什么网上关于git讲解的图片中commit箭头总是感觉画反了，因为每个commit都有一个索引指向它的上游commit。

```
# 创建并切换分支
git checkout -b <branch>

# 拉取远端分支到本地并重命名
git checkout -b <new branch> origin/远程分支名

# 删除远程分支
git push origin -d <branch>

# 删除远程tag
git push origin --delete tag <tag>

# 把远端的最新代码合并到当前分支
git pull --rebase origin master
# 如果再做rebase之前push过代码，那么rebase后远端和现在的分支commit记录已经不一样了，这时是push不上去的，只能通过 --force 强制让远端的分支ocmmit记录和本地的一样
git push --force
```

在工作中遇到了一个需求，项目中的一些配置信息不提交到gitlab上面去，但是没有配置的话项目又跑不起来，而且有些配置文件是写在源文件内的，又不能写在gitignore里面去，所以采用以下办法：

```
# 拉去项目代码后新建并切换一个本地分支，此分支不提交到远端，只在本地
git checkout -b local

# 把配置项填入每个文件，让项目跑起来，然后做一个commit
git add .
git commit -m 'update config'

# 后面再做一些正常的无敏感信息的提交
git add .
git commit -m 'fix bug'

# 在本地分支中创建一个新分支并切换到新分支
git checkout -b <new branch> origin/master

# 在新的分支上选择要提交的commit
git cherry-pick commit1 commit2 ... commitN
或
git cherry-pick commit1..commitN

注意：不包含commit1，实际上选中的commit是commit2到commitN

# push筛选过的分支到远端
git push origin remote <branch name>

```

还有一种比较简单的方式，使用diff生成patch和apply命令

```
# 先把修改的配置项生成patch文件
git diff > conf.patch

# 提交前还原patch
git apply -R conf.patch

# 提交代码
git add .
git commit -m 'xxx'

# 在填写配置
git apply conf.patch

```

临时要切换分支，要先保存现在的修改记录

```
# 暂存修改
git stash save 'xxx'
# 还原修改
git stash pop
# 查看stash
git stash list
```
解决 git 合并遇到的冲突问题

我在远程上开了一个自己的分支 `dev_hqq` ，本地也有两个新提交并 push 到远端的 `dev_hqq` 分支上，现在在 master 分支上有另一个同事的提交，并且改了同一个文件。目前两个分支不能直接的合并。

合并的流程是：

1. 先用命令 `git pull --rebase origin master` 把 master 分支的代码合过来

2. 发现冲突，解决冲突

3. 使用 `git add .` 把解决冲突的文件加入到缓冲区中

4. 使用 `git --rebase continue` 继续合并下一个commit，直到成功