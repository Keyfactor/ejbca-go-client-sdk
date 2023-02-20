GO_PACKAGE_NAME = ejbca
GIT_USER_ID = Keyfactor
GIT_REPO_NAME = ejbca-go-client-sdk
OPENAPI_YAML = ejbca-7.10.0.1-patched.yaml
OPENAPI_GENERATOR_VERSION = 6.3.0
OPENAPI_TEMPLATE_DIR = .openapi-generator/templates/go

generate: clean
	@echo "Initializing..."
	@wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/$(OPENAPI_GENERATOR_VERSION)/openapi-generator-cli-$(OPENAPI_GENERATOR_VERSION).jar -O openapi-generator-cli.jar
	@echo "Generating..."
	@if [ -d $(OPENAPI_TEMPLATE_DIR) ]; \
	then \
  		java -jar openapi-generator-cli.jar generate -i $(OPENAPI_YAML) -p packageName=$(GO_PACKAGE_NAME) $(OPENAPI_ARGS) -g go -p isGoSubmodule=false -p disallowAdditionalPropertiesIfNotPresent=false --git-user-id $(GIT_USER_ID) --git-repo-id $(GIT_REPO_NAME) -t $(OPENAPI_TEMPLATE_DIR); \
 	else \
		java -jar openapi-generator-cli.jar generate -i $(OPENAPI_YAML) -p packageName=$(GO_PACKAGE_NAME) $(OPENAPI_ARGS) -g go -p isGoSubmodule=false -p disallowAdditionalPropertiesIfNotPresent=false --git-user-id $(GIT_USER_ID) --git-repo-id $(GIT_REPO_NAME); \
 	fi
	@mkdir ./api/$(GO_PACKAGE_NAME) || (echo /api/$(GO_PACKAGE_NAME) exists)
	@mv ./*.go ./api/$(GO_PACKAGE_NAME) -f || (echo no files to move)
	@rm .travis.yml || (echo no .travis.yml to remove)
	@rm .openapi-generator-ignore || (echo no .openapi-generator-ignore to remove)
	@rm git_push.sh || (echo no git_push.sh to remove)
	@echo "Cleaning up..."
	@rm openapi-generator-cli.jar
	@echo "Done."

clean:
	@echo "Cleaning..."
	@rm -rf ./api/$(GO_PACKAGE_NAME) || (echo no ./api/$(GO_PACKAGE_NAME) to remove)
	@rm -rf ./go.mod || (echo no ./go.mod to remove)
	@rm -rf ./go.sum || (echo no ./go.sum to remove)
	@rm -rf ./docs || (echo no ./docs to remove)
	@rm -rf ./test || (echo no ./test to remove)

template:
	@echo "Initializing..."
	@wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/$(OPENAPI_GENERATOR_VERSION)/openapi-generator-cli-$(OPENAPI_GENERATOR_VERSION).jar -O openapi-generator-cli.jar
	@mkdir -p $(OPENAPI_TEMPLATE_DIR)
	@echo "Generating Templates..."
	@java -jar openapi-generator-cli.jar author template -g go -o $(OPENAPI_TEMPLATE_DIR)
	@echo "Cleaning up..."
	@rm openapi-generator-cli.jar
	@echo "Done."