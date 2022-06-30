package keeper

import (
	"athlete/x/athlete/types"
)

var _ types.QueryServer = Keeper{}
