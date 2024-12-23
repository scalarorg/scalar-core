package client

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govrest "github.com/cosmos/cosmos-sdk/x/gov/client/rest"

	"github.com/scalarorg/scalar-core/x/scalarnet/client/cli"
)

// ProposalHandler is the call contracts proposal handler.
var ProposalHandler = govclient.NewProposalHandler(
	cli.NewSubmitCallContractsProposalTxCmd,
	func(ctx client.Context) govrest.ProposalRESTHandler {
		return govrest.ProposalRESTHandler{
			SubRoute: "unsupported-scalarnet-client",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				rest.WriteErrorResponse(w, http.StatusBadRequest, "Legacy REST Routes are not supported for scalarnet proposals")
			},
		}
	},
)
