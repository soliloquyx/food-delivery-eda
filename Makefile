LOCAL_IMAGE_REGISTRY ?= localhost:5001

.PHONY: proto kind-up kind-down kind-reset dev debug migrate

proto:
	buf generate
	go mod tidy

kind-up:
	./infra/kind/bootstrap.sh

kind-down:
	kind delete cluster

kind-reset: kind-down kind-up

dev:
	skaffold dev --module services --port-forward --default-repo=${LOCAL_IMAGE_REGISTRY}

debug:
	skaffold debug --port-forward --default-repo=${LOCAL_IMAGE_REGISTRY}

migrate:
	kubectl delete job order-migrate --ignore-not-found
	skaffold run --module migrate --default-repo=${LOCAL_IMAGE_REGISTRY} --tail
	kubectl wait --for=condition=complete --timeout=120s job/order-migrate

