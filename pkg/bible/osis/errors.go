package osis

import "errors"

var (
	ErrBookNotFound    = errors.New("ErrBookNotFound")
	ErrChapterNotFound = errors.New("ErrChapterNotFound")
	ErrVerseNotFound   = errors.New("ErrVerseNotFound")
)
