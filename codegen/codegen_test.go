package codegen

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCodegen(t *testing.T) {
	kDebug = false

	err := RunCodeGen("D:\\GOLANG_CODE\\sqlc-model-codegen\\internal\\database\\sqliteDao", "D:\\GOLANG_CODE\\sqlc-model-codegen\\cmd\\dto", []string{"github.com/crossevol/sqlc-model-codegen/internal/database/sqliteDao", "time"})
	require.Nil(t, err)
}
