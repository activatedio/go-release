package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Version struct {
	Version            string
	Prefix             string
	IncrementingSuffix int
}

func ParseVersion(input string) (*Version, error) {

	pattern, err := regexp.Compile("^(?P<prefix>v.*)\\.(?P<suffix>\\d+)$")

	if err != nil {
		return nil, err
	}

	got := pattern.FindStringSubmatch(input)

	if len(got) != 3 {
		return nil, errors.New("invalid version string")
	}

	suffix, err := strconv.Atoi(got[2])

	if err != nil {
		return nil, err
	}

	return &Version{
		Version:            got[0],
		Prefix:             got[1],
		IncrementingSuffix: suffix,
	}, nil
}

func (v *Version) Increment() *Version {

	next := v.IncrementingSuffix + 1

	return &Version{
		Version:            fmt.Sprintf("%s.%d", v.Prefix, next),
		Prefix:             v.Prefix,
		IncrementingSuffix: next,
	}
}
