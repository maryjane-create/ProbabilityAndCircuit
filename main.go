package main

import "fmt"

func main() {

	fmt.Println(reliabilityInSeries(S1, S2))
}

//Todo: Story Time (NexaScale)

//Todo:     Jumoke is a 2nd year student of Chemistry at the prestigious University of Ibadan. She lives in a room self-contained apartment off-campus. In her room are two switches which she nicknamed S1 and S2, connected to the blue and white bulbs respectively.
//Todo:     S1 is placed a few centimeters from her wooden bed, while S2 is close to the door. From the moment she moved into the apartment, she's noticed some inconsistencies in the switches. Sometimes, they don't respond to triggers, but she's not sure of the number of times this happens.
//Todo:     When she tries to communicate this to an electrician in the neighborhood, he respond by asking her how often the switches misbehave, but she doesn't have the number in her head.
//Todo:     So Jumoke came up with an idea. She will be taking records of the number of times the switches responded to triggers for the next 10 days.

var S1 = []bool{true, true, true, false, true, true, false, true, false, true}
var S2 = []bool{false, false, true, true, true, false, false, true, false, true}

//Todo: The electrician Jumoke consulted came up with an economical and efficient approach to tackle the problem by applying probability theorems

// / This methods counts the number of successful triggers
func countOfTrue(arr []bool) int {
	count := 0
	for _, val := range arr {
		if val == true {
			count++
		}
	}
	return count
}

// This function calculate the probability of successful triggers compared to the total number of triggers
func probabilty(arr []bool) float64 {
	result := float64(countOfTrue(arr)) / float64(len(arr))
	return result
}

// This function estimates how efficient the switches will be when connected side by side

func reliabilityInParallel(arr1 []bool, arr2 []bool) float64 {
	// i.e	when the two switches are connected with each having its own wire, and both connected to a central source cable
	result := (probabilty(S1) + probabilty(S2)) - reliabilityInSeries(arr1, arr2)
	return result
}

// This function estimates how efficient the switches will be when connected with one cable split into two wires
func reliabilityInSeries(arr1 []bool, arr2 []bool) float64 {
	//	i.e when the two switches are connected together with one wire
	result := probabilty(S1) * probabilty(S2)
	return result
}
