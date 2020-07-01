package sign

type SignHandle interface {
	Hash(msg []byte) ([]byte, error)
	Sign(digest []byte) ([]byte, error)
	Verify(sign, digest []byte) (bool, error)
}
