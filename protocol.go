package gecdhm

import "io"

//SendPublicKey writes the public point data to the supplied io.Writer. This can be a local buffer or a socket
//connection
func SendPublicKey(writer io.Writer, publicKey Point) {
	data := publicKey.Marshal()
	writer.Write(data)
}
