package cmd_test

type TestGitter struct{}

//goland:noinspection GoUnusedParameter
func (t TestGitter) Clone(gitUri string, targetCloneDir string) error {
	return nil
}

//goland:noinspection GoUnusedParameter
func (t TestGitter) Pull(gitCloneDir string) error {
	return nil
}

func newTestGitter() TestGitter {
	return TestGitter{}
}
