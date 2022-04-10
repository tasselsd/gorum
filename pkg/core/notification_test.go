package core_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/tasselsd/gorum/pkg/core"
)

func TestSendActivation(t *testing.T) {
	core.LoadConfig([]string{"", "../../myconfig.yaml"})
	core.RefreshApplication()
	err := core.SendActivation("tasselsd@outlook.com", uuid.NewString())
	if err != nil {
		panic(err)
	}
}
