package clonerepo

type Gitter interface {
	Clone(gitUri string, targetCloneDir string) error
	Pull(gitCloneDir string) error
}
