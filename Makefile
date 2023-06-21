
.PHONY: deploy
deploy:
	kubectl apply -f ./hack/deployments
	kubectl get po --watch

.PHONY: reset
reset:
	kubectl delete -f ./hack/deployments
