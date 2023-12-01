package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CricketScoreboard struct {
	Team1       string
	Team2       string
	TotalOvers  int
	RunsTeam1   int
	RunsTeam2   int
	CurrentOver int
	CurrentBall int
}

func (cs *CricketScoreboard) updateScore(runs int) {
	if cs.CurrentOver < cs.TotalOvers {
		if cs.CurrentBall < 6 {
			if cs.CurrentOver%2 == 0 {
				cs.RunsTeam1 += runs
			} else {
				cs.RunsTeam2 += runs
			}

			cs.CurrentBall++

			if cs.CurrentBall == 6 {
				fmt.Println("Over completed. Switching to the next over.")
				cs.CurrentBall = 0
				cs.CurrentOver++
			}

		} else {
			fmt.Println("Over completed. Switching to the next over.")
			cs.CurrentBall = 0
			cs.CurrentOver++
		}

	} else {
		fmt.Println("Inning completed. Game Over.")
		os.Exit(0)
	}
}

func (cs *CricketScoreboard) displayScore() {
	fmt.Printf("\n%s: %d | %s: %d\n", cs.Team1, cs.RunsTeam1, cs.Team2, cs.RunsTeam2)
	fmt.Printf("Overs: %d.%d\n", cs.CurrentOver, cs.CurrentBall)
}

func main() {
	fmt.Print("Enter name for team 1: ")
	team1Name, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	team1Name = strings.TrimSpace(team1Name)

	fmt.Print("Enter name for team 2: ")
	team2Name, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	team2Name = strings.TrimSpace(team2Name)

	fmt.Print("Enter total overs per inning: ")
	totalOversStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	totalOversStr = strings.TrimSpace(totalOversStr)
	totalOvers, err := strconv.Atoi(totalOversStr)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		os.Exit(1)
	}

	cricketScoreboard := CricketScoreboard{
		Team1:       team1Name[:len(team1Name)-1],
		Team2:       team2Name[:len(team2Name)-1],
		TotalOvers:  totalOvers,
		CurrentOver: 0,
		CurrentBall: 0,
	}

	for cricketScoreboard.CurrentOver < cricketScoreboard.TotalOvers {
		fmt.Println("\nScoreboard:")
		cricketScoreboard.displayScore()
		fmt.Print("\nEnter runs scored in the current ball: ")
		runsStr, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		runsStr = strings.TrimSpace(runsStr)
		runs, err := strconv.Atoi(runsStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			return
		}

		cricketScoreboard.updateScore(runs)
	}

	fmt.Println("Game Over!")
}
