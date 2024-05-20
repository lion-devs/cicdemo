# GitOps Demo

(WIP) ![Build Status]()
(WIP) ![Code Coverage]()
(WIP) ![App Status]()

A Golang project for GitOps pipeline demo

## Go

- Install
    - `brew install go`
    - `go version`
- Build
    - `go build -o build/gitops-demo cmd/main.go`
- Run
    - `go run cmd/main.go`


## Docker

- Edit tags
    - edit file "Makefile" => `VERSION`
- build and push image using makefile
    - build: `make docker-build`
    - push: `make docker-push`
- run docker container
    - run: `make docker-run`

## Deploy to Kubernetes

- Edit tags
    - edit file "deploy/base/kustomization.yaml" => `images.newTag`
- Kustomize
    - Build(preview): `kustomize build ./deploy/base`
    - Apply: `kubectl apply -k ./deploy/base`

## Others

- Network Test Container
    - `kubectl run net-test --image=justinhung0407/net-tools:v2.0.1 --restart=Never -n default sleep 3600`

- Update swagger
  - `go install github.com/swaggo/swag/cmd/swag@latest`
  - `swag init -g cmd/main.go --parseInternal -o docs/swagger`

- Run in background
    - `cd /opt/go-admin nohup go run main.go >> /tmp/go-http.log 2>&1 &`

- Update go dependencies
    - `go get -u ./cmd`
    - `go mod tidy`

- Delete single pod's container
    - `kubectl exec -it go-cicd-d64644988-bwtql -c opa-istio -n go-cicd -- /bin/sh -c "kill 1"`