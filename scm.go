package main

import (
	"errors"
	"log"
	"os"

	"github.com/Masterminds/semver/v3"
	git "github.com/go-git/go-git/v5"
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
	}

	log.Println(s.String())
	return errors.New("workspace is not clean")

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
func (r *Repository) IncrementAndCommit(version *semver.Version, incrementMode string) error {

	var next semver.Version
	switch incrementMode {
	case IncrementMajor:
		next = version.IncMajor()
	case IncrementMinor:
		next = version.IncMinor()
	case IncrementPatch:
		next = version.IncPatch()
	default:
		return errors.New("unrecognized increment mode " + incrementMode)
	}

	f, err := os.OpenFile(".version", os.O_WRONLY|os.O_TRUNC, 0600)

	if err != nil {
		return err
	}

	defer func() {
		mustNoError(f.Close())
	}()

	_, err = f.WriteString(next.String() + "\n")

	if err != nil {
		return err
	}

	err = f.Sync()

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
