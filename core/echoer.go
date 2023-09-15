package core

type Echoer interface {
	Echo(string) string
}

type MyEchoer struct{}

func (*MyEchoer) Echo(s string) string {
	return s
}
