package calculatepaintneeded

import "fmt"

func Calculatepaintneeded(height float64, width float64) (float64, error) {
	var paintneededpermeter float64
	var area float64
	paintneededpermeter = 10.00
	if height < 0 {
		return 0, fmt.Errorf("%0.2f", height)
	}
	if width < 0 {
		return 0, fmt.Errorf("%0.2f", width)
	}

	area = height * width
	return area / paintneededpermeter, nil

}
