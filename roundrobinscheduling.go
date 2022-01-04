package main

import "fmt"

func schedule(weeks int, teams []int) [][][]int {

	//Check if team count is odd, if so we add a 'bye'
	if len(teams)%2 != 0 {
		teams = append(teams, len(teams)+1)
	}

	halfTeams := len(teams) / 2

	mTeams := teams[1:]
	count := len(mTeams)

	schedule := make([][][]int, weeks)

	for i := 0; i < weeks; i++ {
		teamIndex := i % count

		schedule[i] = make([][]int, halfTeams)

		schedule[i][0] = make([]int, 2)

		schedule[i][0][0] = teams[0]
		schedule[i][0][1] = mTeams[teamIndex]

		for j := 1; j < halfTeams; j++ {
			teamA := (i + j) % count
			teamB := (i + count - j) % count

			schedule[i][j] = make([]int, 2)

			schedule[i][j][0] = mTeams[teamA]
			schedule[i][j][1] = mTeams[teamB]
		}
	}

	return schedule
}

func main() {
	weeks := 13
	teams := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	fmt.Println(schedule(weeks, teams))
}
