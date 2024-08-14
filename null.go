// Package null реализует поддержку nullable-типов без эмуляции NULL указателем nil.
// Используется механизм шаблонов, реализующий nullable для любых типов.
package null

import (
	"database/sql"
	"encoding/json"
	_ "gopkg.in/yaml.v3"
)

// null промежуточный приватный тип, необходимый для сокрытия структуры типа Null
type null[T any] sql.Null[T]

// Null шаблонный nullable-тип. Работает поверх стандартного шаблонного типа sql.Null
type Null[T any] null[T]

// NewVal получает значение и возвращает nullable-значение
func NewVal[T any](v T) Null[T] {
	return Null[T]{V: v, Valid: true}
}

// NewNull возвращает nullable-значение NULL
func NewNull[T any]() Null[T] {
	return Null[T]{}
}

// New получает указатель и возвращает nullable-значение
func New[T any](v *T) Null[T] {
	if v == nil {
		return Null[T]{}
	}
	return Null[T]{V: *v, Valid: true}
}

// NewSQL получает значение типа sql.Null и возвращает nullable-значение
func NewSQL[T any](v sql.Null[T]) Null[T] {
	if v.Valid {
		return Null[T](v)
	}
	return Null[T]{}
}

// SetVal записывает переданное значение (всегда не NULL)
func (n *Null[T]) SetVal(v T) {
	*n = Null[T]{V: v, Valid: true}
}

// SetNull записывает NULL
func (n *Null[T]) SetNull() {
	*n = Null[T]{}
}

// Set записывает значение переданной по указателю переменной. Если передан nil, устанавливается NULL
func (n *Null[T]) Set(v *T) {
	if v == nil {
		*n = Null[T]{}
	} else {
		*n = Null[T]{V: *v, Valid: true}
	}
}

// SetSQL записывает переданное значение типа sql.Null
func (n *Null[T]) SetSQL(v sql.Null[T]) {
	if v.Valid {
		*n = Null[T](v)
	} else {
		*n = Null[T]{}
	}
}

// Get возвращает значение и флаг "значение не NULL"
func (n Null[T]) Get() (T, bool) {
	return n.V, n.Valid
}

// GetVal возвращает значение n, если оно не NULL, или значение v
func (n Null[T]) GetVal(v T) T {
	if n.Valid {
		return n.V
	}
	return v
}

// GetRef возвращает указатель на значение, если оно не NULL, или nil
func (n Null[T]) GetRef() *T {
	if n.Valid {
		return &n.V
	}
	return nil
}

// GetSQL возвращает значение типа sql.Null
func (n Null[T]) GetSQL() sql.Null[T] {
	return sql.Null[T](n)
}

// IsZero предикат "значение равно NULL"
func (n Null[T]) IsZero() bool {
	return !n.Valid
}

// Equal возвращает результат сравнения на равенство - с учётом NULL
func Equal[T comparable](a, b Null[T]) bool {
	return a.Valid == b.Valid && (!a.Valid || a.V == b.V)
}

func (n Null[T]) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.V)
}

func (n *Null[T]) UnmarshalJSON(data []byte) error {
	var x *T
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	n.Set(x)
	return nil
}

func (n Null[T]) MarshalYAML() (interface{}, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.V, nil
}

func (n *Null[T]) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var x *T
	if err := unmarshal(&x); err != nil {
		return err
	}
	n.Set(x)
	return nil
}
