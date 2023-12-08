package jsonHelper

func (j *JsonStr) JsonArrayInit() {
	j.Str = make([]byte, 1)
	j.Str[0] = byte('[')
}

func (j *JsonStr) JsonArrayEnd() {
	j.Str = append(j.Str, []byte("]")...)
}

func (j *JsonStr) JsonArrayAddStr(value string) {
	if len(j.Str) > 1 {
		j.Str = append(j.Str, []byte(",")...)
	}

	j.Str = append(j.Str, []byte("\"")...)
	j.Str = append(j.Str, []byte(value)...)
	j.Str = append(j.Str, []byte("\"")...)

}

func (j *JsonStr) JsonArrayAddJson(value JsonStr) {
	if len(j.Str) > 1 {
		j.Str = append(j.Str, []byte(",")...)
	}

	j.Str = append(j.Str, value.Str...)
}
