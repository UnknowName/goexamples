package design_pattern

import (
    "log"
    "testing"
)

func TestNewSchool(t *testing.T) {
    school := NewSchool("peking")
    college := NewCollege("computer")
    college.add(&major{name: "python"})
    college.add(&major{name: "golang"})
    college.print()
    log.Println("_____________________________")
    school.add(college)
    school.print()
}
