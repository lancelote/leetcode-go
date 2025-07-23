package p933

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{queue: []int{}}
}

func (rc *RecentCounter) Ping(t int) int {
	for len(rc.queue) > 0 && rc.queue[0]+3000 < t {
		rc.queue = rc.queue[1:]
	}
	rc.queue = append(rc.queue, t)
	return len(rc.queue)
}
