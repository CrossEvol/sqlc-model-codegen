package codegen

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCodegen(t *testing.T) {
	err := RunCodeGen("D:\\GOLANG_CODE\\sqlc-model-codegen\\internal\\database\\sqliteDao", "D:\\GOLANG_CODE\\sqlc-model-codegen\\curdGen")
	require.Nil(t, err)
}
