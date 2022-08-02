# CHANGELOG

## UNRELEASED

## [v0.3.0](https://github.com/alehechka/grpc-graphql-gateway/releases/tag/v0.3.0)

- Resolve camelCase enum values ([#13](https://github.com/alehechka/grpc-graphql-gateway/pull/13))
- Remove field_camel and add field_proto CLI arg ([#14](https://github.com/alehechka/grpc-graphql-gateway/pull/14))
- Refactor getTagName function ([#16](https://github.com/alehechka/grpc-graphql-gateway/pull/16))
- GitHub Action Workflows - Publish tagged BSR repository ([#17](https://github.com/alehechka/grpc-graphql-gateway/pull/17))

## [v0.2.4](https://github.com/alehechka/grpc-graphql-gateway/releases/tag/v0.2.4)

- Add compiler and plugin version info to header comment ([#12](https://github.com/alehechka/grpc-graphql-gateway/pull/12))

## [v0.2.3](https://github.com/alehechka/grpc-graphql-gateway/releases/tag/v0.2.3)

- Resolve issues with goreleaser action [v0.2.3] ([#11](https://github.com/alehechka/grpc-graphql-gateway/pull/11))

## [v0.2.2](https://github.com/alehechka/grpc-graphql-gateway/releases/tag/v0.2.2)

- GitHub Action Workflows - Pull Request Verification ([#7](https://github.com/alehechka/grpc-graphql-gateway/pull/7))
- GitHub Action Workflows - Tagged Releases ([#9](https://github.com/alehechka/grpc-graphql-gateway/pull/9))

## [v0.2.1](https://github.com/alehechka/grpc-graphql-gateway/releases/tag/v0.2.1)

- Containerize plugin for remote code generation via buf.build ([#5](https://github.com/alehechka/grpc-graphql-gateway/pull/5))

## [v0.2.0](https://github.com/alehechka/grpc-graphql-gateway/releases/tag/v0.2.0)

- Resolve errant self imported package functions ([#1](https://github.com/alehechka/grpc-graphql-gateway/pull/1))
- Generate GraphQL types for empty proto messages ([#2](https://github.com/alehechka/grpc-graphql-gateway/pull/2))
- Allow Generation via buf.build and remote BSR proto dependency for graphql.proto by ([#3](https://github.com/alehechka/grpc-graphql-gateway/pull/3))
- Generate google/protobuf GraphQL types by ([#4](https://github.com/alehechka/grpc-graphql-gateway/pull/4))

## [v0.1.0](https://github.com/alehechka/grpc-graphql-gateway/releases/tag/v0.1.0)

- Resolve Import Issues ([#4](https://github.com/alehechka/forked-grpc-graphql-gateway/pull/4))
- Convert module path ([#3](https://github.com/alehechka/forked-grpc-graphql-gateway/pull/3))
- Refactor GraphqlHandler to expose host and DialOptions ([#1](https://github.com/alehechka/forked-grpc-graphql-gateway/pull/1))
- Resolve source_relative not respecting subpath proto files ([#2](https://github.com/alehechka/forked-grpc-graphql-gateway/pull/2))

# Original Changelog for [ysugimoto/grpc-graphql-gateway](https://github.com/ysugimoto/grpc-graphql-gateway)

## v0.22.0

- support parths=source_relative option [#51](https://github.com/ysugimoto/grpc-graphql-gateway/pull/51)

## v0.21.0

- Support Google's ptypes in Mutation response [#47](https://github.com/ysugimoto/grpc-graphql-gateway/pull/46)
- Fix misspelling in example [#46](https://github.com/ysugimoto/grpc-graphql-gateway/pull/46) (@day112)

## v0.20.1

- Support marshaling map type of Protocol Buffers [#38](https://github.com/ysugimoto/grpc-graphql-gateway/pull/38)

## v0.20.0

- Enable to specify Google's ptype in request message [#36](https://github.com/ysugimoto/grpc-graphql-gateway/pull/36)
- Support protobuf v3.14.0 [#36](https://github.com/ysugimoto/grpc-graphql-gateway/pull/36)

## v0.19.2

Bugifxes

- Fix input object when message has dependency of extermal package [#35](https://github.com/ysugimoto/grpc-graphql-gateway/pull/35)

## v0.19.1

Bugifxes

- Fix Input message type generation for nested resolver [#33](https://github.com/ysugimoto/grpc-graphql-gateway/pull/33)

## v0.19.0

Bugifxes

- Fix Input message type generation [#32](https://github.com/ysugimoto/grpc-graphql-gateway/pull/32)
- Fix Query argument field type [#30](https://github.com/ysugimoto/grpc-graphql-gateway/pull/30) (@tarunvelli)

## v0.18.1, v0.18.2

- Fix Google's ptypes package and generation [#27](https://github.com/ysugimoto/grpc-graphql-gateway/pull/27) [#29](https://github.com/ysugimoto/grpc-graphql-gateway/pull/29)

## v0.18.0

- Enable to parse enums within message [#26](https://github.com/ysugimoto/grpc-graphql-gateway/pull/26) (@hychrisli)

## v0.17.0

- Add default error handler which add `code` extension via gRPC error [#25](https://github.com/ysugimoto/grpc-graphql-gateway/pull/25)

## v0.16.0

## New graphql.proto option

- Add new graphql option of `resolver` [#24](https://github.com/ysugimoto/grpc-graphql-gateway/pull/24)

```protobuf
message GraphqlField {
    bool required = 1;
    string name = 2;
    string default = 3;
    bool omit = 4;
    string resolver = 5;
}
```

A new field of `resolve` which resolves as a nested query.

## v0.15.0

Add `omit` option in graphql.field option.

## v0.14.6

Bugfixes

- enum: use protobuf enum type for value [#18](https://github.com/ysugimoto/grpc-graphql-gateway/pull/18) (@bmkessler)

## v0.14.5

Bugfixes

- Implement request transformation from CamelCase to SnakeCase.

## v0.14.3, v0.14.4

Bugfixes

- Fix unexpected panic caused by reflect package in marshaling JSON response with camel-case

## v0.14.1, v0.14.2

Bugfixes

- Fix tiny syntax error
- Fix field camelcase generation
- Fix required sign in repeated array and input object

## v0.14.0

### Partially support google's type

Implement specific types and provide from this repository:

- google.protobuf.Timestamp
- google.protobuf.Wrappers
- google.protobuf.Empty

Note that we could only support the above types hence if a user imports other types e.g. google.protobuf.Any will raise an error.

And for GraphQL spec, google.protobuf.Empty defined empty field like `_: Boolean` because GraphQL raises error when type fields are empty.

### Fix map type in proto3

In proto3, map type can define as `map<string, string>` but PB parse as xxxEntry message. This PR also can use as graphql type with required key and value.

### Version printing in plugin binary

`protoc-gen-graphql` accepts `-v` option for print build version.

## v0.13.0

### Add infix typename

Add infix typename to GraphQL typename in order to avoid conflicting name between `Type` and `Input`.
After this version, GraphQL typename is modified that `[PackageName]_[Type]_[MessageName]` for example:

```
package user

message User {
  int64 user_id = 1;
}
```

Then typename is `User_Type_User`.

### Convert field name to CamelCase option

In protocol buffers, all message field name should define as _lower_snake_case_ referred by [Style Guide](https://developers.google.com/protocol-buffers/docs/style#message_and_field_names). But in GraphQL Schema, typically each field name should define as _lowerCamelCase_, so we add more option in `protoc-gen-graphql`:

```shell
protoc -I.
    --graphql_out=field_camel=true:.
    --go_out=plugins=grpc:.
    example.proto
```

The new option, `field_camel=true` converts all message field name to camel case like:

```
// protobuf
message User {
    int64 user_id = 1 [(graphql.field).required = true];
    string user_name = 2 [(graphql.field).required = true];
}

// Graphql Schema
type User_Type_User {
    userId Int!
    userName String!
}
```

To keep backward compatibility, compile as _lower_snake_case_ as default. If you want to define any graphql field as lowerCamelCase, please supply this option.

## v0.12.0

### Define MiddlewareError and respond with error code

We defined `MiddleWareError` which can return in Middleware function. If middleware function returns that pointer instead of common error,
The runtime responds error with that field data.

The definition is:

```go
type MiddlewareError struct {
  Code string
  Message string
}

// generate error
return runtime.NewMiddlewareError("CODE", "MESSAGE")
```

If middleware returns common error, the runtime error will be:

```
{"data": null, "errors": [{"message": "error message of err.Error()", "extensions": {"code": "MIDDLEWARE_ERROR"}]}
```

If middleware returns `MiddlewareError`, the runtime error will be:

```
{"data": null, "errors": [{"message": "Message field value of MiddlewareError", "extensions": {"code": "Code field value of MiddlewareError"}]}
```

## v0.9.1

### Changed middleware fucntion type

On MiddlewareFunc, you need to return `context.Context` as first return value. this is because we need to make custom metadata to gRPC on middleware process.
If you are already using your onw middleware, plase change its interface. see https://github.com/ysugimoto/grpc-graphql-gateway/pull/10
