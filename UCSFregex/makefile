all: phi_regex.go

phi_regex.go: out.regex_context_pattern out.regex_pattern out.regex_safe_pattern
	# Generating regex_context_pattern
	@( \
		echo "// Copyright (c) 2020-2021 Segmed Inc."; \
	) > phi_regex.go
	@cat  out.regex_context_pattern| awk -F'\n' 'BEGIN { \
		printf("package UCSFregex\n"); \
		printf("var AllRegexContextPattern = []string{\n"); \
	} \
	// { \
		printf("\t\`\%s\`,\n", $$1); \
	} \
	END { \
		printf("}\n"); \
	}' >> phi_regex.go
	# Generating regex_pattern next...
	@cat out.regex_pattern | awk -F'\n' 'BEGIN { \
		printf("var AllRegexPattern = []string{\n"); \
	} \
	// { \
		printf("\t\`\%s\`,\n", $$1); \
	} \
	END { \
		printf("}\n"); \
	}' >> phi_regex.go

	# Generating regex_safe_pattern next...
	@cat out.regex_safe_pattern | awk -F'\n' 'BEGIN { \
		printf("var AllRegexSafePattern = []string{\n"); \
	} \
	// { \
		printf("\t\`\%s\`,\n", $$1); \
	} \
	END { \
		printf("}\n"); \
	}' >> phi_regex.go

	# All should be done.
	go build
	go test

out.regex_context_pattern:
	for f in ./regex_context/*.txt; do (cat "$$f" ; echo) >> out.regex_context_pattern; done
	@wc -l out.regex_context_pattern
	@echo "       ^^ -- this is how many unique regex context pattern we have"

out.regex_pattern:
	find ./regex -type f \( -iname "*_transformed.txt" ! -iname "*_safe_transformed.txt" \) -print0 | while read -d $$'\0' f; do (cat "$$f" ; echo) >> out.regex_pattern; done
	@wc -l out.regex_pattern
	@echo "       ^^ -- this is how many unique regex pattern we have"

out.regex_safe_pattern:
	for f in ./regex/*/*_safe_transformed.txt; do (cat "$$f" ; echo) >> out.regex_safe_pattern; done
	@wc -l out.regex_safe_pattern
	@echo "       ^^ -- this is how many unique whitlist regex pattern we have"

install-tools:
	go get -u github.com/markbates/pkger/cmd/pkger
	go get -u github.com/dlclark/regexp2/...


clean:
	rm -rf out.*

.PHONY: out.regex_pattern out.regex_context_pattern out.regex_safe_pattern
