package codegen

func RunCodeGen(sourceDir string, destDir string, imports []string) error {
	structMetas, err := CollectStructMetas(sourceDir)
	if err != nil {
		return err
	}
	_, structMetasMap, err := GroupStructMetas(structMetas)
	if err != nil {
		return err
	}
	dataMetas := Map2DataMetas(structMetasMap)
	err = CrudGen(dataMetas, destDir, imports)
	if err != nil {
		return err
	}
	return nil
}
