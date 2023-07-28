package session

import "github.com/ginger-core/errors"

// type ChallengeGenerator func() (string, errors.Error)

type ChallengeGenerator func(chars string, length int) (string, errors.Error)
