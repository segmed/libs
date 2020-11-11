// Copyright (c) 2019-2020 Segmed Inc.
// Author: Wojciech Adam Koszek <wkoszek@segmed.ai>
package namedetect

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestNames(t *testing.T) {
	fnum := len(AllFirstNames)
	lnum := len(AllLastNames)

	fmt.Printf("number of first names: %d\n", fnum)
	fmt.Printf("number of last names: %d\n", lnum)

	assert.Equal(t, fnum, 98400, "Number of first names doesn't match")
	assert.Equal(t, lnum, 162255, "Number of last names doesn't match")
}
