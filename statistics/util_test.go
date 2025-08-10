package statistics

import "testing"

func TestCalculateDeviation(t *testing.T) {
	count := 8

	values := make([]int, count)

	values[0] = 2
	values[1] = 4
	values[2] = 4
	values[3] = 4
	values[4] = 5
	values[5] = 5
	values[6] = 7
	values[7] = 9

	average := 5

	ans := CalculateDeviation(values, average, count)

	if ans != 2.0 {
		t.Errorf("Calculate Deviation wiki test = %f; want 2", ans)
	}
}
