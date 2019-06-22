package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegex(t *testing.T) {
	a := parseLineName("NS1")
	require.Equal(t, "NS", a)
}
