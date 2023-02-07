package workflow

import (
	"context"
)

type Session struct {
	WorkflowID string `json:"workflow_id"`

	ClientID                 string `json:"client_id,omitempty"`
	RedirectURI              string `json:"redirect_uri,omitempty"`
	SuppressIDPSessionCookie bool   `json:"suppress_idp_session_cookie,omitempty"`
}

type SessionOutput struct {
	WorkflowID  string `json:"workflow_id"`
	ClientID    string `json:"client_id,omitempty"`
	RedirectURI string `json:"redirect_uri,omitempty"`
}

type SessionOptions struct {
	ClientID                 string
	RedirectURI              string
	SuppressIDPSessionCookie bool
}

func NewSession(opts *SessionOptions) *Session {
	return &Session{
		WorkflowID:               newWorkflowID(),
		ClientID:                 opts.ClientID,
		RedirectURI:              opts.RedirectURI,
		SuppressIDPSessionCookie: opts.SuppressIDPSessionCookie,
	}
}

func (s *Session) ToOutput() *SessionOutput {
	return &SessionOutput{
		WorkflowID:  s.WorkflowID,
		ClientID:    s.ClientID,
		RedirectURI: s.RedirectURI,
	}
}

func (s *Session) Context(ctx context.Context) context.Context {
	return context.WithValue(ctx, contextKeyClientID, s.ClientID)
}
