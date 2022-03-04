package filter

import (
	"strings"

	"github.com/anu1097/golang-mask-utility/masker"
)

type valueFilter struct {
	target   string
	maskType masker.Mtype
}

func ValueFilter(target string) *valueFilter {
	return &valueFilter{
		target: target,
	}
}

func CustomValueFilter(target string, maskType masker.Mtype) *valueFilter {
	return &valueFilter{
		target:   target,
		maskType: maskType,
	}
}

func (x *valueFilter) ReplaceString(s string) string {
	return strings.ReplaceAll(s, x.target, maskerInstance.String(x.maskType, s, GetFilteredLabel()))
}

func (x *valueFilter) MaskString(s string) string {
	return s
}

func (x *valueFilter) ShouldMask(fieldName string, value interface{}, tag string) bool {
	return false
}
