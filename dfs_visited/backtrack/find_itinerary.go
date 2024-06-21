package data_structures

import "sort"

func findItinerary(tickets [][]string) []string {
	// build adjacency map
	ticketAdjacency := make(map[string][]string)
	ticketMap := make(map[string]int, len(tickets))
	for _, ticket := range tickets {
		from, to := ticket[0], ticket[1]
		ticketAdjacency[from] = append(ticketAdjacency[from], to)
		ticketMap[from+to] += 1
	}
	// sort adjacency list
	for _, neighbors := range ticketAdjacency {
		sort.Strings(neighbors)
	}
	return recurse(ticketAdjacency, "JFK", ticketMap, []string{}, len(tickets)+1)
}

func recurse(ticketsAdjacency map[string][]string, currentAirport string, ticketMap map[string]int, itinerary []string, expectedLengthItinerary int) []string {
	itinerary = append(itinerary, currentAirport) // append current airport to itinerary
	if len(itinerary) == expectedLengthItinerary {
		return itinerary
	}
	// try all possible next steps
	for _, next := range ticketsAdjacency[currentAirport] {
		ticket := currentAirport + next
		if ticketCount, ok := ticketMap[ticket]; ok && ticketCount > 0 {
			// use ticket
			ticketMap[ticket] -= 1
			result := recurse(ticketsAdjacency, next, ticketMap, itinerary, expectedLengthItinerary)
			if result != nil {
				return result // if a valid itinerary is found, return it
			}
			// backtrack, if itinerary is not valid
			ticketMap[ticket] += 1
		}
	}
	return nil
}
