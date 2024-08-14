package null

import (
	"database/sql"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestEmptyVariable(t *testing.T) {
	ass := assert.New(t)

	var s Null[string]
	ass.False(s.Valid, "bad NULL flag")
	ass.Equal(s.V, "", "bad value")

	var m Null[map[int]int]
	ass.False(m.Valid, "bad NULL flag")
	ass.Nil(m.V, "bad value")
}

func TestNewVal(t *testing.T) {
	ass := assert.New(t)

	v := NewVal(25)
	ass.True(v.Valid, "bad NULL flag")
	ass.Equal(v.V, 25, "bad value")

}

func TestNewNull(t *testing.T) {
	ass := assert.New(t)

	v := Null[int]{V: 25, Valid: true}
	v = NewNull[int]()
	ass.False(v.Valid, "bad NULL flag")
	ass.Equal(v.V, 0, "bad value")

}

func TestNew(t *testing.T) {
	ass := assert.New(t)

	i := 25
	v := New[int](&i)
	ass.True(v.Valid, "bad NULL flag")
	ass.Equal(v.V, 25, "bad value")

	v = New[int](nil)
	ass.False(v.Valid, "bad NULL flag")
	ass.Equal(v.V, 0, "bad value")
}

func TestNull_SetVal(t *testing.T) {
	ass := assert.New(t)

	var v Null[int]
	v.SetVal(25)
	ass.True(v.Valid, "bad NULL flag")
	ass.Equal(v.V, 25, "bad value")
}

func TestNull_SetNull(t *testing.T) {
	ass := assert.New(t)

	v := Null[int]{V: 25, Valid: true}
	v.SetNull()
	ass.False(v.Valid, "bad NULL flag")
	ass.Equal(v.V, 0, "bad value")
}

func TestNull_Set(t *testing.T) {
	ass := assert.New(t)

	var v Null[int]
	i := 25
	v.Set(&i)
	ass.True(v.Valid, "bad NULL flag")
	ass.Equal(v.V, 25, "bad value")

	v.Set(nil)
	ass.False(v.Valid, "bad NULL flag")
	ass.Equal(v.V, 0, "bad value")
}

func TestNewSQL(t *testing.T) {
	ass := assert.New(t)

	v := NewSQL(sql.Null[int]{V: 25, Valid: true})
	ass.True(v.Valid, "bad NULL flag")
	ass.Equal(v.V, 25, "bad value")

	v = NewSQL(sql.Null[int]{V: 13, Valid: false})
	ass.False(v.Valid, "bad NULL flag")
	ass.Equal(v.V, 0, "bad value")
}

func TestNull_SetSQL(t *testing.T) {
	ass := assert.New(t)

	var v Null[int]
	v.SetSQL(sql.Null[int]{V: 25, Valid: true})
	ass.True(v.Valid, "bad NULL flag")
	ass.Equal(v.V, 25, "bad value")

	v.SetSQL(sql.Null[int]{V: 13, Valid: false})
	ass.False(v.Valid, "bad NULL flag")
	ass.Equal(v.V, 0, "bad value")
}

func TestNull_Get(t *testing.T) {
	ass := assert.New(t)

	var v Null[int]
	res, flg := v.Get()
	ass.Equal(res, 0, "bad result")
	ass.False(flg, "bad Value flag")

	v.SetVal(25)
	res, flg = v.Get()
	ass.Equal(res, 25, "bad result")
	ass.True(flg, "bad Value flag")
}

func TestNull_GetVal(t *testing.T) {
	ass := assert.New(t)

	var v Null[int]
	ass.Equal(v.GetVal(37), 37, "bad result")

	v.SetVal(25)
	ass.Equal(v.GetVal(37), 25, "bad result")
}

func TestNull_GetRef(t *testing.T) {
	ass := assert.New(t)

	var v Null[int]
	res := v.GetRef()
	ass.Nil(res, "bad result")

	v.SetVal(25)
	res = v.GetRef()
	ass.NotNil(res, "bad result")
	ass.Equal(*res, 25, "bad result")
}

func TestNull_GetSQL(t *testing.T) {
	ass := assert.New(t)

	var v Null[int]
	ass.Equal(v.GetSQL(), sql.Null[int]{}, "bad result")

	v.SetVal(25)
	ass.Equal(v.GetSQL(), sql.Null[int]{25, true}, "bad result")
}

func TestNull_IsZero(t *testing.T) {
	ass := assert.New(t)

	var v Null[int]
	ass.True(v.IsZero(), "bad result")

	v.SetVal(25)
	ass.False(v.IsZero(), "bad result")
}

func TestEqual(t *testing.T) {
	ass := assert.New(t)

	var a, b Null[int]
	ass.True(Equal[int](a, b), sql.Null[int]{}, "bad compare")

	a.SetVal(25)
	ass.False(Equal[int](a, b), sql.Null[int]{}, "bad compare")
	ass.False(Equal[int](b, a), sql.Null[int]{}, "bad compare")

	b.SetVal(13)
	ass.False(Equal[int](a, b), sql.Null[int]{}, "bad compare")
	ass.False(Equal[int](b, a), sql.Null[int]{}, "bad compare")

	b.SetVal(25)
	ass.True(Equal[int](a, b), sql.Null[int]{}, "bad compare")
}

func TestNull_MarshalJSON(t *testing.T) {
	ass := assert.New(t)

	var s Null[string]
	res, err := json.Marshal(s)
	ass.Nil(err, "error %w", err)
	ass.Equal(string(res), "null", "bad result")

	s.SetVal("abc")
	res, err = json.Marshal(s)
	ass.Nil(err, "error %w", err)
	ass.Equal(string(res), "\"abc\"", "bad result")

	var a Null[[]int]
	res, err = json.Marshal(a)
	ass.Nil(err, "error %w", err)
	ass.Equal(string(res), "null", "bad result")

	a.SetVal([]int{1, 2, 3})
	res, err = json.Marshal(a)
	ass.Nil(err, "error %w", err)
	ass.Equal(string(res), "[1,2,3]", "bad result")
}

func TestNull_UnmarshalJSON(t *testing.T) {
	ass := assert.New(t)

	var s Null[string]
	err := json.Unmarshal([]byte("null"), &s)
	ass.Nil(err, "error %w", err)
	ass.False(s.Valid, "bad result")
	ass.Equal(s.V, "", "bad result")

	err = json.Unmarshal([]byte("\"abc\""), &s)
	ass.Nil(err, "error %w", err)
	ass.True(s.Valid, "bad result")
	ass.Equal(s.V, "abc", "bad result")

	var a Null[[]int]
	err = json.Unmarshal([]byte("null"), &a)
	ass.Nil(err, "error %w", err)
	ass.False(a.Valid, "bad result")
	ass.Nil(a.V, "bad result")

	err = json.Unmarshal([]byte("[1,2,3]"), &a)
	ass.Nil(err, "error %w", err)
	ass.True(a.Valid, "bad result")
	ass.Equal(a.V, []int{1, 2, 3}, "bad result")

	err = json.Unmarshal([]byte("{\"a\":\"b\"}"), &a)
	ass.NotNil(err, "no error")
}

func TestNull_MarshalYAML(t *testing.T) {
	ass := assert.New(t)

	var s Null[string]
	res, err := yaml.Marshal(s)
	ass.Nil(err, "error %w", err)
	ass.Equal(string(res), "null\n", "bad result")

	s.SetVal("abc")
	res, err = yaml.Marshal(s)
	ass.Nil(err, "error %w", err)
	ass.Equal(string(res), "abc\n", "bad result")

	var a Null[[]int]
	res, err = yaml.Marshal(a)
	ass.Nil(err, "error %w", err)
	ass.Equal(string(res), "null\n", "bad result")

	a.SetVal([]int{1, 2, 3})
	res, err = yaml.Marshal(a)
	ass.Nil(err, "error %w", err)
	ass.Equal(string(res), "- 1\n- 2\n- 3\n", "bad result")
}

func TestNull_UnmarshalYAML(t *testing.T) {
	ass := assert.New(t)

	var s Null[string]
	err := yaml.Unmarshal([]byte("null\n"), &s)
	ass.Nil(err, "error %w", err)
	ass.False(s.Valid, "bad result")
	ass.Equal(s.V, "", "bad result")

	err = yaml.Unmarshal([]byte("abc\n"), &s)
	ass.Nil(err, "error %w", err)
	ass.True(s.Valid, "bad result")
	ass.Equal(s.V, "abc", "bad result")

	var a Null[[]int]
	err = yaml.Unmarshal([]byte("null\n"), &a)
	ass.Nil(err, "error %w", err)
	ass.False(a.Valid, "bad result")
	ass.Nil(a.V, "bad result")

	err = yaml.Unmarshal([]byte("- 1\n- 2\n- 3\n"), &a)
	ass.Nil(err, "error %w", err)
	ass.True(a.Valid, "bad result")
	ass.Equal(a.V, []int{1, 2, 3}, "bad result")

	err = yaml.Unmarshal([]byte("xyz"), &a)
	ass.NotNil(err, "no error")
}
