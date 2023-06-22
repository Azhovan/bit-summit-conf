
.PHONY: deploy
deploy:
	kubectl apply -f ./hack/deployments
	kubectl get po --watch

.PHONY: reset
reset:
	kubectl delete -f ./hack/deployments

.PHONY: dashboard
dashboard:
	helm repo add "dapr https://dapr.github.io/helm-charts/"
	helm repo update
	helm install dapr-dashboard dapr/dapr-dashboard

.PHONY: rollout
rollout:
	 kubectl rollout restart deployment/supplier
	 kubectl rollout restart deployment/market
#	 kubectl rollout restart deployment/buyer1
#	 kubectl rollout restart deployment/buyer2

.PHONY: redis
redis:
	helm repo add bitnami https://charts.bitnami.com/bitnami
	helm repo update
	helm install redis bitnami/redis

.PHONY: tracing
tracing:
	#kubectl create deployment zipkin --image openzipkin/zipkin
	#kubectl expose deployment zipkin --type ClusterIP --port 9411
	kubectl port-forward svc/zipkin 9411:9411