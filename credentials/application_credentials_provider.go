package credentials

import (
	"fmt"
	"os"

	"google.golang.org/api/option"
)

// An ApplicationCredentialsProvider has a file system path to fetch the
// credentials from.
type ApplicationCredentialsProvider struct {
	// filePath holds the path to the credentials file.
	filePath string
}

// NewApplicationCredentials fetches credentials from a file name or path in
// the file system. If the file name is empty, it uses the value of the
// `GOOGLE_APPLICATION_CREDENTIALS` shell variable.
func NewApplicationCredentials(filepath string) (*ApplicationCredentialsProvider, error) {
	if filepath != "" {
		return &ApplicationCredentialsProvider{
			filePath: filepath,
		}, nil
	}
	fpath, err := credentialsPath()
	if err != nil {
		return nil, err
	}
	return &ApplicationCredentialsProvider{
		filePath: fpath,
	}, nil
}

// FilePath returns the path associated with an ApplicationCredentialsProvider.
func (acp ApplicationCredentialsProvider) FilePath() string {
	return acp.filePath
}

// Provide returns the ClientOption instance holding the credentials.
func (p *ApplicationCredentialsProvider) Provide() option.ClientOption {
	return option.WithCredentialsFile(p.filePath)
}

// credentialsPath returns the filename to use to read google application
// credentials. credentialsPath returns an error if the user's home
// directory path cannot be found.
func credentialsPath() (string, error) {
	varname := "GOOGLE_APPLICATION_CREDENTIALS"
	fpath := os.Getenv(varname)
	if fpath == "" {
		return "", fmt.Errorf("unable to read " + varname)
	}
	return fpath, nil
}
