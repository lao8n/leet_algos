package data_structures

/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */
// no need to use kahn's topological sorting because we have a tree.

func getImportance(employees []*Employee, id int) int {
	employeeMap := make(map[int]*Employee)
	for _, employee := range employees {
		employeeMap[employee.Id] = employee
	}
	var dfs func(int) int
	dfs = func(id int) int {
		employee := employeeMap[id]
		sum := employee.Importance
		for _, sub := range employee.Subordinates {
			sum += dfs(sub)
		}
		return sum
	}
	return dfs(id)
}
