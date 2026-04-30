// Package github provides a client for interacting with the GitHub API
// as part of the MCP (Model Context Protocol) server implementation.
package github

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/v57/github"
	"golang.org/x/oauth2"
)

// Client wraps the GitHub API client with additional configuration.
type Client struct {
	gh      *github.Client
	baseURL string
}

// ClientOption is a functional option for configuring the Client.
type ClientOption func(*Client) error

// WithBaseURL sets a custom base URL for the GitHub API (e.g., for GitHub Enterprise).
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		c.baseURL = baseURL
		return nil
	}
}

// NewClient creates a new GitHub API client using the provided token.
// If token is empty, it falls back to the GITHUB_TOKEN environment variable.
func NewClient(token string, opts ...ClientOption) (*Client, error) {
	if token == "" {
		token = os.Getenv("GITHUB_TOKEN")
	}

	if token == "" {
		return nil, fmt.Errorf("GitHub token is required: set GITHUB_TOKEN environment variable or provide a token")
	}

	c := &Client{}

	// Apply options before constructing the underlying client.
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, fmt.Errorf("applying client option: %w", err)
		}
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), ts)

	if c.baseURL != "" {
		ghClient, err := github.NewEnterpriseClient(c.baseURL, c.baseURL, httpClient)
		if err != nil {
			return nil, fmt.Errorf("creating GitHub Enterprise client: %w", err)
		}
		c.gh = ghClient
	} else {
		c.gh = github.NewClient(httpClient)
	}

	return c, nil
}

// GetRawClient returns the underlying go-github client for direct API access.
func (c *Client) GetRawClient() *github.Client {
	return c.gh
}

// CheckAuth verifies that the current token is valid by fetching the
// authenticated user's information.
func (c *Client) CheckAuth(ctx context.Context) (*github.User, *http.Response, error) {
	user, resp, err := c.gh.Users.Get(ctx, "")
	if err != nil {
		return nil, resp, fmt.Errorf("authenticating with GitHub: %w", err)
	}
	return user, resp, nil
}
