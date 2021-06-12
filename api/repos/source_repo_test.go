package repos

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckSourceExist(t *testing.T) {
	resval := Hi()
	fmt.Println(resval)
	assert.Equal(t, resval, true)
}

func TestCheckSourceExist2(t *testing.T) {
	resval := GetSources2()
	assert.Equal(t, resval, true)
}
