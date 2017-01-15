package checksum

func Calc(msg []byte) uint16 {
	sum := 0

	len := len(msg)
	for i := 0; i < len-1; i += 2 {
		sum += int(msg[i])<<8 + int(msg[i+1])
	}
	if len%2 == 1 {
		sum += int(msg[len-1]) << 8
	}

	sum = (sum & 0xffff) + (sum >> 16)
	sum += sum >> 16
	return uint16(^sum)
}
