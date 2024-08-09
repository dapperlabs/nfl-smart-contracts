package nfl

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEmbed(t *testing.T) {
	content, err := os.ReadFile("scripts/user/account_is_all_setup.cdc")
	assert.NoError(t, err)
	assert.Equal(t, UserAccountIsAllSetup, content)
}
