package main

import "log"

func Perform(config *Config) error {

	log.Println("starting perform")

	repo, err := NewRepository()

	if err != nil {
		return err
	}

	log.Println("verifying clean workspace")

	if err := repo.VerifyWorkspaceClean(); err != nil {
		return err
	}

	if config.Perform != "" {
		if err := RunInShell(config.Perform); err != nil {
			return err
		}
	}

	log.Println("tagging release")

	if err := repo.Tag(config.Version.Version); err != nil {
		return err
	}

	log.Println("incrementing and commiting new version")

	if err := repo.IncrementAndCommit(config.Version); err != nil {
		return err
	}

	log.Println("pushing to origin")

	if err := repo.PushToOrigin(); err != nil {
		return err
	}

	return nil
}
