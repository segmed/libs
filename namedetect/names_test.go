// Copyright (c) 2019-2021 Segmed Inc.
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
	rnum := len(AllFrenchNames)

	fmt.Printf("number of first names: %d\n", fnum)
	fmt.Printf("number of last names: %d\n", lnum)
	fmt.Printf("number of french names: %d\n", rnum)

	assert.Equal(t, fnum, 98400, "Number of first names doesn't match")
	assert.Equal(t, lnum, 162255, "Number of last names doesn't match")
	assert.Equal(t, lnum, 34262, "Number of french names doesn't match")
}
