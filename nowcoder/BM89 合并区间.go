package nowcoder

import "sort"

type Interval struct {
	Start int
	End   int
}

func merge(intervals []*Interval) []*Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	merged := make([]*Interval, 0)
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		cur := intervals[i]

		if last.End < cur.Start {
			merged = append(merged, cur)
		} else {
			merged[len(merged)-1] = &Interval{
				Start: last.Start,
				End:   max(last.End, cur.End),
			}
		}
	}

	return merged
}
