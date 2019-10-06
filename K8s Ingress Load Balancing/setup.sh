terraform init
terraform apply
kubectl apply -f managedcertificates-crd.yaml
kubectl apply -f managed-certificate-controller.yaml
kubectl create -f app-tls-cert.yaml
kubectl create -f app-ingress-https.yaml
kubectl describe ManagedCertificate