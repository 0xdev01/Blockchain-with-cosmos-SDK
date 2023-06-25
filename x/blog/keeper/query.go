package keeper

import (
	"union/x/blog/types"
)

var _ types.QueryServer = Keeper{}
