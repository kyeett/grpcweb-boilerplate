package backend

import (
	"context"
	"fmt"

	"github.com/kyeett/grpcweb-boilerplate/proto/server"
)

// Backend should be used to implement the server interface
// exposed by the generated server proto.
type Backend struct {
}

// Ensure struct implements interface
var _ server.BackendServer = (*Backend)(nil)

func (b *Backend) NewPlayer(ctx context.Context, _ *server.Empty) (*server.PlayerID, error) {
	fmt.Println("Send send :-)")
	return &server.PlayerID{
		ID: "Magnus",
	}, nil
}
