.DEFAULT_GOAL := help
.PHONY: submodule generate help

REPO_ROOT:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

BUF_VERSION="1.17.0"
BIN="bin"

bin: ## install deps (library & development)
	mkdir -p bin/
	@ curl --fail --silent --show-error --location \
		"https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$(shell uname -s)-$(shell uname -m)" \
		-o "$(BIN)/buf" && \
		chmod +x "$(BIN)/buf"

submodule:  ## retrieve public protospecs
	git submodule update --init --recursive --remote

generate: submodule bin ## generate code from specifications
	cp buf.yaml api/buf.yaml
	cp buf.gen.yaml api/buf.gen.yaml

	cd api; ../bin/buf mod update
	cd api; ../bin/buf generate

help: ## Show help message
	@IFS=$$'\n' ; \
	help_lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##/:/'`); \
	printf "%s\n\n" "Usage: make [task]"; \
	printf "%-20s %s\n" "task" "help" ; \
	printf "%-20s %s\n" "------" "----" ; \
	for help_line in $${help_lines[@]}; do \
		IFS=$$':' ; \
		help_split=($$help_line) ; \
		help_command=`echo $${help_split[0]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		help_info=`echo $${help_split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		printf '\033[36m'; \
		printf "%-20s %s" $$help_command ; \
		printf '\033[0m'; \
		printf "%s\n" $$help_info; \
	done
