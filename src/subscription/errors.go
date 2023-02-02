package subscription

import "errors"

var ErrZeroKeysInServer = errors.New("not keys count in server")
var ErrNegativeTotalPrice = errors.New("total price must be positive")
