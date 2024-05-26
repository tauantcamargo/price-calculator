package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

func StringsToFloat(values []string) ([]float64, error) {
  floatValues := make([]float64, len(values)) 

  for valueIndex, value := range values {
    floatValue, err := strconv.ParseFloat(value, 64)

    if err != nil {
      fmt.Println("Error parsing value")
      fmt.Println(err)
      return nil, errors.New("Error parsing value")
    }

    floatValues[valueIndex] = floatValue
  }

  return floatValues, nil
}