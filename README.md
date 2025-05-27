# Connectwise PSA Go SDK
This is an unofficial Go SDK for Connectwise PSA. I am adding endpoints as I need it, therefore this doesn't have close to full API coverage.

## Example Usage
```
ctx := context.Background()

creds := connectwise.Creds{
    PublicKey:  "my_public_key",
    PrivateKey: "my_private_key",
    ClientId:   "my_client_id",
    CompanyId:  "my_company_id",
}

client := connectwise.NewClient(creds, http.DefaultClient)

params := &connectwise.QueryParams{
    OrderBy: "name asc",
    Conditions: "name='Help Desk'",
}

boards, err := client.ListBoards(ctx, params)
if err != nil {
    return nil, err
}
```