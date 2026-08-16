type TimeMap struct {
	Storage map[string][]Timestamp
}

type Timestamp struct {
	time int
	val  string
}

func Constructor() TimeMap {
	return TimeMap{
		make(map[string][]Timestamp),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if this.Storage[key] == nil {
		this.Storage[key] = make([]Timestamp, 0)
	}
	this.Storage[key] = append(this.Storage[key], Timestamp{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	row := this.Storage[key]
	if row == nil || timestamp < row[0].time {
		return ""
	}
	ans := row[0].val
	for l, r := 0, len(row)-1; l <= r; {
		m := (l + r) / 2
		if row[m].time == timestamp {
			return row[m].val
		} else if row[m].time < timestamp {
			l = m + 1
		} else {
			r = m - 1
		}
		ans = row[r].val
	}
	return ans
}