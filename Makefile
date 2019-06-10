PROTOTOOL_IMAGE = uber/prototool:1.6.0

.PHONY: help
help: ## Display this help message
	@echo "Please use \`make <target>\` where <target> is one of:"
	@cat $(MAKEFILE_LIST) | grep -e "^[-a-zA-Z_\.]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: proto
proto: ## Generate .go files from .proto file
	docker run -it --rm -v $(CURDIR):/work $(PROTOTOOL_IMAGE) prototool lint --error-format="filename:line:column:id:message" proto
	docker run -it --rm -v $(CURDIR):/work $(PROTOTOOL_IMAGE) prototool generate proto
	echo Change generated file owner:
	sudo chown -R $(USER):$(USER) gen/
