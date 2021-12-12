package platforms

import (
	"io"
)

// Success status value
const successStatusValue = "success"

// Platform interface Github and Gitlab
type Platform interface {
	CheckAllStatusSucceeded(string, string, string, []string) (bool, error)
	CreateFile(string, string, string, string, string, string) error
	CreateRelease(string, string, *Release) error
	CreateRepository(string, string, string) error
	CreateStatus(string, string, *Status) error
	DeleteRepository(string, string) error
	GetStatus(string, string, string, string) (*Status, error)
	ListReleases(string, string) ([]*Release, error)
	ListStatuses(string, string, string) ([]*Status, error)
	PublishRelease(string, string, *Release) (bool, error)
	ReadFile(string, string, string) (io.Reader, error)
	UpdateRelease(string, string, *Release) error
}

// Release represent a release regarding the platform
type Release struct {
	// CommitSha attached to the release
	CommitSha string

	// ID of the release, Github use an int, Gitlab use a string
	ID interface{}

	// Name of the release
	Name string

	// Platform, either github or gitlab
	Platform string

	// Tag associated to the release
	Tag string

	// ReleaseNote attached to the release
	ReleaseNote string

	// Published represent the state of the release. For Github it translates to
	// draft. For Gitlab it translates to a future release
	Published bool
}

// Status contains commit status informations
type Status struct {
	// CommitSha
	CommitSha string

	// Name of the status
	Name string

	// State is only used by Github checks, must be one of success or in_progress
	State string

	// Status the commit status:
	// For Github must be one of: queued, in_progress or completed
	// For Gitlab must be one of: pending, running, success, failed or cancelled
	Status string
}
