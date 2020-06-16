test:
	@echo
	@echo '>> Running the tests'
	go test -v --cover ./style/

bench:
	@echo
	@echo '>> Running the benchmarks'
	go test --bench=. ./style/

all: test bench