package parse_git_uri_test

import (
	"testing"

	"github.com/yngvark.com/clonerepo/pkg/clonerepo/parse_git_uri"

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
			gitUri:      "https://github.com/someone/somerepo.git",
			expectError: false,
		},
		{
			gitUri:      "https://github.com/someone/somerepo", // No ".git" at the end, which is valid
			expectError: false,
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

			_, _, err := parse_git_uri.GetOrgAndRepoFromGitUri(tc.gitUri)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGitUriParser(t *testing.T) {
	t.Parallel()

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
			name:       "Should get correct org and repo from git native URI",
			gitUri:     "git@github.com:some-org/some-repo.git",
			expectOrg:  "some-org",
			expectRepo: "some-repo",
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

			// When
			org, repo, err := parse_git_uri.GetOrgAndRepoFromGitUri(tc.gitUri)
			require.NoError(t, err)

			// Then
			assert.Equal(t, tc.expectOrg, org)
			assert.Equal(t, tc.expectRepo, repo)
		})
	}
}
