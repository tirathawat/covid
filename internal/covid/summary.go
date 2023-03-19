package covid

// AgeGroup is the age group of a person.
type AgeGroup string

const (
	AgeGroupUnspecified AgeGroup = "N/A"
	AgeGroupYoungAdults AgeGroup = "0-30"
	AgeGroupMiddleAged  AgeGroup = "31-60"
	AgeGroupElderly     AgeGroup = "61+"
)

// ClassifyAgeGroup classifies the age group of a person.
// It will return AgeGroupUnspecified if the age is nil or less than or equal to zero.
// It will return AgeGroupYoungAdults if the age is between 1 and 30.
// It will return AgeGroupMiddleAged if the age is between 31 and 60.
// It will return AgeGroupElderly if the age is greater than 60.
func ClassifyAgeGroup(age *int) AgeGroup {
	if age == nil {
		return AgeGroupUnspecified
	}

	switch {
	case *age > 60:
		return AgeGroupElderly
	case *age > 30:
		return AgeGroupMiddleAged
	case *age > 0:
		return AgeGroupYoungAdults
	default:
		return AgeGroupUnspecified
	}
}

// ProvinceCount is the count of cases by province.
// The key of the map is the name of the province.
// The value of the map is the count of cases.
type ProvinceCount map[string]uint

// Province represents the province.
// It contains the ID and name of the province.
type Province struct {
	ID   int
	Name string
}

// CountCasesByProvince counts the number of cases by province.
// It will return the count of cases by province.
func CountCasesByProvince(provinces []Province) ProvinceCount {
	nameByID := make(map[int]string, len(provinces))
	for _, p := range provinces {
		nameByID[p.ID] = p.Name
	}

	countByID := make(map[int]uint, len(provinces))
	for _, p := range provinces {
		countByID[p.ID]++
	}

	count := make(ProvinceCount)
	for id, name := range nameByID {
		count[name] = countByID[id]
	}

	return count
}

// AgeGroupCount is the count of cases by age group.
// The key of the map is the age group.
// The value of the map is the count of cases.
type AgeGroupCount map[AgeGroup]uint

// CountCasesByAgeGroup counts the number of cases by age group.
func CountCasesByAgeGroup(ages []*int) AgeGroupCount {
	count := make(AgeGroupCount)
	for _, age := range ages {
		count[ClassifyAgeGroup(age)]++
	}

	return count
}
