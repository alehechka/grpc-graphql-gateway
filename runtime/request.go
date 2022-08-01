package runtime

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/runtime/protoiface"
)

type GraphqlRequest struct {
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
	OperationName string                 `json:"operationName"`
}

// ParseRequest parses graphql query and variables from each request methods
func parseRequest(r *http.Request) (*GraphqlRequest, error) {
	var body []byte

	// Get request body
	switch r.Method {
	case http.MethodPost:
		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, errors.New("malformed request body, " + err.Error())
		}
		body = buf
	case http.MethodGet:
		body = []byte(r.URL.Query().Get("query"))
	default:
		return nil, errors.New("invalid request method: '" + r.Method + "'")
	}

	// And try to parse
	var req GraphqlRequest
	if err := json.Unmarshal(body, &req); err != nil {
		// If error, the request body may come with single query line
		req.Query = string(body)
	}
	return &req, nil
}

// MarshalRequest marshals graphql request arguments to gRPC request message
func MarshalRequest(args map[string]interface{}, message protoiface.MessageV1) error {
	buf := new(bytes.Buffer)

	if err := json.NewEncoder(buf).Encode(args); err != nil {
		return err
	}

	return jsonpb.Unmarshal(buf, message)
}
