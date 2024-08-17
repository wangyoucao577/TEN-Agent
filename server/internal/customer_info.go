package internal

type CustomerInfo map[string]string

func (u CustomerInfo) Empty() bool {
	return len(u) == 0
}
