package gocef

type LifeSpanHandler interface {
	OnBeforeClose(b *Browser)
}
