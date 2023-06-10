package test

import (
	"github.com/shine-bright-team/LAAS/v2/initialize"
	"github.com/shine-bright-team/LAAS/v2/utils"
	"gorm.io/gorm/utils/tests"
	"testing"
)

func TestJwtEndCodeAndDecode(t *testing.T) {
	initialize.LookForEnv()

	t.Run("Test JWT Endcode and Decode", func(t *testing.T) {
		token, err := utils.SignToken("1", false)
		if err != nil {
			t.Fatal(err)
		}
		user, err := utils.ValidateToken(token)
		if err != nil {
			t.Fatal(err)
		}
		tests.AssertEqual(t, user.UserId, "1")
		tests.AssertEqual(t, user.IsLender, false)
	})
}
