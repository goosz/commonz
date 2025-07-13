package commonz

import "reflect"

func TypeNameWithDepth(t reflect.Type, maxDepth int) string {
	return typeNameWithDepth(t, maxDepth)
}
