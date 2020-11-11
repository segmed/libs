all: names.go

names.go: out.names_first out.names_last
	# Generating first names first...
	@( \
		echo "// Copyright (c) 2019-2020 Segmed Inc."; \
		echo "// Author: Wojciech Adam Koszek <wkoszek@segmed.ai>"; \
	) > names.go
	@cat  out.names_first| awk 'BEGIN { \
		printf("package namedetect\n"); \
		printf("var AllFirstNames = []string{\n"); \
	} \
	// { \
		printf("\t\"\%s\",\n", $$1); \
	} \
	END { \
		printf("}\n"); \
	}' >> names.go
	# Generating lastnames next...
	@cat out.names_last | awk 'BEGIN { \
		printf("var AllLastNames = []string{\n"); \
	} \
	// { \
		printf("\t\"\%s\",\n", $$1); \
	} \
	END { \
		printf("}\n"); \
	}' >> names.go
	# All should be done.
	go build
	go test

out.names_first:
	cat names_first/*.txt               | cut -d , -f 1 | tr 'A-Z' 'a-z' | sort | uniq > out.names_first
	@wc -l out.names_first
	@echo "       ^^ -- this is how many unique first names we have"
out.names_last:
	cat names_last/Names_2010Census.csv | cut -d , -f 1 | tr 'A-Z' 'a-z' | sort | uniq > out.names_last
	@wc -l out.names_last
	@echo "       ^^ -- this is how many unique last names we have"

install-tools:
	go get -u github.com/markbates/pkger/cmd/pkger

clean:
	rm -rf out.*

.PHONY: out.names_first out.names_last
