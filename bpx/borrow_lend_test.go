package bpx

import (
	"testing"

	"bpx-api-client-go/types"

	"github.com/stretchr/testify/require"
)

func TestBorrowLend_GetBorrowLendPositions(t *testing.T) {
	c := NewClient(key, secret).BorrowLend()
	resp, err := c.GetBorrowLendPositions(types.GetBorrowLendPositionReq{})
	require.NoError(t, err)
	require.NotNil(t, resp)
}
