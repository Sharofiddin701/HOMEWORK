func plusMinus(arr []int32) {
    positiveCount := 0
    negativeCount := 0
    zeroCount := 0

  for _, num := range arr {
   if num > 0 {
      positiveCount++
    } else if num < 0 {
      negativeCount++
    } else {
      zeroCount++
    }
  }

  total := float64(len(arr))

  positiveRatio := float64(positiveCount) / total
  negativeRatio := float64(negativeCount) / total
  zeroRatio := float64(zeroCount) / total

  fmt.Printf("%.6fn", positiveRatio)
  fmt.Printf("%.6fn", negativeRatio)
  fmt.Printf("%.6fn", zeroRatio)
}