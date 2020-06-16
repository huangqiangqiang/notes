# 生成一个测试Pod，退出自动删除

```
kubectl run ${RANDOM} --rm -it --image=cirros:0.4 -- /bin/sh
```

# 通过 ip 验证 pod 接口
使用 `kubectl get po` 或 `kubectl get svc` 获取 Pod 和 Service 的 ip，然后生成一个测试 pod，因为只有在 pod 里面才能访问这个 ip，在 node 上是不行的。
```
curl -k http://<pod ip 或 service ip>
```

# 通过 kube dns 验证 pod 接口

```
curl -k https://<service name>.<namespace>/path
```

这种方式也可以作为容器内服务发现的一种方法