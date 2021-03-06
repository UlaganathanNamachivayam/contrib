/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package comment

import (
	"strings"
	"time"

	"github.com/google/go-github/github"
)

// Matcher is an interface to match a comment
type Matcher interface {
	Match(comment *github.IssueComment) bool
}

// CreatedAfter matches comments created after the time
type CreatedAfter time.Time

// Match returns true if the comment is created after the time
func (c CreatedAfter) Match(comment *github.IssueComment) bool {
	if comment == nil || comment.CreatedAt == nil {
		return false
	}
	return comment.CreatedAt.After(time.Time(c))
}

// CreatedBefore matches comments created before the time
type CreatedBefore time.Time

// Match returns true if the comment is created before the time
func (c CreatedBefore) Match(comment *github.IssueComment) bool {
	if comment == nil || comment.CreatedAt == nil {
		return false
	}
	return comment.CreatedAt.Before(time.Time(c))
}

// ValidAuthor validates that a comment has the author set
type ValidAuthor struct{}

// Match if the comment has a valid author
func (ValidAuthor) Match(comment *github.IssueComment) bool {
	return comment != nil && comment.User != nil && comment.User.Login != nil
}

// AuthorLogin matches comment made by this Author
type AuthorLogin string

// Match if the Author is a match (ignoring case)
func (a AuthorLogin) Match(comment *github.IssueComment) bool {
	if !(ValidAuthor{}).Match(comment) {
		return false
	}

	return strings.ToLower(*comment.User.Login) == strings.ToLower(string(a))
}
