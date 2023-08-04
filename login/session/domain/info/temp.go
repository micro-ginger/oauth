package info

type Temp map[string]any

func (t Temp) Set(k string, v any) {
	t[k] = v
}

func (t Temp) Del(k string) {
	delete(t, k)
}

func (t Temp) Get(k string) any {
	return t[k]
}
