//go:build !goverter

package factory

import (
	conv "github.com/moumou/server/biz/conv"
	genconv "github.com/moumou/server/gen/conv"
)

func NewConverter() conv.IConverter {
	return &genconv.IConverterImpl{}
}
