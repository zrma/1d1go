package p1100_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"1d1go/boj/p1k/p1100/mocks"
)

func newIOWithMock(t *testing.T) *mocks.MockInOut {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})
	return mocks.NewMockInOut(ctrl)
}
