package lzma

type readerDict struct {
	buffer
	bufferLen int
}

func newReaderDict(historyLen, bufferLen int) (rd *readerDict, err error) {
	if !(1 <= historyLen && int64(historyLen) < MaxDictLen) {
		return nil, newError("historyLen out of range")
	}
	if bufferLen <= 0 {
		return nil, newError("bufferLen must be greater than zero")
	}
	capacity := historyLen
	if bufferLen > capacity {
		capacity = bufferLen
	}
	rd = &readerDict{bufferLen: bufferLen}
	err = initBuffer(&rd.buffer, capacity)
	return
}

func (rd *readerDict) WriteRep(dist int64, n int) (written int, err error) {
	panic("TODO")
}

func (rd *readerDict) Offset() int64 {
	return rd.end
}

func (rd *readerDict) Byte(dist int) byte {
	panic("TODO")
}

type writerDict struct {
	buffer
}

func newWriterDict(historyLen, bufferLen int) (wd *writerDict, err error) {
	if !(1 <= historyLen && int64(historyLen) < MaxDictLen) {
		return nil, newError("historyLen out of range")
	}
	if bufferLen <= 0 {
		return nil, newError("bufferLen must be greater than zero")
	}
	capacity := historyLen + bufferLen
	wd = &writerDict{}
	err = initBuffer(&wd.buffer, capacity)
	if err != nil {
		return nil, err
	}
	wd.writeLimit = bufferLen
	return wd, nil
}

func (wd *writerDict) Byte(dist int) byte {
	panic("TODO")
}

func (wd *writerDict) Offset() int64 {
	return wd.cursor
}
