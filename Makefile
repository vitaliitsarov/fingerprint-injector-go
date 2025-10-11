.PHONY: help test run-basic run-custom run-stealth run-proxy run-multi run-random clean deps

help: ## Показать справку
	@echo "Доступные команды:"
	@echo "  make deps          - Установить зависимости"
	@echo "  make test          - Запустить тесты"
	@echo "  make run-basic     - Запустить базовый пример"
	@echo "  make run-custom    - Запустить пример с кастомным fingerprint"
	@echo "  make run-stealth   - Запустить stealth режим"
	@echo "  make run-proxy     - Запустить пример с прокси"
	@echo "  make run-multi     - Запустить множественные сессии"
	@echo "  make run-random    - Запустить со случайным fingerprint"
	@echo "  make clean         - Очистить временные файлы"

deps: ## Установить зависимости
	go mod download
	go mod tidy

test: ## Запустить тесты
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

run-basic: ## Запустить базовый пример
	cd examples/basic && go run main.go

run-custom: ## Запустить пример с кастомным fingerprint
	cd examples/custom && go run main.go

run-stealth: ## Запустить stealth режим
	cd examples/stealth && go run main.go

run-proxy: ## Запустить пример с прокси
	cd examples/with-proxy && go run main.go

run-multi: ## Запустить множественные сессии
	cd examples/multi-session && go run main.go

run-random: ## Запустить со случайным fingerprint
	cd examples/random && go run main.go

run-platform: ## Запустить выбор платформы (использование: make run-platform PLATFORM=linux)
	cd examples/platform-selector && go run main.go -platform=$(or $(PLATFORM),windows)

run-viewport: ## Демонстрация viewport для разных устройств
	cd examples/viewport-demo && go run main.go

run-generator: ## Умная генерация fingerprint с базой устройств
	cd examples/smart-generator && go run main.go

clean: ## Очистить временные файлы
	rm -rf chrome-data*/
	rm -rf examples/*/chrome-data*/
	rm -f coverage.out coverage.html
	go clean

lint: ## Запустить линтер
	golangci-lint run

format: ## Форматировать код
	go fmt ./...
	gofmt -s -w .

build: ## Собрать все примеры
	cd examples/basic && go build -o ../../bin/basic main.go
	cd examples/custom && go build -o ../../bin/custom main.go
	cd examples/stealth && go build -o ../../bin/stealth main.go
	cd examples/with-proxy && go build -o ../../bin/with-proxy main.go
	cd examples/multi-session && go build -o ../../bin/multi-session main.go
	cd examples/random && go build -o ../../bin/random main.go

