LOCAL_REGISTRY ?= localhost:5001

.PHONY: proto kind-up kind-down kind-reset dev debug

proto:
	buf generate
	go mod tidy

kind-up:
	./infra/kind/bootstrap.sh

kind-down:
	kind delete cluster

kind-reset: kind-down kind-up

dev:
	skaffold dev --port-forward --default-repo=${LOCAL_REGISTRY}

debug:
	skaffold debug --port-forward --default-repo=${LOCAL_REGISTRY}

