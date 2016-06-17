package pinyin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamMgr(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(T(), "")
	assert.Equal(T(" "), "")

	assert.Equal(T("中国人"), "zhong guo ren")
	assert.Equal(T("中国人", "-"), "zhong-guo-ren")
	assert.Equal(TT("中国人"), "zhōng guó rén")
	assert.Equal(TT("中国人", "-"), "zhōng-guó-rén")
}
