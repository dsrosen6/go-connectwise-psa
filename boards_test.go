package connectwise

import (
	"context"
	"net/http"
	"os"
	"testing"
)

var (
	testClient = NewClient(testCreds, http.DefaultClient)
	testCtx    = context.Background()
	testCreds  = Creds{
		PublicKey:  os.Getenv("CW_PSA_PUBLIC_KEY"),
		PrivateKey: os.Getenv("CW_PSA_PRIVATE_KEY"),
		ClientId:   os.Getenv("CW_PSA_CLIENT_ID"),
		CompanyId:  os.Getenv("CW_PSA_COMPANY_ID"),
	}
)

func TestClient_GetBoards(t *testing.T) {
	params := map[string]string{"pagesize": "1", "conditions": "name='Projects'"}
	boards, err := testClient.GetBoards(testCtx, params)
	if err != nil {
		t.Fatal(err)
	}

	for _, board := range boards {
		t.Log(board.Name)
	}
}

func TestClient_GetBoard(t *testing.T) {
	board, err := testClient.GetBoard(testCtx, 1, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Received board:", board.Name)
}
