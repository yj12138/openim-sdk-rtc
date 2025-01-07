package sdk

var asyncIdInc = uint64(0)

func genAsyncId() uint64 {
	asyncIdInc++
	return asyncIdInc
}

