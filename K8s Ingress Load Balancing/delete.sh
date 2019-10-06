kubectl delete -f managedcertificates-crd.yaml
kubectl delete -f managed-certificate-controller.yaml
kubectl delete -f app-tls-cert.yaml
kubectl delete -f app-ingress-https.yaml