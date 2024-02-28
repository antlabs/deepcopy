package deepcopy

type options struct {
	// maxDepth is the maximum depth to traverse.
	// If maxDepth is 0, it will be treated as no limit.
	maxDepth int
	// tagName is the tag name to use.
	// If tagName is empty, it will be treated as no tag.
	tagName string
	// // OnlyField is the field name to copy.
	// // If OnlyField is empty, it will be treated as no field.
	// OnlyField string
	modifySrcMap map[string]ModifySrcValue

	modifyDstMap map[string]ModifyDstFieldFunc
}

type Option func(*options)
type ModifyDstFieldFunc func(dstArg interface{}) (newDst interface{})

func WithMaxDepth(maxDepth int) Option {
	return func(o *options) {
		o.maxDepth = maxDepth
	}
}

func WithTagName(tagName string) Option {
	return func(o *options) {
		o.tagName = tagName
	}
}

type ModifySrcValue struct {
	DstFieldName string // copy到的字段
	SrcFieldName string // 被copy的字段
	Callback     ModifySrcFieldFunc
}
type ModifySrcFieldFunc func(srcArg interface{}) (newDst interface{})

// 该函数的作用是在拷贝的时候，插入一段回调函数，修改拷贝的源字段
// 目前只支持基础类型, int, int8, in16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string, bool
func WithModifySrcField(m map[string]ModifySrcValue) Option {
	return func(o *options) {
		o.modifySrcMap = m
	}
}

func WithModifyDstField(m map[string]ModifyDstFieldFunc) Option {
	return func(o *options) {
		o.modifyDstMap = m
	}
}
