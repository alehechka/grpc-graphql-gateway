// Code generated by protoc-gen-graphql. DO NOT EDIT.
// versions:
// 	protoc-gen-graphql dev
// 	protoc             v3.21.4
// source: greeter.proto

package greeter

import (
	"context"

	"github.com/alehechka/grpc-graphql-gateway/runtime"
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

var (
	gql__type_HelloRequest    *graphql.Object      // message HelloRequest in greeter.proto
	gql__type_HelloReply      *graphql.Object      // message HelloReply in greeter.proto
	gql__type_GoodbyeRequest  *graphql.Object      // message GoodbyeRequest in greeter.proto
	gql__type_GoodbyeReply    *graphql.Object      // message GoodbyeReply in greeter.proto
	gql__input_HelloRequest   *graphql.InputObject // message HelloRequest in greeter.proto
	gql__input_HelloReply     *graphql.InputObject // message HelloReply in greeter.proto
	gql__input_GoodbyeRequest *graphql.InputObject // message GoodbyeRequest in greeter.proto
	gql__input_GoodbyeReply   *graphql.InputObject // message GoodbyeReply in greeter.proto
)

func Gql__type_HelloRequest() *graphql.Object {
	if gql__type_HelloRequest == nil {
		gql__type_HelloRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Greeter_Type_HelloRequest",
			Fields: graphql.Fields{
				"name": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.String),
					Description: `Below line means the "name" field is required in GraphQL argument`,
				},
			},
		})
	}
	return gql__type_HelloRequest
}

func Gql__type_HelloReply() *graphql.Object {
	if gql__type_HelloReply == nil {
		gql__type_HelloReply = graphql.NewObject(graphql.ObjectConfig{
			Name: "Greeter_Type_HelloReply",
			Fields: graphql.Fields{
				"message": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_HelloReply
}

func Gql__type_GoodbyeRequest() *graphql.Object {
	if gql__type_GoodbyeRequest == nil {
		gql__type_GoodbyeRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Greeter_Type_GoodbyeRequest",
			Fields: graphql.Fields{
				"name": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.String),
					Description: `Below line means the "name" field is required in GraphQL argument`,
				},
			},
		})
	}
	return gql__type_GoodbyeRequest
}

func Gql__type_GoodbyeReply() *graphql.Object {
	if gql__type_GoodbyeReply == nil {
		gql__type_GoodbyeReply = graphql.NewObject(graphql.ObjectConfig{
			Name: "Greeter_Type_GoodbyeReply",
			Fields: graphql.Fields{
				"message": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_GoodbyeReply
}

func Gql__input_HelloRequest() *graphql.InputObject {
	if gql__input_HelloRequest == nil {
		gql__input_HelloRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Greeter_Input_HelloRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"name": &graphql.InputObjectFieldConfig{
					Description: `Below line means the "name" field is required in GraphQL argument`,
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
		})
	}
	return gql__input_HelloRequest
}

func Gql__input_HelloReply() *graphql.InputObject {
	if gql__input_HelloReply == nil {
		gql__input_HelloReply = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Greeter_Input_HelloReply",
			Fields: graphql.InputObjectConfigFieldMap{
				"message": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_HelloReply
}

func Gql__input_GoodbyeRequest() *graphql.InputObject {
	if gql__input_GoodbyeRequest == nil {
		gql__input_GoodbyeRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Greeter_Input_GoodbyeRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"name": &graphql.InputObjectFieldConfig{
					Description: `Below line means the "name" field is required in GraphQL argument`,
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
		})
	}
	return gql__input_GoodbyeRequest
}

func Gql__input_GoodbyeReply() *graphql.InputObject {
	if gql__input_GoodbyeReply == nil {
		gql__input_GoodbyeReply = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Greeter_Input_GoodbyeReply",
			Fields: graphql.InputObjectConfigFieldMap{
				"message": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_GoodbyeReply
}

// graphql__resolver_Greeter is a struct for making query, mutation and resolve fields.
// This struct must be implemented runtime.SchemaBuilder interface.
type graphql__resolver_Greeter struct {

	// Automatic connection host
	host string

	// grpc dial options
	dialOptions []grpc.DialOption

	// grpc client connection.
	// this connection may be provided by user
	conn *grpc.ClientConn
}

// new_graphql_resolver_Greeter creates pointer of service struct
func new_graphql_resolver_Greeter(conn *grpc.ClientConn, host string, opts []grpc.DialOption) *graphql__resolver_Greeter {
	return &graphql__resolver_Greeter{
		conn:        conn,
		host:        host,
		dialOptions: opts,
	}
}

// CreateConnection() returns grpc connection which user specified or newly connected and closing function
func (x *graphql__resolver_Greeter) CreateConnection(ctx context.Context) (*grpc.ClientConn, func(), error) {
	// If x.conn is not nil, user injected their own connection
	if x.conn != nil {
		return x.conn, func() {}, nil
	}

	// Otherwise, this handler opens connection with specified host
	conn, err := grpc.DialContext(ctx, x.host, x.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	return conn, func() { conn.Close() }, nil
}

// GetQueries returns acceptable graphql.Fields for Query.
func (x *graphql__resolver_Greeter) GetQueries(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{
		"hello": &graphql.Field{
			Type: Gql__type_HelloReply(),
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type:         graphql.NewNonNull(graphql.String),
					Description:  `Below line means the "name" field is required in GraphQL argument`,
					DefaultValue: "",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req HelloRequest
				if err := runtime.MarshalRequest(p.Args, &req); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for hello")
				}
				client := NewGreeterClient(conn)
				resp, err := client.SayHello(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC SayHello")
				}
				return runtime.MarshalResponse(resp), nil
			},
		},
		"goodbye": &graphql.Field{
			Type: Gql__type_GoodbyeReply(),
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type:         graphql.NewNonNull(graphql.String),
					Description:  `Below line means the "name" field is required in GraphQL argument`,
					DefaultValue: "",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req GoodbyeRequest
				if err := runtime.MarshalRequest(p.Args, &req); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for goodbye")
				}
				client := NewGreeterClient(conn)
				resp, err := client.SayGoodbye(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC SayGoodbye")
				}
				return runtime.MarshalResponse(resp), nil
			},
		},
	}
}

// GetMutations returns acceptable graphql.Fields for Mutation.
func (x *graphql__resolver_Greeter) GetMutations(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{}
}

// Register package divided graphql handler "without" *grpc.ClientConn,
// therefore gRPC connection will be opened and closed automatically.
// Occasionally you may worry about open/close performance for each handling graphql request,
// then you can call RegisterGreeterGraphqlHandler with *grpc.ClientConn manually.
func RegisterGreeterGraphql(mux *runtime.ServeMux, host string, opts ...grpc.DialOption) error {
	return RegisterGreeterGraphqlHandler(mux, nil, host, opts...)
}

// Register package divided graphql handler "with" *grpc.ClientConn.
// this function accepts your defined grpc connection, so that we reuse that and never close connection inside.
// You need to close it manually when application will terminate.
func RegisterGreeterGraphqlHandler(mux *runtime.ServeMux, conn *grpc.ClientConn, host string, opts ...grpc.DialOption) error {
	return mux.AddHandler(new_graphql_resolver_Greeter(conn, host, opts))
}
