package main

import "log"

func Perform(config *Config) error {

	log.Println("starting perform")

	repo, err := NewRepository()

	if err != nil {
		return err
	}

	if config.SkipCleanWorkspaceCheck {

		log.Println("skipping clean workspace check")

	} else {

		log.Println("verifying clean workspace")

		if err := repo.VerifyWorkspaceClean(); err != nil {
			return err
		}

	}

	if config.Perform != "" {
		if err := RunInShell(config.Perform); err != nil {
			return err
		}
	}

	log.Println("tagging release")

	if err := repo.Tag("v" + config.Version.String()); err != nil {
		return err
	}

	log.Println("incrementing and committing new version")

	if err := repo.IncrementAndCommit(config.Version, config.Increment); err != nil {
		return err
	}

	if !config.SkipPush {

		log.Println("pushing to origin")
		if err := repo.PushToOrigin(); err != nil {
			return err
		}

	} else {
		log.Println("not pushing to origin")
	}

	return nil
}
