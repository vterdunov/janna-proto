PROTOTOOL_IMAGE = uber/prototool:1.8.0

.PHONY: help
help: ## Display this help message
	@echo "Please use \`make <target>\` where <target> is one of:"
	@cat $(MAKEFILE_LIST) | grep -e "^[-a-zA-Z_\.]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: generate
generate: ## Generate .go files from .proto file
	docker run -t --rm -v $(CURDIR):/work $(PROTOTOOL_IMAGE) prototool lint --error-format="filename:line:column:id:message" proto
	docker run -t --rm -v $(CURDIR):/work $(PROTOTOOL_IMAGE) prototool generate proto
	docker run -t --rm -v $(CURDIR):/work -w /work golang:1.14.2-alpine3.11 go generate box/box.go
	echo Change generated file owner:
	sudo chown -R $(USER):$(USER) gen/
