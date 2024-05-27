package main

import (
	"errors"
	git "github.com/go-git/go-git/v5"
	"log"
	"os"
)

type Repository struct {
	repository *git.Repository
}

func NewRepository() (*Repository, error) {

	r, err := git.PlainOpen(".")

	if err != nil {
		return nil, err
	}

	return &Repository{repository: r}, nil
}

func (r *Repository) VerifyWorkspaceClean() error {

	w, err := r.repository.Worktree()

	if err != nil {
		return err
	}

	s, err := w.Status()

	if err != nil {
		return err
	}

	if s.IsClean() {

		return nil
	} else {

		log.Println(s.String())

		return errors.New("workspace is not clean")
	}

}

func (r *Repository) Tag(tag string) error {

	log.Printf("tagging version %s\n", tag)

	h, err := r.repository.Head()

	if err != nil {
		return err
	}

	_, err = r.repository.CreateTag(tag, h.Hash(), &git.CreateTagOptions{
		Message: tag,
	})

	return err
}

// TODO - this isn't purely an scm operation
func (r *Repository) IncrementAndCommit(version *Version) error {

	next := version.Increment()

	f, err := os.OpenFile(".version", os.O_WRONLY|os.O_TRUNC, 0644)
	defer f.Close()

	if err != nil {
		return err
	}

	_, err = f.WriteString(next.Version + "\n")

	f.Sync()

	if err != nil {
		return err
	}

	w, err := r.repository.Worktree()

	if err != nil {
		return err
	}

	_, err = w.Add(".version")

	if err != nil {
		return err
	}

	_, err = w.Commit("incrementing version", &git.CommitOptions{})

	return err

}

func (r *Repository) PushToOrigin() error {

	return r.repository.Push(&git.PushOptions{
		RemoteName: "origin",
	})
}
