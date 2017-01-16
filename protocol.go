package gecdhm

import (
	"encoding/binary"
	"io"
)

//Enumerated Packet Types
const (
	ConnectRequest uint8 = iota
	AcceptConnection
	RejectConnection
	KeyTransferRequest
	KeyTransferResponse
)

func getLengthBytes(data []byte) []byte {
	buffer := make([]byte, 10)
	n := binary.PutUvarint(buffer, uint64(len(data)))
	return data[:n]
}

//SendPacket is just a convenience function to send a packet
func SendPacket(writer io.Writer, packetType byte, packetData []byte) {
	data := append([]byte{packetType}, getLengthBytes(packetData)...)
	writer.Write(append(data, packetData...))
}

//ReadPacket is just a convenience function to split the short 1-byte header from the data
func ReadPacket(reader io.Reader, data []byte) (packetType uint8, packetData []byte) {
	return data[0], data[1:]
}
