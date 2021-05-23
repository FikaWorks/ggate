package cmd

import (
	"github.com/fikaworks/grgate/pkg/config"
	"github.com/fikaworks/grgate/pkg/platforms"
)

func newPlatform() (platform platforms.Platform, err error) {
	if *config.Main.Platform == config.GitlabPlatform {
		platform, err = platforms.NewGitlab(&platforms.GitlabConfig{
			Token: config.Main.Gitlab.Token,
		})
	} else {
		platform, err = platforms.NewGithub(&platforms.GithubConfig{
			AppID:          config.Main.Github.AppID,
			InstallationID: config.Main.Github.InstallationID,
			PrivateKeyPath: config.Main.Github.PrivateKeyPath,
		})
	}
	return
}
