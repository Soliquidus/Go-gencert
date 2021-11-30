package cert

import (
	"testing"
)

// Course tests
func TestValideCertData(t *testing.T) {
	c, err := New("Golang", "Bob", "2021-05-31")
	if err != nil {
		t.Errorf("Cert data should be valid. err = %v", err)
	}
	if c == nil {
		t.Errorf("Cert should be a valid reference. Got nil")
	}

	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course name is not valid. Expected 'GOLANG COURSE', got %v", c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "Bob", "2021-05-31")
	if err == nil {
		t.Errorf("Error should be returned on an empty course")
	}
}

func TestCourseTooLong(t *testing.T) {
	course := "azodihanfopuiagbhoguifhnmlgzogihzezgiojkbzengpiouzehgbpijdbngspoiugjbpzioujegb"
	_, err := New(course, "Bob", "2021-05-31")
	if err == nil {
		t.Errorf("Error should be returned on a too long course name (got %s)", course)
	}
}

// Name tests
func TestNameEmptyValue(t *testing.T) {
	_, err := New("Golang", "", "2021-05-31")
	if err == nil {
		t.Errorf("Error should be returned on empty student name")
	}
}

func TestNameTooLong(t *testing.T) {
	name := "ozefghzogihzgpoichsjgpzuoihgzpeoiughzepogihcspoigzjaheaozidhzgf"
	_, err := New("Golang", name, "2021-05-31")
	if err == nil {
		t.Errorf("Error should be returned on a too long student name (got %s)", name)
	}
}
