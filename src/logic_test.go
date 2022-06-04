package logic_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	logic "gocircleci/app/src"
	"testing"
)

func TestWhichInput(t *testing.T) {
	res := logic.WhichInput(100)
	//res := 100
	fmt.Printf("Result: %d", res)
	if res != 100 {
		fmt.Errorf("got %q, expected %q", res, 100)
	}
	assert.Equal(t, 100, res)
}
