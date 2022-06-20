GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

default: terraform-provider-rocketchat

terraform-provider-rocketchat: fmtcheck
	go build

install: terraform-provider-rocketchat
	mkdir -p ~/.terraform.d/plugins/registry.terraform.io/magenta-aps/rocketchat/1.0.0/linux_amd64
	cp $+ ~/.terraform.d/plugins/registry.terraform.io/magenta-aps/rocketchat/1.0.0/linux_amd64

test_resource: install
	cd examples/resource/; rm .terraform.lock.hcl; rm -rf .terraform; terraform init; terraform apply

test_datasource: install
	cd examples/datasource/; rm .terraform.lock.hcl; rm -rf .terraform; terraform init; terraform apply

test: fmtcheck
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)

.PHONY: build test testacc vet fmt fmtcheck errcheck

