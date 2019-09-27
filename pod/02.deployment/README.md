# deployment的常用命令


## 查看部署状态
kubectl rollout status deployment <pod-name>  -n <namespace>
kubectl describe deployment <pod-name>  - <namespace>


## 升级
kubectl set image deployment review-demo review-demo=library/review-demo:0.0.1 --namespace=scm

kubectl edit deployment/review-demo --namespace=scm
**编辑.spec.template.spec.containers[0].image的值**

## 终止升级
kubectl rollout pause deployment/review-demo --namespace=scm

## 继续升级
kubectl rollout resume deployment/review-demo --namespace=scm

## 回滚
kubectl rollout undo deployment/review-demo --namespace=scm

**查看deployments版本**
kubectl rollout history deployments --namespace=scm

**回滚到指定版本**

kubectl rollout undo deployment/review-demo --to-revision=2 --namespace=scm

## 升级历史
kubectl describe deployment/review-demo --namespace=scm