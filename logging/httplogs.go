package logging

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// LogHTTPRequest logs an info message with standard fields and extract HTTP specific data from the request
func LogHTTPRequest(req *http.Request) {
	fields := logrus.Fields{
		"request_id":     req.Header.Get("x-request-id"),
		"forwarded_for":  req.Header.Get("x-forwarded-for"),
		"protocol":       req.Proto,
		"remote_address": req.RemoteAddr,
		"url":            req.RequestURI,
		"method":         req.Method,
		"content_length": req.ContentLength,
	}

	logger.WithFields(fields).Info("HTTP request")
}

// LogHTTPResponse logs an info message with standard fields and extract HTTP specific data from the response
func LogHTTPResponse(res *http.Response) {
	fields := logrus.Fields{
		"request_id":     res.Request.Header.Get("x-request-id"),
		"forwarded_for":  res.Request.Header.Get("x-forwarded-for"),
		"protocol":       res.Proto,
		"remote_address": res.Request.RemoteAddr,
		"url":            res.Request.RequestURI,
		"method":         res.Request.Method,
		"content_length": res.ContentLength,
		"status":         res.StatusCode,
	}

	logger.WithFields(fields).Info("HTTP response")
}

// InfoWithRequestID logs an info level message with standard fields and request_id
func InfoWithRequestID(msg, correlationID string) {
	logger.WithField("request_id", correlationID).Info(msg)
}

// WarnWithRequestID logs a warn level message with standard fields and request_id
func WarnWithRequestID(msg, correlationID string) {
	logger.WithField("request_id", correlationID).Warn(msg)
}

// ErrorWithRequestID logs an error level message with standard fields and request_id
func ErrorWithRequestID(msg, correlationID string) {
	logger.WithField("request_id", correlationID).Error(msg)
}

// FatalWithRequestID logs a fatal level message with standard fields and request_id
func FatalWithRequestID(msg, correlationID string) {
	logger.WithField("request_id", correlationID).Fatal(msg)
}
