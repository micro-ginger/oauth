package flow

type Flows map[Section]*Flow

func (f Flows) Get(s Section) *Flow {
	return f[s]
}
