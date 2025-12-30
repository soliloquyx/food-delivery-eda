# Food Delivery Platform

An example project built to explore architectural concepts and technologies of interest.

## Setup

### Prerequisites

- [Go](https://go.dev)
- [Docker](https://www.docker.com)
- [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)
- [kind](https://kind.sigs.k8s.io)
- [Skaffold](https://skaffold.dev)
- [ko](https://ko.build)
- [Buf](https://buf.build)

### Local development

Create a local cluster:

```bash
make kind-up
```

Run the services:

```bash
make dev
```

The HTTP gateway is available at:

```text
http://localhost:8080
```
