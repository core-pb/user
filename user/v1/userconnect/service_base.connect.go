// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: user/v1/service_base.proto

package userconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/core-pb/user/user/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// BaseName is the fully-qualified name of the Base service.
	BaseName = "user.v1.Base"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// BaseListAuthenticateProcedure is the fully-qualified name of the Base's ListAuthenticate RPC.
	BaseListAuthenticateProcedure = "/user.v1.Base/ListAuthenticate"
	// BaseAddAuthenticateProcedure is the fully-qualified name of the Base's AddAuthenticate RPC.
	BaseAddAuthenticateProcedure = "/user.v1.Base/AddAuthenticate"
	// BaseSetAuthenticateProcedure is the fully-qualified name of the Base's SetAuthenticate RPC.
	BaseSetAuthenticateProcedure = "/user.v1.Base/SetAuthenticate"
	// BaseDeleteAuthenticateProcedure is the fully-qualified name of the Base's DeleteAuthenticate RPC.
	BaseDeleteAuthenticateProcedure = "/user.v1.Base/DeleteAuthenticate"
	// BaseListUserProcedure is the fully-qualified name of the Base's ListUser RPC.
	BaseListUserProcedure = "/user.v1.Base/ListUser"
	// BaseAddUserProcedure is the fully-qualified name of the Base's AddUser RPC.
	BaseAddUserProcedure = "/user.v1.Base/AddUser"
	// BaseSetUserProcedure is the fully-qualified name of the Base's SetUser RPC.
	BaseSetUserProcedure = "/user.v1.Base/SetUser"
	// BaseSetUserInfoProcedure is the fully-qualified name of the Base's SetUserInfo RPC.
	BaseSetUserInfoProcedure = "/user.v1.Base/SetUserInfo"
	// BaseDeleteUserProcedure is the fully-qualified name of the Base's DeleteUser RPC.
	BaseDeleteUserProcedure = "/user.v1.Base/DeleteUser"
	// BaseSetUserAuthProcedure is the fully-qualified name of the Base's SetUserAuth RPC.
	BaseSetUserAuthProcedure = "/user.v1.Base/SetUserAuth"
	// BaseDeleteUserAuthProcedure is the fully-qualified name of the Base's DeleteUserAuth RPC.
	BaseDeleteUserAuthProcedure = "/user.v1.Base/DeleteUserAuth"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	baseServiceDescriptor                  = v1.File_user_v1_service_base_proto.Services().ByName("Base")
	baseListAuthenticateMethodDescriptor   = baseServiceDescriptor.Methods().ByName("ListAuthenticate")
	baseAddAuthenticateMethodDescriptor    = baseServiceDescriptor.Methods().ByName("AddAuthenticate")
	baseSetAuthenticateMethodDescriptor    = baseServiceDescriptor.Methods().ByName("SetAuthenticate")
	baseDeleteAuthenticateMethodDescriptor = baseServiceDescriptor.Methods().ByName("DeleteAuthenticate")
	baseListUserMethodDescriptor           = baseServiceDescriptor.Methods().ByName("ListUser")
	baseAddUserMethodDescriptor            = baseServiceDescriptor.Methods().ByName("AddUser")
	baseSetUserMethodDescriptor            = baseServiceDescriptor.Methods().ByName("SetUser")
	baseSetUserInfoMethodDescriptor        = baseServiceDescriptor.Methods().ByName("SetUserInfo")
	baseDeleteUserMethodDescriptor         = baseServiceDescriptor.Methods().ByName("DeleteUser")
	baseSetUserAuthMethodDescriptor        = baseServiceDescriptor.Methods().ByName("SetUserAuth")
	baseDeleteUserAuthMethodDescriptor     = baseServiceDescriptor.Methods().ByName("DeleteUserAuth")
)

// BaseClient is a client for the user.v1.Base service.
type BaseClient interface {
	ListAuthenticate(context.Context, *connect.Request[v1.ListAuthenticateRequest]) (*connect.Response[v1.ListAuthenticateResponse], error)
	AddAuthenticate(context.Context, *connect.Request[v1.AddAuthenticateRequest]) (*connect.Response[v1.AddAuthenticateResponse], error)
	SetAuthenticate(context.Context, *connect.Request[v1.SetAuthenticateRequest]) (*connect.Response[v1.SetAuthenticateResponse], error)
	DeleteAuthenticate(context.Context, *connect.Request[v1.DeleteAuthenticateRequest]) (*connect.Response[v1.DeleteAuthenticateResponse], error)
	ListUser(context.Context, *connect.Request[v1.ListUserRequest]) (*connect.Response[v1.ListUserResponse], error)
	AddUser(context.Context, *connect.Request[v1.AddUserRequest]) (*connect.Response[v1.AddUserResponse], error)
	SetUser(context.Context, *connect.Request[v1.SetUserRequest]) (*connect.Response[v1.SetUserResponse], error)
	SetUserInfo(context.Context, *connect.Request[v1.SetUserInfoRequest]) (*connect.Response[v1.SetUserInfoResponse], error)
	DeleteUser(context.Context, *connect.Request[v1.DeleteUserRequest]) (*connect.Response[v1.DeleteUserResponse], error)
	SetUserAuth(context.Context, *connect.Request[v1.SetUserAuthRequest]) (*connect.Response[v1.SetUserAuthResponse], error)
	DeleteUserAuth(context.Context, *connect.Request[v1.DeleteUserAuthRequest]) (*connect.Response[v1.DeleteUserAuthResponse], error)
}

// NewBaseClient constructs a client for the user.v1.Base service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewBaseClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) BaseClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &baseClient{
		listAuthenticate: connect.NewClient[v1.ListAuthenticateRequest, v1.ListAuthenticateResponse](
			httpClient,
			baseURL+BaseListAuthenticateProcedure,
			connect.WithSchema(baseListAuthenticateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		addAuthenticate: connect.NewClient[v1.AddAuthenticateRequest, v1.AddAuthenticateResponse](
			httpClient,
			baseURL+BaseAddAuthenticateProcedure,
			connect.WithSchema(baseAddAuthenticateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		setAuthenticate: connect.NewClient[v1.SetAuthenticateRequest, v1.SetAuthenticateResponse](
			httpClient,
			baseURL+BaseSetAuthenticateProcedure,
			connect.WithSchema(baseSetAuthenticateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteAuthenticate: connect.NewClient[v1.DeleteAuthenticateRequest, v1.DeleteAuthenticateResponse](
			httpClient,
			baseURL+BaseDeleteAuthenticateProcedure,
			connect.WithSchema(baseDeleteAuthenticateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listUser: connect.NewClient[v1.ListUserRequest, v1.ListUserResponse](
			httpClient,
			baseURL+BaseListUserProcedure,
			connect.WithSchema(baseListUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		addUser: connect.NewClient[v1.AddUserRequest, v1.AddUserResponse](
			httpClient,
			baseURL+BaseAddUserProcedure,
			connect.WithSchema(baseAddUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		setUser: connect.NewClient[v1.SetUserRequest, v1.SetUserResponse](
			httpClient,
			baseURL+BaseSetUserProcedure,
			connect.WithSchema(baseSetUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		setUserInfo: connect.NewClient[v1.SetUserInfoRequest, v1.SetUserInfoResponse](
			httpClient,
			baseURL+BaseSetUserInfoProcedure,
			connect.WithSchema(baseSetUserInfoMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteUser: connect.NewClient[v1.DeleteUserRequest, v1.DeleteUserResponse](
			httpClient,
			baseURL+BaseDeleteUserProcedure,
			connect.WithSchema(baseDeleteUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		setUserAuth: connect.NewClient[v1.SetUserAuthRequest, v1.SetUserAuthResponse](
			httpClient,
			baseURL+BaseSetUserAuthProcedure,
			connect.WithSchema(baseSetUserAuthMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteUserAuth: connect.NewClient[v1.DeleteUserAuthRequest, v1.DeleteUserAuthResponse](
			httpClient,
			baseURL+BaseDeleteUserAuthProcedure,
			connect.WithSchema(baseDeleteUserAuthMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// baseClient implements BaseClient.
type baseClient struct {
	listAuthenticate   *connect.Client[v1.ListAuthenticateRequest, v1.ListAuthenticateResponse]
	addAuthenticate    *connect.Client[v1.AddAuthenticateRequest, v1.AddAuthenticateResponse]
	setAuthenticate    *connect.Client[v1.SetAuthenticateRequest, v1.SetAuthenticateResponse]
	deleteAuthenticate *connect.Client[v1.DeleteAuthenticateRequest, v1.DeleteAuthenticateResponse]
	listUser           *connect.Client[v1.ListUserRequest, v1.ListUserResponse]
	addUser            *connect.Client[v1.AddUserRequest, v1.AddUserResponse]
	setUser            *connect.Client[v1.SetUserRequest, v1.SetUserResponse]
	setUserInfo        *connect.Client[v1.SetUserInfoRequest, v1.SetUserInfoResponse]
	deleteUser         *connect.Client[v1.DeleteUserRequest, v1.DeleteUserResponse]
	setUserAuth        *connect.Client[v1.SetUserAuthRequest, v1.SetUserAuthResponse]
	deleteUserAuth     *connect.Client[v1.DeleteUserAuthRequest, v1.DeleteUserAuthResponse]
}

// ListAuthenticate calls user.v1.Base.ListAuthenticate.
func (c *baseClient) ListAuthenticate(ctx context.Context, req *connect.Request[v1.ListAuthenticateRequest]) (*connect.Response[v1.ListAuthenticateResponse], error) {
	return c.listAuthenticate.CallUnary(ctx, req)
}

// AddAuthenticate calls user.v1.Base.AddAuthenticate.
func (c *baseClient) AddAuthenticate(ctx context.Context, req *connect.Request[v1.AddAuthenticateRequest]) (*connect.Response[v1.AddAuthenticateResponse], error) {
	return c.addAuthenticate.CallUnary(ctx, req)
}

// SetAuthenticate calls user.v1.Base.SetAuthenticate.
func (c *baseClient) SetAuthenticate(ctx context.Context, req *connect.Request[v1.SetAuthenticateRequest]) (*connect.Response[v1.SetAuthenticateResponse], error) {
	return c.setAuthenticate.CallUnary(ctx, req)
}

// DeleteAuthenticate calls user.v1.Base.DeleteAuthenticate.
func (c *baseClient) DeleteAuthenticate(ctx context.Context, req *connect.Request[v1.DeleteAuthenticateRequest]) (*connect.Response[v1.DeleteAuthenticateResponse], error) {
	return c.deleteAuthenticate.CallUnary(ctx, req)
}

// ListUser calls user.v1.Base.ListUser.
func (c *baseClient) ListUser(ctx context.Context, req *connect.Request[v1.ListUserRequest]) (*connect.Response[v1.ListUserResponse], error) {
	return c.listUser.CallUnary(ctx, req)
}

// AddUser calls user.v1.Base.AddUser.
func (c *baseClient) AddUser(ctx context.Context, req *connect.Request[v1.AddUserRequest]) (*connect.Response[v1.AddUserResponse], error) {
	return c.addUser.CallUnary(ctx, req)
}

// SetUser calls user.v1.Base.SetUser.
func (c *baseClient) SetUser(ctx context.Context, req *connect.Request[v1.SetUserRequest]) (*connect.Response[v1.SetUserResponse], error) {
	return c.setUser.CallUnary(ctx, req)
}

// SetUserInfo calls user.v1.Base.SetUserInfo.
func (c *baseClient) SetUserInfo(ctx context.Context, req *connect.Request[v1.SetUserInfoRequest]) (*connect.Response[v1.SetUserInfoResponse], error) {
	return c.setUserInfo.CallUnary(ctx, req)
}

// DeleteUser calls user.v1.Base.DeleteUser.
func (c *baseClient) DeleteUser(ctx context.Context, req *connect.Request[v1.DeleteUserRequest]) (*connect.Response[v1.DeleteUserResponse], error) {
	return c.deleteUser.CallUnary(ctx, req)
}

// SetUserAuth calls user.v1.Base.SetUserAuth.
func (c *baseClient) SetUserAuth(ctx context.Context, req *connect.Request[v1.SetUserAuthRequest]) (*connect.Response[v1.SetUserAuthResponse], error) {
	return c.setUserAuth.CallUnary(ctx, req)
}

// DeleteUserAuth calls user.v1.Base.DeleteUserAuth.
func (c *baseClient) DeleteUserAuth(ctx context.Context, req *connect.Request[v1.DeleteUserAuthRequest]) (*connect.Response[v1.DeleteUserAuthResponse], error) {
	return c.deleteUserAuth.CallUnary(ctx, req)
}

// BaseHandler is an implementation of the user.v1.Base service.
type BaseHandler interface {
	ListAuthenticate(context.Context, *connect.Request[v1.ListAuthenticateRequest]) (*connect.Response[v1.ListAuthenticateResponse], error)
	AddAuthenticate(context.Context, *connect.Request[v1.AddAuthenticateRequest]) (*connect.Response[v1.AddAuthenticateResponse], error)
	SetAuthenticate(context.Context, *connect.Request[v1.SetAuthenticateRequest]) (*connect.Response[v1.SetAuthenticateResponse], error)
	DeleteAuthenticate(context.Context, *connect.Request[v1.DeleteAuthenticateRequest]) (*connect.Response[v1.DeleteAuthenticateResponse], error)
	ListUser(context.Context, *connect.Request[v1.ListUserRequest]) (*connect.Response[v1.ListUserResponse], error)
	AddUser(context.Context, *connect.Request[v1.AddUserRequest]) (*connect.Response[v1.AddUserResponse], error)
	SetUser(context.Context, *connect.Request[v1.SetUserRequest]) (*connect.Response[v1.SetUserResponse], error)
	SetUserInfo(context.Context, *connect.Request[v1.SetUserInfoRequest]) (*connect.Response[v1.SetUserInfoResponse], error)
	DeleteUser(context.Context, *connect.Request[v1.DeleteUserRequest]) (*connect.Response[v1.DeleteUserResponse], error)
	SetUserAuth(context.Context, *connect.Request[v1.SetUserAuthRequest]) (*connect.Response[v1.SetUserAuthResponse], error)
	DeleteUserAuth(context.Context, *connect.Request[v1.DeleteUserAuthRequest]) (*connect.Response[v1.DeleteUserAuthResponse], error)
}

// NewBaseHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewBaseHandler(svc BaseHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	baseListAuthenticateHandler := connect.NewUnaryHandler(
		BaseListAuthenticateProcedure,
		svc.ListAuthenticate,
		connect.WithSchema(baseListAuthenticateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	baseAddAuthenticateHandler := connect.NewUnaryHandler(
		BaseAddAuthenticateProcedure,
		svc.AddAuthenticate,
		connect.WithSchema(baseAddAuthenticateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	baseSetAuthenticateHandler := connect.NewUnaryHandler(
		BaseSetAuthenticateProcedure,
		svc.SetAuthenticate,
		connect.WithSchema(baseSetAuthenticateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	baseDeleteAuthenticateHandler := connect.NewUnaryHandler(
		BaseDeleteAuthenticateProcedure,
		svc.DeleteAuthenticate,
		connect.WithSchema(baseDeleteAuthenticateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	baseListUserHandler := connect.NewUnaryHandler(
		BaseListUserProcedure,
		svc.ListUser,
		connect.WithSchema(baseListUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	baseAddUserHandler := connect.NewUnaryHandler(
		BaseAddUserProcedure,
		svc.AddUser,
		connect.WithSchema(baseAddUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	baseSetUserHandler := connect.NewUnaryHandler(
		BaseSetUserProcedure,
		svc.SetUser,
		connect.WithSchema(baseSetUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	baseSetUserInfoHandler := connect.NewUnaryHandler(
		BaseSetUserInfoProcedure,
		svc.SetUserInfo,
		connect.WithSchema(baseSetUserInfoMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	baseDeleteUserHandler := connect.NewUnaryHandler(
		BaseDeleteUserProcedure,
		svc.DeleteUser,
		connect.WithSchema(baseDeleteUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	baseSetUserAuthHandler := connect.NewUnaryHandler(
		BaseSetUserAuthProcedure,
		svc.SetUserAuth,
		connect.WithSchema(baseSetUserAuthMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	baseDeleteUserAuthHandler := connect.NewUnaryHandler(
		BaseDeleteUserAuthProcedure,
		svc.DeleteUserAuth,
		connect.WithSchema(baseDeleteUserAuthMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/user.v1.Base/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case BaseListAuthenticateProcedure:
			baseListAuthenticateHandler.ServeHTTP(w, r)
		case BaseAddAuthenticateProcedure:
			baseAddAuthenticateHandler.ServeHTTP(w, r)
		case BaseSetAuthenticateProcedure:
			baseSetAuthenticateHandler.ServeHTTP(w, r)
		case BaseDeleteAuthenticateProcedure:
			baseDeleteAuthenticateHandler.ServeHTTP(w, r)
		case BaseListUserProcedure:
			baseListUserHandler.ServeHTTP(w, r)
		case BaseAddUserProcedure:
			baseAddUserHandler.ServeHTTP(w, r)
		case BaseSetUserProcedure:
			baseSetUserHandler.ServeHTTP(w, r)
		case BaseSetUserInfoProcedure:
			baseSetUserInfoHandler.ServeHTTP(w, r)
		case BaseDeleteUserProcedure:
			baseDeleteUserHandler.ServeHTTP(w, r)
		case BaseSetUserAuthProcedure:
			baseSetUserAuthHandler.ServeHTTP(w, r)
		case BaseDeleteUserAuthProcedure:
			baseDeleteUserAuthHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedBaseHandler returns CodeUnimplemented from all methods.
type UnimplementedBaseHandler struct{}

func (UnimplementedBaseHandler) ListAuthenticate(context.Context, *connect.Request[v1.ListAuthenticateRequest]) (*connect.Response[v1.ListAuthenticateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("user.v1.Base.ListAuthenticate is not implemented"))
}

func (UnimplementedBaseHandler) AddAuthenticate(context.Context, *connect.Request[v1.AddAuthenticateRequest]) (*connect.Response[v1.AddAuthenticateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("user.v1.Base.AddAuthenticate is not implemented"))
}

func (UnimplementedBaseHandler) SetAuthenticate(context.Context, *connect.Request[v1.SetAuthenticateRequest]) (*connect.Response[v1.SetAuthenticateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("user.v1.Base.SetAuthenticate is not implemented"))
}

func (UnimplementedBaseHandler) DeleteAuthenticate(context.Context, *connect.Request[v1.DeleteAuthenticateRequest]) (*connect.Response[v1.DeleteAuthenticateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("user.v1.Base.DeleteAuthenticate is not implemented"))
}

func (UnimplementedBaseHandler) ListUser(context.Context, *connect.Request[v1.ListUserRequest]) (*connect.Response[v1.ListUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("user.v1.Base.ListUser is not implemented"))
}

func (UnimplementedBaseHandler) AddUser(context.Context, *connect.Request[v1.AddUserRequest]) (*connect.Response[v1.AddUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("user.v1.Base.AddUser is not implemented"))
}

func (UnimplementedBaseHandler) SetUser(context.Context, *connect.Request[v1.SetUserRequest]) (*connect.Response[v1.SetUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("user.v1.Base.SetUser is not implemented"))
}

func (UnimplementedBaseHandler) SetUserInfo(context.Context, *connect.Request[v1.SetUserInfoRequest]) (*connect.Response[v1.SetUserInfoResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("user.v1.Base.SetUserInfo is not implemented"))
}

func (UnimplementedBaseHandler) DeleteUser(context.Context, *connect.Request[v1.DeleteUserRequest]) (*connect.Response[v1.DeleteUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("user.v1.Base.DeleteUser is not implemented"))
}

func (UnimplementedBaseHandler) SetUserAuth(context.Context, *connect.Request[v1.SetUserAuthRequest]) (*connect.Response[v1.SetUserAuthResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("user.v1.Base.SetUserAuth is not implemented"))
}

func (UnimplementedBaseHandler) DeleteUserAuth(context.Context, *connect.Request[v1.DeleteUserAuthRequest]) (*connect.Response[v1.DeleteUserAuthResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("user.v1.Base.DeleteUserAuth is not implemented"))
}
