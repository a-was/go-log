package handlers

import (
	"fmt"
	"net/smtp"
	"sync"

	"github.com/a-was/log"
	"github.com/jordan-wright/email"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type SMTPConfig struct {
	From    string
	To      []string
	Subject string

	Server string
	Auth   smtp.Auth
}

func SMTPHandler(c SMTPConfig) *log.Handler {
	return log.NewHandler(log.HandlerConfig{
		Type:         log.HandlerTypeText,
		Writer:       newSMTPWriter(c),
		WriterSynced: false,
		Encoders: log.HandlerEncoders{
			Time: zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.999"),
		},
		Enabler: zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.ErrorLevel
		}),
	})
}

type smtpWriter struct {
	SMTPConfig
	email email.Email
	mu    sync.Mutex
}

func newSMTPWriter(c SMTPConfig) *smtpWriter {
	return &smtpWriter{
		SMTPConfig: c,
		email: email.Email{
			From:    c.From,
			To:      c.To,
			Subject: c.Subject,
		},
	}
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
