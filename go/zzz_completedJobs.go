// 2022.02.23

package main

import "sort"

type Job []int // Job{id, priorityGroup (0 highest), processingTime (hours), deadline (hours)}

func completedJobs(jobs []Job) int {
	sort.Slice(jobs, func(i, j int) bool { // nlogn
		if jobs[i][1] == jobs[j][1] {
			if jobs[i][3]-jobs[i][2] == jobs[j][3]-jobs[j][2] {
				return jobs[i][3] < jobs[j][3]
			}
			return jobs[i][3]-jobs[i][2] < jobs[j][3]-jobs[j][2]
		}
		return jobs[i][1] < jobs[j][1]
	})

	ans := 0
	timeNow := 0
	for _, job := range jobs { // n
		if timeNow+job[2] <= job[3] {
			ans++
			timeNow += job[2]
		}
	}
	return ans
}
