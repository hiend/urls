OUT := ${CURDIR}/urls

all: build run clean

${OUT}: clean
	@echo Building...
	@go build -o $@

.PHONY: build
build: ${OUT}

.PHONY: run
run: ${OUT}
	@echo Running...
	@echo
	@${OUT}
	@echo

.PHONY: clean
clean:
	@echo Cleaning...
	@rm ${OUT} || true
