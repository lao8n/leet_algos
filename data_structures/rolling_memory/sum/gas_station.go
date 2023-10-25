package data_structures

func canCompleteCircuit(gas []int, cost []int) int {
    startStation, gasInTank, netGasInTank := 0, 0, 0
    for i := 0; i < len(gas); i++ {
        netGas := gas[i] - cost[i]
        if gasInTank + netGas < 0 {
            // reset
            startStation = i + 1
            gasInTank = 0
        } else {
            gasInTank += netGas
        }
        netGasInTank += netGas
    }
    if netGasInTank < 0 {
        return -1
    }
    return startStation
}