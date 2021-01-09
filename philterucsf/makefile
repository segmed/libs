all: names.go

names.go: out.names_first out.names_last
	# Generating first names first...
	@( \
		echo "// Copyright (c) 2019-2020 Segmed Inc."; \
		echo "// Author: Wojciech Adam Koszek <wkoszek@segmed.ai>"; \
	) > names.go
	jq -r 'keys[]' firstnames_minus_fps.json | awk 'BEGIN { \
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
	jq -r 'keys[]' lastnames_minus_fps.json | awk 'BEGIN { \
		printf("var AllLastNames = []string{\n"); \
	} \
	// { \
		printf("\t\"\%s\",\n", $$1); \
	} \
	END { \
		printf("}\n"); \
	}' >> names.go
	# Generating lastnames next...
	jq -r 'keys[]' whitelist_plus_fps_091718_nonames.json | awk 'BEGIN { \
		printf("var PermitListPlusFps = []string{\n"); \
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


install-tools:
	go get -u github.com/markbates/pkger/cmd/pkger

clean:
	rm -rf out.*

.PHONY: out.names_first out.names_last
