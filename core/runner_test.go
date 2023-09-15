package core

import (
	"log"

	"testing"

	"github.com/mahdiou-diallo/goroutinemock/mocks/core"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type RunnerSuite struct {
	suite.Suite
	logger *log.Logger
}

func (s *RunnerSuite) SetupSuite() {
	s.logger = log.Default()
}

func TestSuite_Package(t *testing.T) {
	suite.Run(t, &RunnerSuite{})
}

func (s *RunnerSuite) Test_Sync() {
	m := core.NewMockEchoer(s.T())
	runner := Runner{
		echoer: m,
	}
	m.EXPECT().Echo(mock.Anything).RunAndReturn(func(s string) string { return s }).Once()

	res := runner.EchoSync("hello")

	m.AssertExpectations(s.T())
	require.Equal(s.T(), 32, res)
}

func (s *RunnerSuite) Test_Async() {
	m := core.NewMockEchoer(s.T())
	runner := Runner{
		echoer: m,
	}
	m.EXPECT().Echo(mock.Anything).RunAndReturn(func(s string) string { return s }).Once()

	res := runner.EchoAsync("hello")

	m.AssertExpectations(s.T())
	require.Equal(s.T(), 32, res)
}

func (s *RunnerSuite) Test_SyncUnexpected() {
	m := core.NewMockEchoer(s.T())
	runner := Runner{
		echoer: m,
	}

	// missing or wrong expectation

	res := runner.EchoSync("hello")

	m.AssertExpectations(s.T())
	require.Equal(s.T(), 32, res)
}

func (s *RunnerSuite) Test_AsyncUnexpected() {
	m := core.NewMockEchoer(s.T())
	runner := Runner{
		echoer: m,
	}

	res := runner.EchoAsync("hello")

	m.AssertExpectations(s.T())
	require.Equal(s.T(), 32, res)
}

func (s *RunnerSuite) Test_AsyncPanic() {
	m := core.NewMockEchoer(s.T())
	runner := Runner{
		echoer: m,
	}

	m.EXPECT().Echo(mock.Anything).RunAndReturn(func(string) string {
		panic("adder panic")
	})

	res := runner.EchoSync("hello")

	m.AssertExpectations(s.T())
	require.Equal(s.T(), 32, res)
}
