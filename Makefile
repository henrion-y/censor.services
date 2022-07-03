fmt:
	command -v gofumpt || (WORK=$(shell pwd) && cd /tmp && GO111MODULE=on go install mvdan.cc/gofumpt@latest && cd $(WORK))
	gofumpt -w -d .

lint:
	command -v golangci-lint || (WORK=$(shell pwd) && cd /tmp && GO111MODULE=on go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0 && cd $(WORK))
	golangci-lint run  -v
