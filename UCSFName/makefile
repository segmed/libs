all: names.go

names.go: out.names_first out.names_last out.names_whitelist
	# Generating first names first...
	@( \
		echo "// Copyright (c) 2020-2021 Segmed Inc."; \
	) > names.go
	@cat  out.names_first| awk 'BEGIN { \
		printf("package UCSFName\n"); \
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
	# Generating whitelist names next...
	@cat out.names_whitelist | awk 'BEGIN { \
		printf("var AllWhitelistNames = []string{\n"); \
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
	jq 'keys' names/firstnames_minus_fps.json | tr -d '"' | cut -d , -f 1 | sort | uniq > out.names_first
	@wc -l out.names_first
	@echo "       ^^ -- this is how many unique first names we have"
out.names_last:
	jq 'keys' names/lastnames_minus_fps.json | tr -d '"' | cut -d , -f 1 | sort | uniq > out.names_last
	@wc -l out.names_last
	@echo "       ^^ -- this is how many unique last names we have"
out.names_whitelist:
	jq 'keys' names/names_whitelist.json | tr -d '"' | cut -d , -f 1 | sort | uniq > out.names_whitelist
	@wc -l out.names_whitelist
	@echo "       ^^ -- this is how many unique whitelist names we have, we should skip these during deid"

install-tools:
	go get -u github.com/markbates/pkger/cmd/pkger

clean:
	rm -rf out.*

.PHONY: out.names_first out.names_last out.names_whitelist
