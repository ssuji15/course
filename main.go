package main

import "fmt"

type course struct {
	prerequisiteCount int
	requisite         []int
}

func courseOrder(numOfCourse int, prerequisite [][]int) ([]int, error) {

	// Check input constraints
	if numOfCourse <= 0 || numOfCourse > 2000 {
		return nil, fmt.Errorf("Invalid number of courses")
	}
	if len(prerequisite) > numOfCourse*(numOfCourse-1) {
		return nil, fmt.Errorf("number of prerequisites exceeds the maximum possible number")
	}

	var courseMap map[int]course = make(map[int]course)
	var result []int

	// Initialize the coursemap
	for i := 0; i < numOfCourse; i++ {
		courseMap[i] = course{
			prerequisiteCount: 0,
			requisite:         []int{},
		}
	}

	// Populate the courseMap with prerequisites
	for _, pre := range prerequisite {

		if len(pre) != 2 {
			return nil, fmt.Errorf("prerequisite length must be 2")
		}
		if pre[0] < 0 || pre[0] >= numOfCourse {
			return nil, fmt.Errorf("invalid prerequisite course index %d", pre[0])
		}
		if pre[1] < 0 || pre[1] >= numOfCourse {
			return nil, fmt.Errorf("invalid prerequisite course index %d", pre[1])
		}
		if pre[0] == pre[1] {
			return nil, fmt.Errorf("prerequisite course cannot map to itself")
		}

		if c, ok := courseMap[pre[0]]; ok {
			c.prerequisiteCount++
			courseMap[pre[0]] = c
		} else {
			return nil, fmt.Errorf("prerequisite course %d not found in the course list", pre[0])
		}

		if c, ok := courseMap[pre[1]]; ok {
			c.requisite = append(c.requisite, pre[0])
			courseMap[pre[1]] = c
		} else {
			return nil, fmt.Errorf("prerequisite course %d not found in the course list", pre[1])
		}

	}

	queue := make([]int, 0, numOfCourse)

	// Add courses with no prerequisites to the queue
	for key, course := range courseMap {
		if course.prerequisiteCount == 0 {
			queue = append(queue, key)
		}
	}

	// perform topological sort to determine the order of course.

	for len(queue) != 0 {

		c := queue[0]
		queue = queue[1:]
		result = append(result, c)

		if curCourse, ok := courseMap[c]; ok {

			for _, req := range curCourse.requisite {

				if reqCourse, ok := courseMap[req]; ok {
					reqCourse.prerequisiteCount--
					courseMap[req] = reqCourse

					if reqCourse.prerequisiteCount == 0 {
						queue = append(queue, req)
					}
				} else {
					return nil, fmt.Errorf("course %d not found in the course list", req)
				}
			}
		} else {
			return nil, fmt.Errorf("course %d not found in the course list", c)
		}
	}

	// Check if all courses are completed
	if len(result) != numOfCourse {
		return nil, fmt.Errorf("Cannot complete all courses!")
	}

	return result, nil
}

func main() {
	order, err := courseOrder(1, [][]int{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(order)
	}
}
