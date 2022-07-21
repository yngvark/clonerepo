package parse_git_uri

import (
	"errors"
	"fmt"
	"regexp"
)

var ErrInvalidGitUri = errors.New("not a valid or supported git URI")

// GetOrgAndRepoFromGitUri parses git URIs looking like this:
//
// git@github.com:someorg/somerepo.git
//
// https://github.com/someorg/somerepo
//
// to
//
// someorg, somerepo
func GetOrgAndRepoFromGitUri(gitUri string) (string, string, error) {
	re := regexp.MustCompile(`(git@github.com:|https://github.com/)(\w+)/(\w+)(\.git)?`)

	match := re.FindStringSubmatch(gitUri)
	if match == nil {
		return "", "", fmt.Errorf("%q: %w", gitUri, ErrInvalidGitUri)
	}

	org := match[2]
	repo := match[3]

	return org, repo, nil
}
