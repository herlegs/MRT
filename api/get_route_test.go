package api

import (
	"testing"
	"time"

	"github.com/herlegs/MRT/core"

	"github.com/stretchr/testify/require"
)

func TestParseTime(t *testing.T) {
	tt, err := parseTime("2019-06-21T07:30")
	require.NoError(t, err)
	require.EqualValues(t, time.Date(2019, 06, 21, 7, 30, 0, 0, time.UTC), tt)
}

func TestGetTrafficMode(t *testing.T) {
	require.Equal(t, core.Night, getTrafficMode("2019-06-21T00:30"))

	require.Equal(t, core.Peak, getTrafficMode("2019-06-21T06:30"))

	require.Equal(t, core.Normal, getTrafficMode("2019-06-21T11:30"))
}
