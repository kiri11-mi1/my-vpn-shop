package subscription

func GetActualPrice(keysCount int, totalVPNPrice float64) float64 {
	return totalVPNPrice / float64(keysCount)
}
