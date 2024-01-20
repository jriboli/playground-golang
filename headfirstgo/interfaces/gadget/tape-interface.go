package gadget

type Player interface {
	Play(song string)
	Stop()
}