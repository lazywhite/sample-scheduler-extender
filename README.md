```
mkdir -p /etc/kubernetes/scheduler

# 1.14版本只支持json格式policy
cp deploy/policy.json /etc/kubernetes/scheduler

cp deploy/config.yaml /etc/kubernetes/scheduler

cp deploy/kube-scheduler.yaml /etc/kubernetes/manifests/kube-scheduler.yaml

kubectl get -n kube-system pod #等待kube-scheduler pod重启
```
