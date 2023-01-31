package subscription

func GetActualPrice(keysCount, totalVPNPrice int) float64 {
	return float64(totalVPNPrice) / float64(keysCount)
}
