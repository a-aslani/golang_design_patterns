PACKAGE_NAME = "github.com/a-aslani/golang_design_patterns"

install-tools:
	@echo installing tools && \
	@go install go.uber.org/mock/mockgen@latest \
	@echo done

.PHONY: mock
mock:
	mockgen -package mockfacade -destination structural/facade/mock/weather.go ${PACKAGE_NAME}/structural/facade CurrentWeatherDataRetriever
