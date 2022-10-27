// Package daemonapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package daemonapi

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /auth/token)
	PostAuthToken(w http.ResponseWriter, r *http.Request)

	// (GET /daemon/running)
	GetDaemonRunning(w http.ResponseWriter, r *http.Request)

	// (GET /daemon/status)
	GetDaemonStatus(w http.ResponseWriter, r *http.Request, params GetDaemonStatusParams)

	// (POST /daemon/stop)
	PostDaemonStop(w http.ResponseWriter, r *http.Request)

	// (POST /node/monitor)
	PostNodeMonitor(w http.ResponseWriter, r *http.Request)

	// (GET /nodes/info)
	GetNodesInfo(w http.ResponseWriter, r *http.Request)

	// (POST /object/abort)
	PostObjectAbort(w http.ResponseWriter, r *http.Request)

	// (POST /object/clear)
	PostObjectClear(w http.ResponseWriter, r *http.Request)

	// (GET /object/config)
	GetObjectConfig(w http.ResponseWriter, r *http.Request, params GetObjectConfigParams)

	// (GET /object/file)
	GetObjectFile(w http.ResponseWriter, r *http.Request, params GetObjectFileParams)

	// (POST /object/monitor)
	PostObjectMonitor(w http.ResponseWriter, r *http.Request)

	// (GET /object/selector)
	GetObjectSelector(w http.ResponseWriter, r *http.Request, params GetObjectSelectorParams)

	// (POST /object/status)
	PostObjectStatus(w http.ResponseWriter, r *http.Request)

	// (GET /public/openapi)
	GetSwagger(w http.ResponseWriter, r *http.Request)

	// (GET /relay/message)
	GetRelayMessage(w http.ResponseWriter, r *http.Request, params GetRelayMessageParams)

	// (POST /relay/message)
	PostRelayMessage(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// PostAuthToken operation middleware
func (siw *ServerInterfaceWrapper) PostAuthToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostAuthToken(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetDaemonRunning operation middleware
func (siw *ServerInterfaceWrapper) GetDaemonRunning(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetDaemonRunning(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetDaemonStatus operation middleware
func (siw *ServerInterfaceWrapper) GetDaemonStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetDaemonStatusParams

	// ------------- Optional query parameter "namespace" -------------
	if paramValue := r.URL.Query().Get("namespace"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "namespace", r.URL.Query(), &params.Namespace)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "namespace", Err: err})
		return
	}

	// ------------- Optional query parameter "relatives" -------------
	if paramValue := r.URL.Query().Get("relatives"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "relatives", r.URL.Query(), &params.Relatives)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "relatives", Err: err})
		return
	}

	// ------------- Optional query parameter "selector" -------------
	if paramValue := r.URL.Query().Get("selector"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "selector", r.URL.Query(), &params.Selector)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "selector", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetDaemonStatus(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostDaemonStop operation middleware
func (siw *ServerInterfaceWrapper) PostDaemonStop(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostDaemonStop(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostNodeMonitor operation middleware
func (siw *ServerInterfaceWrapper) PostNodeMonitor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostNodeMonitor(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetNodesInfo operation middleware
func (siw *ServerInterfaceWrapper) GetNodesInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetNodesInfo(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostObjectAbort operation middleware
func (siw *ServerInterfaceWrapper) PostObjectAbort(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostObjectAbort(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostObjectClear operation middleware
func (siw *ServerInterfaceWrapper) PostObjectClear(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostObjectClear(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetObjectConfig operation middleware
func (siw *ServerInterfaceWrapper) GetObjectConfig(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetObjectConfigParams

	// ------------- Required query parameter "path" -------------
	if paramValue := r.URL.Query().Get("path"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "path"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "path", r.URL.Query(), &params.Path)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "path", Err: err})
		return
	}

	// ------------- Optional query parameter "evaluate" -------------
	if paramValue := r.URL.Query().Get("evaluate"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "evaluate", r.URL.Query(), &params.Evaluate)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "evaluate", Err: err})
		return
	}

	// ------------- Optional query parameter "impersonate" -------------
	if paramValue := r.URL.Query().Get("impersonate"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "impersonate", r.URL.Query(), &params.Impersonate)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "impersonate", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetObjectConfig(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetObjectFile operation middleware
func (siw *ServerInterfaceWrapper) GetObjectFile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetObjectFileParams

	// ------------- Required query parameter "path" -------------
	if paramValue := r.URL.Query().Get("path"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "path"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "path", r.URL.Query(), &params.Path)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "path", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetObjectFile(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostObjectMonitor operation middleware
func (siw *ServerInterfaceWrapper) PostObjectMonitor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostObjectMonitor(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetObjectSelector operation middleware
func (siw *ServerInterfaceWrapper) GetObjectSelector(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetObjectSelectorParams

	// ------------- Required query parameter "selector" -------------
	if paramValue := r.URL.Query().Get("selector"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "selector"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "selector", r.URL.Query(), &params.Selector)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "selector", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetObjectSelector(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostObjectStatus operation middleware
func (siw *ServerInterfaceWrapper) PostObjectStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostObjectStatus(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetSwagger operation middleware
func (siw *ServerInterfaceWrapper) GetSwagger(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetSwagger(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetRelayMessage operation middleware
func (siw *ServerInterfaceWrapper) GetRelayMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetRelayMessageParams

	// ------------- Required query parameter "nodename" -------------
	if paramValue := r.URL.Query().Get("nodename"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "nodename"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "nodename", r.URL.Query(), &params.Nodename)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "nodename", Err: err})
		return
	}

	// ------------- Required query parameter "cluster_id" -------------
	if paramValue := r.URL.Query().Get("cluster_id"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "cluster_id"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "cluster_id", r.URL.Query(), &params.ClusterId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "cluster_id", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetRelayMessage(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostRelayMessage operation middleware
func (siw *ServerInterfaceWrapper) PostRelayMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostRelayMessage(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/auth/token", wrapper.PostAuthToken)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/daemon/running", wrapper.GetDaemonRunning)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/daemon/status", wrapper.GetDaemonStatus)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/daemon/stop", wrapper.PostDaemonStop)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/node/monitor", wrapper.PostNodeMonitor)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/nodes/info", wrapper.GetNodesInfo)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/object/abort", wrapper.PostObjectAbort)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/object/clear", wrapper.PostObjectClear)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/object/config", wrapper.GetObjectConfig)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/object/file", wrapper.GetObjectFile)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/object/monitor", wrapper.PostObjectMonitor)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/object/selector", wrapper.GetObjectSelector)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/object/status", wrapper.PostObjectStatus)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/public/openapi", wrapper.GetSwagger)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/relay/message", wrapper.GetRelayMessage)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/relay/message", wrapper.PostRelayMessage)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
