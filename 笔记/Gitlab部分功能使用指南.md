# deploy key

gitlab 有一个 `deploy key` 的功能，可以给每个项目配置一个 deploy key，我查看了一下，功能和 ssh key 差不多，配置好了就可以 pull/push 这个项目的代码了，但是功能可以分得更细，比如只给 pull 权限，不给 push 权限，而且还是只对这个项目而言，对其他的项目依然没有 pull/push 权限。

但是对我感觉没什么用，这个和 CI 搭不上关系，只是对 SSH key 的更精细化操作而已。一般来说还是配置一个 SSH key 然后所有项目都有权限比较方便。