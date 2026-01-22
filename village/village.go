package village

type VillageElement interface {
	Update()
	FlushInfo() string
}

type Village struct {
	Elements []VillageElement
}

func (v *Village) AddElement(el VillageElement) {
	v.Elements = append(v.Elements, el)
}

func (v Village) UpdateAll() {
	for _, e := range v.Elements {
		e.Update()
	}
}

func (v Village) ShowAllInfo() string {
	info := ""
	for _, e := range v.Elements {
		info = e.FlushInfo()
	}
	return info
}
