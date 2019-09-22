kubectl apply -f namespace.yaml
kubectl create clusterrolebinding cluster-admin-binding --clusterrole=cluster-admin --user=Marian.Ferenc2@ibm.com
kubectl apply -f prometheus-rbac.yaml
kubectl apply -f prometheus-config.yaml
kubectl apply -f prometheus-deploy.yaml
kubectl apply -f prometheus-svc.yaml
kubectl apply -f grafana.yaml
kubectl apply -f node-exporter.yaml
kubectl apply -f state-metrics-deploy.yaml
kubectl apply -f state-metrics-rbac.yaml

kubectl expose deployment grafana --type=LoadBalancer --namespace=monitoring