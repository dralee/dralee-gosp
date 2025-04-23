/*
类型转换
2025.4.23 by dralee
*/
package basic

import (
	"fmt"
)

type BoolType bool

func (b *BoolType) Scan(value interface{}) error {
	switch v := value.(type) {
	case bool:
		*b = BoolType(v)
	case []byte:
		if v == nil {
			*b = false
		} else {
			*b = BoolType(v[0] != 0)
		}
	case uint8:
		*b = BoolType(v != 0)
	case uint16:
		*b = BoolType(v != 0)
	case uint32:
		*b = BoolType(v != 0)
	default:
		return fmt.Errorf("invalid type: %T", v)
	}
	return nil
}
