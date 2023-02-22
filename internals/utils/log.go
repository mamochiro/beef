package utils

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
)

func ExtLogError(span opentracing.Span, err error, fields ...log.Field) {
	if err != nil {
		ext.LogError(span, err, fields...)
	}
}
