MKFILE_PATH=$(abspath $(lastword $(MAKEFILE_LIST)))
BUILD_PATH=$(patsubst %/,%,$(dir $(MKFILE_PATH)))/build/working_dir
VERSION=0.3.0
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

ifeq ($(GOOS),linux)
MD5=md5sum
SED=sed -i
else
MD5=md5 -q
SED=sed
endif

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

build-plugin: fmt vet ## Build plugins locally.
	go mod tidy
	rm -f .devstream/sum.md5
	mkdir -p .devstream

	go build -buildmode=plugin -trimpath -gcflags="all=-N -l" -o .devstream/githubactions-golang-${GOOS}-${GOARCH}_${VERSION}.so ./cmd/githubactions/golang
	${MD5} .devstream/githubactions-golang-${GOOS}-${GOARCH}_${VERSION}.so > .devstream/githubactions-golang-${GOOS}-${GOARCH}_${VERSION}.md5
	cat .devstream/githubactions-golang-${GOOS}-${GOARCH}_${VERSION}.md5 >> .devstream/sum.md5

	go build -buildmode=plugin -trimpath -gcflags="all=-N -l" -o .devstream/githubactions-python-${GOOS}-${GOARCH}_${VERSION}.so ./cmd/githubactions/python
	${MD5} .devstream/githubactions-python-${GOOS}-${GOARCH}_${VERSION}.so > .devstream/githubactions-python-${GOOS}-${GOARCH}_${VERSION}.md5 >> .devstream/sum.md5
	cat .devstream/githubactions-python-${GOOS}-${GOARCH}_${VERSION}.md5 >> .devstream/sum.md5

	go build -buildmode=plugin -trimpath -gcflags="all=-N -l" -o .devstream/githubactions-nodejs-${GOOS}-${GOARCH}_${VERSION}.so ./cmd/githubactions/nodejs
	${MD5} .devstream/githubactions-nodejs-${GOOS}-${GOARCH}_${VERSION}.so > .devstream/githubactions-nodejs-${GOOS}-${GOARCH}_${VERSION}.md5
	cat .devstream/githubactions-nodejs-${GOOS}-${GOARCH}_${VERSION}.md5 >> .devstream/sum.md5

	go build -buildmode=plugin -trimpath -gcflags="all=-N -l" -o .devstream/trello-${GOOS}-${GOARCH}_${VERSION}.so ./cmd/trello/
	${MD5} .devstream/trello-${GOOS}-${GOARCH}_${VERSION}.so > .devstream/trello-${GOOS}-${GOARCH}_${VERSION}.md5
	cat .devstream/trello-${GOOS}-${GOARCH}_${VERSION}.md5 >> .devstream/sum.md5

	go build -buildmode=plugin -trimpath -gcflags="all=-N -l" -o .devstream/trello-github-integ-${GOOS}-${GOARCH}_${VERSION}.so ./cmd/trellogithub/
	${MD5} .devstream/trello-github-integ-${GOOS}-${GOARCH}_${VERSION}.so > .devstream/trello-github-integ-${GOOS}-${GOARCH}_${VERSION}.md5
	cat .devstream/trello-github-integ-${GOOS}-${GOARCH}_${VERSION}.md5 >> .devstream/sum.md5

	go build -buildmode=plugin -trimpath -gcflags="all=-N -l" -o .devstream/argocd-${GOOS}-${GOARCH}_${VERSION}.so ./cmd/argocd/
	${MD5} .devstream/argocd-${GOOS}-${GOARCH}_${VERSION}.so > .devstream/argocd-${GOOS}-${GOARCH}_${VERSION}.md5
	cat .devstream/argocd-${GOOS}-${GOARCH}_${VERSION}.md5 >> .devstream/sum.md5

	go build -buildmode=plugin -trimpath -gcflags="all=-N -l" -o .devstream/argocdapp-${GOOS}-${GOARCH}_${VERSION}.so ./cmd/argocdapp/
	${MD5} .devstream/argocdapp-${GOOS}-${GOARCH}_${VERSION}.so > .devstream/argocdapp-${GOOS}-${GOARCH}_${VERSION}.md5
	cat .devstream/argocdapp-${GOOS}-${GOARCH}_${VERSION}.md5 >> .devstream/sum.md5

	go build -buildmode=plugin -trimpath -gcflags="all=-N -l" -o .devstream/jenkins-${GOOS}-${GOARCH}_${VERSION}.so ./cmd/jenkins/
	${MD5} .devstream/jenkins-${GOOS}-${GOARCH}_${VERSION}.so > .devstream/jenkins-${GOOS}-${GOARCH}_${VERSION}.md5
	cat .devstream/jenkins-${GOOS}-${GOARCH}_${VERSION}.md5 >> .devstream/sum.md5

	go build -buildmode=plugin -trimpath -gcflags="all=-N -l" -o .devstream/kube-prometheus-${GOOS}-${GOARCH}_${VERSION}.so ./cmd/kubeprometheus/
	${MD5} .devstream/kube-prometheus-${GOOS}-${GOARCH}_${VERSION}.so > .devstream/kube-prometheus-${GOOS}-${GOARCH}_${VERSION}.md5
	cat .devstream/kube-prometheus-${GOOS}-${GOARCH}_${VERSION}.md5 >> .devstream/sum.md5

	go build -buildmode=plugin -trimpath -gcflags="all=-N -l" -o .devstream/github-repo-scaffolding-golang-${GOOS}-${GOARCH}_${VERSION}.so ./cmd/reposcaffolding/github/golang/
	${MD5} .devstream/github-repo-scaffolding-golang-${GOOS}-${GOARCH}_${VERSION}.so > .devstream/github-repo-scaffolding-golang-${GOOS}-${GOARCH}_${VERSION}.md5
	cat .devstream/github-repo-scaffolding-golang-${GOOS}-${GOARCH}_${VERSION}.md5 >> .devstream/sum.md5

	go build -buildmode=plugin -trimpath -gcflags="all=-N -l" -o .devstream/devlake-${GOOS}-${GOARCH}_${VERSION}.so ./cmd/devlake/
	${MD5} .devstream/devlake-${GOOS}-${GOARCH}_${VERSION}.so > .devstream/devlake-${GOOS}-${GOARCH}_${VERSION}.md5
	cat .devstream/devlake-${GOOS}-${GOARCH}_${VERSION}.md5 >> .devstream/sum.md5

	go build -buildmode=plugin -trimpath -gcflags="all=-N -l" -o .devstream/gitlabci-golang-${GOOS}-${GOARCH}_${VERSION}.so ./cmd/gitlabci/golang
	${MD5} .devstream/gitlabci-golang-${GOOS}-${GOARCH}_${VERSION}.so > .devstream/gitlabci-golang-${GOOS}-${GOARCH}_${VERSION}.md5
	cat .devstream/gitlabci-golang-${GOOS}-${GOARCH}_${VERSION}.md5 >> .devstream/sum.md5

	go build -buildmode=plugin -trimpath -gcflags="all=-N -l" -o .devstream/jira-github-integ-${GOOS}-${GOARCH}_${VERSION}.so ./cmd/jiragithub/
	${MD5} .devstream/jira-github-integ-${GOOS}-${GOARCH}_${VERSION}.so > .devstream/jira-github-integ-${GOOS}-${GOARCH}_${VERSION}.md5
	cat .devstream/jira-github-integ-${GOOS}-${GOARCH}_${VERSION}.md5 >> .devstream/sum.md5

	go build -buildmode=plugin -trimpath -gcflags="all=-N -l" -o .devstream/openldap-${GOOS}-${GOARCH}_${VERSION}.so ./cmd/openldap/
	${MD5} .devstream/openldap-${GOOS}-${GOARCH}_${VERSION}.so > .devstream/openldap-${GOOS}-${GOARCH}_${VERSION}.md5
	cat .devstream/openldap-${GOOS}-${GOARCH}_${VERSION}.md5 >> .devstream/sum.md5
	data=shell awk BEGIN{RS=EOF}'{gsub(/\n/,":");print}' .devstream/sum.md5
	${SED} 's/  /_/g' .devstream/sum.md5

build: fmt vet build-plugin ## Build dtm & plugins locally.
	go mod tidy
	go build -trimpath -gcflags="all=-N -l" -ldflags "-X github.com/merico-dev/stream/cmd/devstream/version.Version=${VERSION} -X github.com/merico-dev/stream/cmd/devstream/version.MD5String=$(shell awk BEGIN{RS=EOF}'{gsub(/\n/,":");print}' .devstream/sum.md5)" -o dtm-${GOOS}-${GOARCH} ./cmd/devstream/

build-core: fmt vet build-plugin ## Build dtm & plugins locally.
	go mod tidy
	go build -trimpath -gcflags="all=-N -l" -ldflags "-X github.com/merico-dev/stream/cmd/devstream/version.Version=${VERSION} -X github.com/merico-dev/stream/cmd/devstream/version.MD5String=$(shell awk BEGIN{RS=EOF}'{gsub(/\n/,":");print}' .devstream/sum.md5)" -o dtm-${GOOS}-${GOARCH} ./cmd/devstream/

clean: ## Remove local plugins and locally built artifacts.
	rm -rf .devstream
	rm -f dtm*
	rm -rf build/working_dir

build-linux-amd64: ## Cross-platform build for "linux/amd64".
	echo "Building in ${BUILD_PATH}"
	mkdir -p .devstream
	rm -rf ${BUILD_PATH} && mkdir ${BUILD_PATH}
	docker buildx build --platform linux/amd64 --load -t mericodev/stream-builder:v${VERSION} --build-arg VERSION=${VERSION} -f build/package/Dockerfile .
	cp -r go.mod go.sum cmd internal pkg build/package/build_linux_amd64.sh ${BUILD_PATH}/
	chmod +x ${BUILD_PATH}/build_linux_amd64.sh
	docker run --rm --platform linux/amd64 -v ${BUILD_PATH}:/devstream mericodev/stream-builder:v${VERSION}
	mv ${BUILD_PATH}/output/*.so .devstream/
	mv ${BUILD_PATH}/output/dtm* .
	rm -rf ${BUILD_PATH}

fmt: ## Run 'go fmt' & goimports against code.
	go install golang.org/x/tools/cmd/goimports@latest
	goimports -local="github.com/merico-dev/stream" -d -w cmd
	goimports -local="github.com/merico-dev/stream" -d -w pkg
	goimports -local="github.com/merico-dev/stream" -d -w internal
	goimports -local="github.com/merico-dev/stream" -d -w test
	go fmt ./...

vet: ## Run go vet against code.
	go vet ./...

e2e: build ## Run e2e tests.
	./dtm-${GOOS}-${GOARCH} apply -f config.yaml
	./dtm-${GOOS}-${GOARCH} verify -f config.yaml
	./dtm-${GOOS}-${GOARCH} delete -f config.yaml

e2e-up: ## Start kind cluster for e2e tests.
	sh hack/e2e/e2e-up.sh

e2e-down: ## Stop kind cluster for e2e tests.
	sh hack/e2e/e2e-down.sh
