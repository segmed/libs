// Copyright (c) 2020-2021 Segmed Inc.
package UCSFName

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestNames(t *testing.T) {
	fnum := len(AllFirstNames)
	lnum := len(AllLastNames)
	wnum := len(AllWhitelistNames)

	fmt.Printf("number of first names: %d\n", fnum)
	fmt.Printf("number of last names: %d\n", lnum)
	fmt.Printf("number of whitelist names: %d\n", wnum)

	assert.Equal(t, fnum, 97016, "Number of first names doesn't match")
	assert.Equal(t, lnum, 161472, "Number of last names doesn't match")
	assert.Equal(t, wnum, 195254, "Number of whitelist names doesn't match")
}
