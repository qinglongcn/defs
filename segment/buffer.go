package segment

// ReadSegmentFromBuffer 从buffer读取段
// func ReadSegmentFromBuffer(buffer *bytes.Buffer, segmentType string, xref *FileXref) ([]byte, error) {
// 	xref.mu.RLock()
// 	defer xref.mu.RUnlock()

// 	// 检查 segmentType 是否为空
// 	if len(segmentType) == 0 {
// 		return nil, fmt.Errorf("segmentType cannot be empty")
// 	}

// 	// 查找 xref 表以获取 segment 的信息
// 	entry, ok := xref.XrefTable[segmentType]
// 	if !ok {
// 		return nil, fmt.Errorf("segment not found in xref table")
// 	}

// 	// 跳到buffer中 segment 的位置
// 	if _, err := buffer.Read(make([]byte, entry.Offset)); err != nil {
// 		return nil, err
// 	}

// 	// 读取 segmentType 的长度
// 	var segmentTypeLength uint32
// 	if err := binary.Read(buffer, binary.BigEndian, &segmentTypeLength); err != nil {
// 		return nil, err
// 	}

// 	// 读取 segmentType
// 	readSegmentType := make([]byte, segmentTypeLength)
// 	if _, err := io.ReadFull(buffer, readSegmentType); err != nil {
// 		return nil, err
// 	}

// 	// 读取数据长度和CRC32校验和
// 	var length uint32
// 	var checksum uint32
// 	if err := binary.Read(buffer, binary.BigEndian, &length); err != nil {
// 		return nil, err
// 	}
// 	if err := binary.Read(buffer, binary.BigEndian, &checksum); err != nil {
// 		return nil, err
// 	}

// 	// 读取实际数据
// 	data := make([]byte, length)
// 	if _, err := io.ReadFull(buffer, data); err != nil {
// 		return nil, err
// 	}

// 	// 校验数据
// 	if checksum != crc32.ChecksumIEEE(data) {
// 		return nil, fmt.Errorf("data corruption detected")
// 	}

// 	return data, nil
// }
