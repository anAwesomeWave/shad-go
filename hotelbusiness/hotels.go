//go:build !solution

package hotelbusiness

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func ComputeLoad(guests []Guest) []Load {

	dates := make(map[int]int)
	maxDate := 0
	for _, v := range guests {
		maxDate = max(maxDate, v.CheckInDate, v.CheckOutDate)
		dates[v.CheckInDate]++
		dates[v.CheckOutDate]--
	}
	curAns := 0
	var ans []Load
	for i := 0; i <= maxDate; i++ {
		if dates[i] == 0 {
			continue
		}
		curAns += dates[i]
		ans = append(ans, Load{i, curAns})
	}

	return ans
}
