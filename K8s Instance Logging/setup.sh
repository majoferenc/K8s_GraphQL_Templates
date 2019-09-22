kubectl create namespace logging
kubectl create -f elastic.yaml -n logging
kubectl create -f kibana.yaml -n logging
kubectl create -f fluentd-rbac.yaml
kubectl create -f fluentd-daemonset.yaml
