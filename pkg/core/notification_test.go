package core_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/tasselsd/gorum/pkg/core"
)

func TestSendEmail(t *testing.T) {
	core.LoadConfig("../../myconfig.yaml")
	core.RefreshApplication()
	err := core.SendActivation("319348135@qq.com", uuid.NewString())
	if err != nil {
		panic(err)
	}
}
