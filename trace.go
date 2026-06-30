package main

import (
	"crypto/tls"
	"fmt"
	"log/slog"
	"net/http/httptrace"
)

var Trace = &httptrace.ClientTrace{
	GetConn: func(hostPort string) {
		slog.Info(fmt.Sprintf("GetConn: %s", hostPort))
	},
	GotConn: func(info httptrace.GotConnInfo) {
		slog.Info(fmt.Sprintf("GotConn: reused=%t wasIdle=%t idleTimeMs=%d", info.Reused, info.WasIdle, info.IdleTime.Milliseconds()))
	},
	PutIdleConn: func(err error) {
		slog.Info(fmt.Sprintf("PutIdleConn: %v", err))
	},
	GotFirstResponseByte: func() {
		slog.Info("GotFirstResponseByte")
	},
	DNSStart: func(info httptrace.DNSStartInfo) {
		slog.Info(fmt.Sprintf("DNSStart: host=%s", info.Host))
	},
	DNSDone: func(info httptrace.DNSDoneInfo) {
		slog.Info(fmt.Sprintf("DNSDone: addrs=%v coalesced=%t err=%v", info.Addrs, info.Coalesced, info.Err))
	},
	ConnectStart: func(network, addr string) {
		slog.Info(fmt.Sprintf("ConnectStart: network=%s addr=%s", network, addr))
	},
	ConnectDone: func(network, addr string, err error) {
		slog.Info(fmt.Sprintf("ConnectDone: network=%s addr=%s err=%v", network, addr, err))
	},
	TLSHandshakeStart: func() {
		slog.Info("TLSHandshakeStart")
	},
	TLSHandshakeDone: func(state tls.ConnectionState, err error) {
		slog.Info(fmt.Sprintf("TLSHandshakeDone: version=%s cipherSuite=%s serverName=%s err=%v",
			tls.VersionName(state.Version), tls.CipherSuiteName(state.CipherSuite), state.ServerName, err))
	},
	WroteHeaderField: func(key string, value []string) {
		slog.Info(fmt.Sprintf("WroteHeaderField: %s=%v", key, value))
	},
	WroteHeaders: func() {
		slog.Info("WroteHeaders")
	},
	WroteRequest: func(info httptrace.WroteRequestInfo) {
		slog.Info(fmt.Sprintf("WroteRequest: err=%v", info.Err))
	},
}

var ConnTrace = &httptrace.ClientTrace{
	GotConn: func(info httptrace.GotConnInfo) {
		slog.Info(fmt.Sprintf("GotConn: reused=%t wasIdle=%t idleTimeMs=%d", info.Reused, info.WasIdle, info.IdleTime.Milliseconds()))
	},
}
