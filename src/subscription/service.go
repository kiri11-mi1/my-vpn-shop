package subscription

func GetActualPrice(keysCount int, totalVPNPrice float64) (float64, error) {
	if keysCount == 0 {
		return -1, ZeroKeysInServer
	}
	return totalVPNPrice / float64(keysCount), nil
}
