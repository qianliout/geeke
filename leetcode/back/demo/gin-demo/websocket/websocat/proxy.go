package websocat

//
// import (
// 	"crypto/tls"
// 	"errors"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net"
// 	"net/http"
// 	"net/url"
// 	"strings"
// 	"sync"
//
// 	"gitlab.com/security-rd/go-pkg/logging"
// )
//
// /*
//  *
//  *  * Licensed to the Apache Software Foundation (ASF) under one or more
//  *  * contributor license agreements.  See the NOTICE file distributed with
//  *  * this work for additional information regarding copyright ownership.
//  *  * The ASF licenses this file to You under the Apache License, Version 2.0
//  *  * (the "License"); you may not use this file except in compliance with
//  *  * the License.  You may obtain a copy of the License at
//  *  *
//  *  *     http://www.apache.org/licenses/LICENSE-2.0
//  *  *
//  *  * Unless required by applicable law or agreed to in writing, software
//  *  * distributed under the License is distributed on an "AS IS" BASIS,
//  *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  *  * See the License for the specific language governing permissions and
//  *  * limitations under the License.
//  *
//  */
//
// const (
// 	WsScheme  = "ws"
// 	WssScheme = "wss"
// 	BufSize   = 1024 * 32
// )
//
// var ErrFormatAddr = errors.New("remote websockets addr format error")
//
// type WebsocketProxy struct {
// 	// ws, wss
// 	scheme string
// 	// The target address: host:port
// 	remoteAddr string
// 	// path
// 	defaultPath string
// 	tlsc        *tls.Config
// 	// Send handshake before callback
// 	beforeHandshake func(r *http.Request) error
// }
//
// type Options func(wp *WebsocketProxy)
//
// // You must carry a port number，ws://ip:80/ssss, wss://ip:443/aaaa
// // ex: ws://ip:port/ajaxchattest
// func NewProxy(addr string, beforeCallback func(r *http.Request) error, options ...Options) (*WebsocketProxy, error) {
// 	u, err := url.Parse(addr)
// 	if err != nil {
// 		logging.Get().Err(err).Str("addr", addr).Msg("NewProxy Parse")
// 		return nil, ErrFormatAddr
// 	}
// 	host, port, err := net.SplitHostPort(u.Host)
// 	if err != nil {
// 		logging.Get().Err(err).Msg("NewProxy")
// 		return nil, ErrFormatAddr
// 	}
// 	logging.Get().Info().Str("addr", addr).Str("host", host).
// 		Str("port", port).Str("scheme", u.Scheme).Msg("NewProxy SplitHostPort")
//
// 	if u.Scheme != WsScheme && u.Scheme != WssScheme {
// 		return nil, ErrFormatAddr
// 	}
// 	wp := &WebsocketProxy{
// 		scheme:          u.Scheme,
// 		remoteAddr:      fmt.Sprintf("%s:%s", host, port),
// 		defaultPath:     u.Path,
// 		beforeHandshake: beforeCallback,
// 	}
// 	if u.Scheme == WssScheme {
// 		wp.tlsc = &tls.Config{InsecureSkipVerify: true}
// 	}
// 	for op := range options {
// 		options[op](wp)
// 	}
// 	return wp, nil
// }
//
// func (wp *WebsocketProxy) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
// 	wp.Proxy(writer, request)
// }
//
// func (wp *WebsocketProxy) Proxy(writer http.ResponseWriter, request *http.Request) {
// 	if strings.ToLower(request.Header.Get("Connection")) != "upgrade" ||
// 		strings.ToLower(request.Header.Get("Upgrade")) != "websocat" {
// 		_, _ = writer.Write([]byte(`Must be a websocat request`))
// 		return
// 	}
// 	hijacker, ok := writer.(http.Hijacker)
// 	if !ok {
// 		logging.Get().Error().Msg("writer not implement http.Hijacker")
// 		return
// 	}
// 	conn, _, err := hijacker.Hijack()
// 	if err != nil {
// 		logging.Get().Err(err).Msg("hijacker.Hijack")
// 		return
// 	}
// 	defer conn.Close()
// 	req := request.Clone(request.Context())
// 	req.URL.Path, req.URL.RawPath, req.RequestURI = wp.defaultPath, wp.defaultPath, wp.defaultPath
// 	req.Host = wp.remoteAddr
// 	if wp.beforeHandshake != nil {
// 		// Add headers, permission authentication + masquerade sources
// 		err = wp.beforeHandshake(req)
// 		if err != nil {
// 			logging.Get().Err(err).Msg("beforeHandshake")
// 			_, _ = writer.Write([]byte(err.Error()))
// 			return
// 		}
// 	}
// 	var remoteConn net.Conn
// 	switch wp.scheme {
// 	case WsScheme:
// 		remoteConn, err = net.Dial("tcp", wp.remoteAddr)
// 	case WssScheme:
// 		remoteConn, err = tls.Dial("tcp", wp.remoteAddr, wp.tlsc)
// 	}
// 	if err != nil {
// 		logging.Get().Err(err).Msg("Proxy Dial")
// 		_, _ = writer.Write([]byte(err.Error()))
// 		return
// 	}
// 	logging.Get().Info().Str("addr", req.URL.String()).Msg("proxy Dial success")
// 	defer remoteConn.Close()
// 	err = req.Write(remoteConn)
// 	if err != nil {
// 		logging.Get().Err(err).Msg("req.Write")
// 		return
// 	}
// 	errChan := make(chan error, 2)
// 	copyConn := func(a, b net.Conn) {
// 		buf := ByteSliceGet(BufSize)
// 		defer ByteSlicePut(buf)
// 		_, err := io.CopyBuffer(a, b, buf)
// 		errChan <- err
// 	}
// 	go copyConn(conn, remoteConn) // response
// 	go copyConn(remoteConn, conn) // request
// 	select {
// 	case err = <-errChan:
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}
// }
//
// func SetTLSConfig(tlsc *tls.Config) Options {
// 	return func(wp *WebsocketProxy) {
// 		wp.tlsc = tlsc
// 	}
// }
//
// var (
// 	byteSlicePool = sync.Pool{
// 		New: func() interface{} {
// 			return []byte{}
// 		},
// 	}
// 	byteSliceChan = make(chan []byte, 10)
// )
//
// func ByteSliceGet(length int) (data []byte) {
// 	select {
// 	case data = <-byteSliceChan:
// 	default:
// 		data = byteSlicePool.Get().([]byte)[:0]
// 	}
//
// 	if cap(data) < length {
// 		data = make([]byte, length)
// 	} else {
// 		data = data[:length]
// 	}
//
// 	return data
// }
//
// func ByteSlicePut(data []byte) {
// 	select {
// 	case byteSliceChan <- data:
// 	default:
// 		byteSlicePool.Put(data) //nolint:staticcheck
// 	}
// }
