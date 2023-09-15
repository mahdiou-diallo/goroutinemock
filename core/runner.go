package core

type Runner struct {
	echoer Echoer
}

func NewRunner(e Echoer) *Runner {
	return &Runner{
		echoer: e,
	}
}

func (r *Runner) EchoSync(s string) string {
	return r.echoer.Echo(s)
}

func (r *Runner) EchoAsync(s string) string {
	c := make(chan string)
	go func() {
		c <- r.echoer.Echo(s)
	}()
	res := <-c
	return res
}
