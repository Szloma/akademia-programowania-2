package academy

type Student struct {
	Name       string
	Grades     []int
	Project    int
	Attendance []bool
}

// AverageGrade returns an average grade given a
// slice containing all grades received during a
// semester, rounded to the nearest integer.
func AverageGrade(grades []int) int {
	if len(grades) == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < len(grades); i++ {
		sum += grades[i]
	}
	result := float64(sum) / float64(len(grades))
	b := int(result)
	final := float64(b) - result
	if final > 0 {
		return int(result - 0.50)
	} else {
		return int(result + 0.50)
	}

}

// AttendancePercentage returns a percentage of class
// attendance, given a slice containing information
// whether a student was present (true) of absent (false).
//
// The percentage of attendance is represented as a
// floating-point number ranging from 0 to 1.
func AttendancePercentage(attendance []bool) float64 {
	total := len(attendance)
	presentD := 0
	for _, present := range attendance {
		if present {
			presentD++
		}
	}
	if total == 0 {
		return 0.0
	}
	return float64(presentD) / float64(total)
}

// FinalGrade returns a final grade achieved by a student,
// ranging from 1 to 5.
//
// The final grade is calculated as the average of a project grade
// and an average grade from the semester, with adjustments based
// on the student's attendance. The final grade is rounded
// to the nearest integer.

// If the student's attendance is below 80%, the final grade is
// decreased by 1. If the student's attendance is below 60%, average
// grade is 1 or project grade is 1, the final grade is 1.
func FinalGrade(s Student) int {
	avg := AverageGrade(s.Grades)
	attendance := AttendancePercentage(s.Attendance)
	var adjustment int
	if attendance < 0.6 || avg == 1 || s.Project == 1 {
		return 1
	} else if attendance < 0.8 {
		adjustment = -1
	}
	final := (s.Project + avg) / 2
	final += adjustment
	if final < 1 {
		return 1
	} else if final > 5 {
		return 5
	}
	return final
}

// GradeStudents returns a map of final grades for a given slice of
// Student structs. The key is a student's name and the value is a
// final grade.
func GradeStudents(students []Student) map[string]uint8 {
	final := make(map[string]uint8)
	for _, students := range students {
		final[students.Name] = uint8(FinalGrade(students))
	}
	return final
}
