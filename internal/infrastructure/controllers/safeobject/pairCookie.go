package safeobject

type PairToken struct {
	Access  string
	Refresh string
}

func NewPairToken(access string, refresh string) *PairToken {
	return &PairToken{Access: access, Refresh: refresh}
}
