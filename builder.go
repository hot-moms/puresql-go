package sql

import (
	"bytes"
)

type Builder struct {
	writer         bytes.Buffer
	_whereClosures []interface{}
}

func (b *Builder) ToSQL() (query string, args []interface{}) {
	return b.writer.String(), b._whereClosures
}

func (b *Builder) AddCustomArg(arg string) {

	b._whereClosures = append(b._whereClosures, arg)

}

func (b *Builder) QueryString() (query string) {
	return b.writer.String()
}

func Init(sqlQuery string) *Builder {
	b := &Builder{}
	b.writer.WriteString(sqlQuery)

	return b
}
