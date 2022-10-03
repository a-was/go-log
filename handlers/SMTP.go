package handlers

import (
	"fmt"
	"net/smtp"
	"sync"

	"github.com/a-was/go-log"
	"github.com/jordan-wright/email"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type SMTPConfig struct {
	From    string
	To      []string
	Subject string

	Server string // host:port
	Auth   smtp.Auth

	Enabler zapcore.LevelEnabler
}

func SMTPHandler(c SMTPConfig) *log.Handler {
	if c.Enabler == nil {
		c.Enabler = zap.ErrorLevel
	}
	return log.NewHandler(log.HandlerConfig{
		Type: log.HandlerTypeText,
		Writer: &smtpWriter{
			SMTPConfig: c,
			email: email.Email{
				From:    c.From,
				To:      c.To,
				Subject: c.Subject,
			},
		},
		WriterSynced: true,
		Encoders: log.HandlerEncoders{
			Time: zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.999"),
		},
		Enabler: c.Enabler,
	})
}

type smtpWriter struct {
	SMTPConfig
	email email.Email
	mu    sync.Mutex
}

func (w *smtpWriter) Write(p []byte) (int, error) {
	go func(w *smtpWriter) {
		w.mu.Lock()
		defer w.mu.Unlock()
		email := w.email
		email.Text = p
		email.HTML = p
		if err := email.Send(w.Server, w.Auth); err != nil {
			fmt.Println("SMTP handler error:", err.Error())
		}
	}(w)
	return len(p), nil
}
