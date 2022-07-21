package parse_git_uri

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGitUriParserErrors(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		gitUri      string
		expectError bool
	}{
		{
			gitUri:      "git@github.com:someone/somerepo.git",
			expectError: false,
		},
		{
			gitUri:      "https://github.com/yngvark/gclone.git",
			expectError: false,
		},
		{
			gitUri:      "https://github.com/someone/somerepo",
			expectError: true,
		},
		{
			gitUri:      "somerepo",
			expectError: true,
		},
		{
			gitUri:      "somerepo/",
			expectError: true,
		},
		{
			gitUri:      "/somerepo",
			expectError: true,
		},
		{
			gitUri:      "someorg/somerepo",
			expectError: true,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.gitUri, func(t *testing.T) {
			t.Parallel()

			_, _, err := GetOrgAndRepoFromGitUri(tc.gitUri)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGitUriParser(t *testing.T) {
	testCases := []struct {
		name       string
		gitUri     string
		expectOrg  string
		expectRepo string
	}{
		{
			name:       "Should get correct org and repo from git native URI",
			gitUri:     "git@github.com:someorg/somerepo.git",
			expectOrg:  "someorg",
			expectRepo: "somerepo",
		},
		{
			name:       "Should get correct org and repo from HTTPS git URI",
			gitUri:     "https://github.com/someorg/somerepo.git",
			expectOrg:  "someorg",
			expectRepo: "somerepo",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			org, repo, err := GetOrgAndRepoFromGitUri(tc.gitUri)
			require.NoError(t, err)

			assert.Equal(t, tc.expectOrg, org)
			assert.Equal(t, tc.expectRepo, repo)
		})
	}
}
