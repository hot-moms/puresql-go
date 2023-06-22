package sql

import (
	"strconv"
)

// Private functions
func _writeWhereClosure(b *Builder, writeFunc func()) {

	if len(b._whereClosures) == 1 {
		b.writer.WriteString(" WHERE (")
	} else {
		b.writer.WriteString(" AND (")
	}
	writeFunc()
	b.writer.WriteByte(')')

}

func _columnLwGtBase[V int32 | int64 | float32](b *Builder, column string, comparingTo *V, mode string, timestamp bool) {
	if comparingTo != nil && *comparingTo != 0 {
		b._whereClosures = append(b._whereClosures, *comparingTo)
		_writeWhereClosure(b, func() {
			b.writer.WriteString(column)
			b.writer.WriteString(mode)
			if timestamp {
				b.writer.WriteString(`TO_TIMESTAMP(`)
				b.varPlaceholder()
				b.writer.WriteByte(')')
			} else {
				b.varPlaceholder()
			}
		})

	}
}

func (b *Builder) EqualToString(column string, str *string) {
	if str != nil {

		b._whereClosures = append(b._whereClosures, *str)

		_writeWhereClosure(b, func() {
			b.writer.WriteString(column)
			b.writer.WriteString(` = `)
			b.varPlaceholder()
		})
	}
}

func (b *Builder) LikeString(column string, str *string) {
	if str != nil {

		b._whereClosures = append(b._whereClosures, *str)

		_writeWhereClosure(b, func() {
			b.writer.WriteString(column)
			b.writer.WriteString(` LIKE `)
			b.varPlaceholder()
		})
	}
}

func (b *Builder) EqualToInt(column string, num *int32) {
	if num != nil {
		b._whereClosures = append(b._whereClosures, *num)
		_writeWhereClosure(b, func() {

			b.writer.WriteString(column)
			b.writer.WriteString(` = `)
			b.varPlaceholder()

		})

	}
}

func (b *Builder) Limit(num int) {
	b.writer.WriteString(` LIMIT `)
	b.writer.WriteString(strconv.Itoa(num))
}

func (b *Builder) BetweenInts(column string, lowerBound *int32, upperBound *int32) {
	if lowerBound != nil && upperBound != nil {
		b._whereClosures = append(b._whereClosures, *lowerBound)
		_writeWhereClosure(b, func() {
			b.writer.WriteString(column)
			b.writer.WriteString(` BETWEEN `)
			b.varPlaceholder()
			b.writer.WriteString(` AND `)
			b._whereClosures = append(b._whereClosures, *upperBound)
			b.varPlaceholder()
		})

	}
}

func (b *Builder) Offset(num int) {
	b.writer.WriteString(` OFFSET `)
	b.writer.WriteString(strconv.Itoa(num))
}

func (b *Builder) OrderBy(column *string, ordering *string, defaultColumn string) {
	b.writer.WriteString(` ORDER BY `)
	if column != nil && ordering != nil {
		b.writer.WriteString(*column)
		b.writer.WriteByte(' ')
		b.writer.WriteString(*ordering)
	} else {
		b.writer.WriteString(defaultColumn)
	}

}

func (b *Builder) ColumnLwGtInt(column string, num *int32, sorting string) {
	_columnLwGtBase[int32](b, column, num, sorting, false)
}

func (b *Builder) ColumnLwGFloat32(column string, num *float32, sorting string) {
	_columnLwGtBase[float32](b, column, num, sorting, false)
}

// Usefull for time.Time or uuids
func (b *Builder) ColumnLwGtTime(column string, timestamp *int64, sorting string) {
	_columnLwGtBase[int64](b, column, timestamp, sorting, true)
}

func (b *Builder) AnyOfInts(column string, ints []int) {

	if len(ints) == 0 {
		return
	}
	b._whereClosures = append(b._whereClosures, ints)
	_writeWhereClosure(b, func() {
		b.writer.WriteString(column)
		b.writer.WriteString(` = ANY(`)
		b.varPlaceholder()
		b.writer.WriteByte(')')
	})

}

func (b *Builder) ContainedByStrings(column string, strs []string) {

	if len(strs) == 0 {
		return
	}
	b._whereClosures = append(b._whereClosures, strs)
	_writeWhereClosure(b, func() {
		b.writer.WriteString(column)
		b.writer.WriteString(` @> (`)
		b.varPlaceholder()
		b.writer.WriteByte(')')
	})

}

func (b *Builder) ContainedByInts(column string, strs []int32) {

	if len(strs) == 0 {
		return
	}
	b._whereClosures = append(b._whereClosures, strs)
	_writeWhereClosure(b, func() {
		b.writer.WriteString(column)
		b.writer.WriteString(` @> (`)
		b.varPlaceholder()
		b.writer.WriteByte(')')
	})

}

func (b *Builder) AnyOfStrings(column string, strs []string) {

	if len(strs) == 0 {
		return
	}
	b._whereClosures = append(b._whereClosures, strs)
	_writeWhereClosure(b, func() {
		b.writer.WriteString(column)
		b.writer.WriteString(` = ANY(`)
		b.varPlaceholder()
		b.writer.WriteByte(')')
	})

}

func (b *Builder) varPlaceholder() {
	b.writer.WriteByte('$')
	b.writer.WriteString(strconv.Itoa(len(b._whereClosures)))
}
