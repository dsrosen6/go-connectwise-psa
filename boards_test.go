package connectwise

import (
	"context"
	"net/http"
	"os"
	"testing"
)

var (
	testCreds = Creds{
		PublicKey:  os.Getenv("CW_PSA_PUBLIC_KEY"),
		PrivateKey: os.Getenv("CW_PSA_PRIVATE_KEY"),
		ClientId:   os.Getenv("CW_PSA_CLIENT_ID"),
		CompanyId:  os.Getenv("CW_PSA_COMPANY_ID"),
	}
	testHttpClient      = http.DefaultClient
	testClient          = NewClient(testCreds, testHttpClient)
	testCtx             = context.Background()
	testGoodBoardParams = &QueryParams{OrderBy: "name asc"}
	testBadBoardParams  = &QueryParams{Conditions: "something=idk"}
)

func TestClient_ListBoards(t *testing.T) {
	type args struct {
		ctx    context.Context
		params *QueryParams
	}
	tests := []struct {
		name    string
		args    args
		want    []Board
		wantErr bool
	}{
		{
			name: "list boards no params", args: args{ctx: testCtx, params: nil}, want: nil, wantErr: false,
		},
		{
			name: "list boards with good params", args: args{ctx: testCtx, params: testGoodBoardParams}, want: nil, wantErr: false,
		},
		{
			name: "list boards with bad params", args: args{ctx: testCtx, params: testBadBoardParams}, want: nil, wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := testClient.ListBoards(tt.args.ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListBoards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
