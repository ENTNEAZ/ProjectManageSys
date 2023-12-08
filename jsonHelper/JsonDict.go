package jsonHelper

import "strconv"

func (j *JsonStr) JsonDictInit() {
	j.Str = make([]byte, 1)
	j.Str[0] = byte('{')
}

func (j *JsonStr) JsonDictEnd() {
	j.Str = append(j.Str, []byte("}")...)
}

func (j *JsonStr) JsonDictAddStrStr(key, value string) {
	if len(j.Str) > 1 {
		j.Str = append(j.Str, []byte(",")...)
	}
	j.Str = append(j.Str, []byte("\"")...)
	j.Str = append(j.Str, []byte(key)...)
	j.Str = append(j.Str, []byte("\":\"")...)
	j.Str = append(j.Str, []byte(value)...)
	j.Str = append(j.Str, []byte("\"")...)
}

func (j *JsonStr) JsonDictAddStrInt(key string, value int) {
	if len(j.Str) > 1 {
		j.Str = append(j.Str, []byte(",")...)
	}
	j.Str = append(j.Str, []byte("\"")...)
	j.Str = append(j.Str, []byte(key)...)
	j.Str = append(j.Str, []byte("\":")...)
	j.Str = append(j.Str, []byte(strconv.Itoa(value))...)
}
