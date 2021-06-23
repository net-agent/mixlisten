package mixlistener

type flexListener struct {
	ProtoListener
}

// Flex 监听Flex协议
func Flex() ProtoListener {
	return &tunnelListener{
		ProtoListener: NewProtobase(FlexName),
	}
}

func (proto *flexListener) Taste(buf []byte) bool {
	return buf != nil && buf[0] == 0 && buf[1] == 0 && buf[2] == 0
}
