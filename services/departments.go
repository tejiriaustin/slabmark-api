package services

type (
	DepartmentService struct {
	}
	DepartmentsServiceInterface interface {
		GetRoles() ([]string, error)
		BelongsTo(name string) bool
		IsValidDepartment(name string) bool
	}
)

var _ DepartmentsServiceInterface = (*DepartmentService)(nil)

func NewDeptService() *DepartmentService {
	return &DepartmentService{}
}

var Departments = []string{"ADMINISTRATOR", "REFINERY", "FRACTIONATION", "STORE", "QUALITY-CONTROL"}

func (s *DepartmentService) GetRoles() ([]string, error) {
	return Departments, nil
}

func (s *DepartmentService) BelongsTo(name string) bool {
	if s.findDepartment(name) == "" {
		return false
	}
	return true
}

func (s *DepartmentService) IsValidDepartment(name string) bool {
	if s.findDepartment(name) == "" {
		return false
	}
	return true
}

func (s *DepartmentService) findDepartment(name string) string {
	for _, r := range Departments {
		if r == name {
			return r
		}
	}
	return ""
}
