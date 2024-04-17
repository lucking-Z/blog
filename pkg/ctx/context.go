package ctx

type BlogContext struct {
	data map[string]interface{}
}

func (b *BlogContext) GetUID() uint64 {
	if uid, ok := b.data["UID"]; ok {
		res, _ := uid.(uint64)
		return res
	}
	return 0
}

func (b *BlogContext) SetUID(uid uint64) {
	b.data["UID"] = uid
}
