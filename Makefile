
# determine number of cores so we can create equivelant amount of DBs for tests
CORES=$(shell cat /proc/cpuinfo | grep processor | wc -l)

# gather options for tests
TESTARGS=$(TESTOPTIONS)

# gather options for coverage
COVERAGEARGS=$(COVERAGEOPTIONS)


update-deps:
		rm glide.lock vendor -rf && glide update

# build all binaries
build: build-script build-incomplete

build-script:
		go build -o build/script bin/script/main.go
build-incomplete:
		go build -o build/incomplete bin/incomplete/main.go

# concat all coverage reports together
coverage-concat:
	echo "mode: set" > coverage/full && \
    grep -h -v "^mode:" coverage/*.out >> coverage/full

# full coverage report
coverage: coverage-concat
	go tool cover -func=coverage/full $(COVERAGEARGS)

# full coverage report
coverage-html: coverage-concat
	go tool cover -html=coverage/full $(COVERAGEARGS)


