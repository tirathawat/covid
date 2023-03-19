package covid

// Provinces converts the records to provinces.
// It will ignore the records that have nil ProvinceID or Province.
func (d Data) Provinces() []Province {
	provinces := make([]Province, 0, len(d.Records))
	for _, r := range d.Records {
		if r.ProvinceID == nil || r.Province == nil {
			continue
		}

		provinces = append(provinces, Province{
			ID:   *r.ProvinceID,
			Name: *r.Province,
		})
	}

	return provinces
}

// Ages converts the records to ages.
// It will convert the nil Age to nil and append it to the result.
func (d Data) Ages() []*int {
	ages := make([]*int, 0, len(d.Records))
	for _, r := range d.Records {
		ages = append(ages, r.Age)
	}

	return ages
}
