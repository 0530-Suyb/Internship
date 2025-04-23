package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
)

func Encode(msg string) ([]byte, error) {
	length := int32(len(msg)) // 4字节消息长度
	pkg := new(bytes.Buffer)
	err := binary.Write(pkg, binary.LittleEndian, length) // 小端：低位字节排放内存低地址端，用在计算机内；大端：高字节位排放在内存低地址端，用于通信。
	if err != nil {
		return nil, err
	}
	err = binary.Write(pkg, binary.LittleEndian, []byte(msg))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

func Decode(reader *bufio.Reader) (string, error) {
	lenByte, _ := reader.Peek(4)
	lenBuff := bytes.NewBuffer(lenByte)
	var len int32
	err := binary.Read(lenBuff, binary.LittleEndian, &len)
	if err != nil {
		return "", err
	}
	if int32(reader.Buffered()) < len+4 {
		return "", fmt.Errorf("msg data len err")
	}

	pack := make([]byte, int(len+4))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
