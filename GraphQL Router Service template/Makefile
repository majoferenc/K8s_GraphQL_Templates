deploy_all:
	sh mongo-db/deploy_k8s.sh
	helm install --name prisma prisma/helm-prisma
	sh deploy_k8s.sh

remove_all:
	sh mongo-db/delete_k8s.sh
	sh prisma/delete_k8s.sh
	sh delete_k8s.sh