package jsonHelper

type JsonStr struct {
	Str []byte
}

func (j *JsonStr) CombineWith(anotherJson JsonStr) {
	switch {
	case len(j.Str) <= 2:
		// j is empty
		j.Str = append(anotherJson.Str)
		return
	case len(anotherJson.Str) <= 2:
		// nothing in anotherJson
		return
	}

	j.Str = append(j.Str[:len(j.Str)-1], append([]byte(","), anotherJson.Str[1:]...)...)
	return
}
