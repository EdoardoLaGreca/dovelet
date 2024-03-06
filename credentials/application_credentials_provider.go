package credentials

import (
	"fmt"
	"os"

	"google.golang.org/api/option"
)

// A ApplicationCredentialsProvider retrieves credentials from the current user's home
// directory, and keeps track if those credentials are expired.
type ApplicationCredentialsProvider struct {
	// FilePath holds the path to the credentials file.
	FilePath string
}

// NewApplicationCredentials returns a pointer to a new Credentials object
// wrapping the file provider.
func NewApplicationCredentials(filename string) (*ApplicationCredentialsProvider, error) {
	if filename != "" {
		return &ApplicationCredentialsProvider{
			FilePath: filename,
		}, nil
	}
	fpath, err := credentialsPath()
	if err != nil {
		return nil, err
	}
	return &ApplicationCredentialsProvider{
		FilePath: fpath,
	}, nil
}

// Provide reads and extracts the shared credentials from the current
// users home directory.
func (p *ApplicationCredentialsProvider) Provide() option.ClientOption {
	return option.WithCredentialsFile(p.FilePath)
}

// credentialsPath returns the filename to use to read google application credentials.
// Will return an error if the user's home directory path cannot be found.
func credentialsPath() (string, error) {
	varname := "GOOGLE_APPLICATION_CREDENTIALS"
	fpath := os.Getenv(varname)
	if fpath == "" {
		return "", fmt.Errorf("unable to read " + varname)
	}
	return fpath, nil
}
