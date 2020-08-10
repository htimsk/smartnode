package account

import (
    "errors"

    "github.com/urfave/cli"

    "github.com/rocket-pool/smartnode/shared/services"
    "github.com/rocket-pool/smartnode/shared/types/api"
)


func initPassword(c *cli.Context, password string) (*api.InitPasswordResponse, error) {

    // Get services
    pm, err := services.GetPasswordManager(c)
    if err != nil { return nil, err }

    // Response
    response := api.InitPasswordResponse{}

    // Check if password is already set
    if pm.IsPasswordSet() {
        return nil, errors.New("The node password is already set")
    }

    // Set password
    if err := pm.SetPassword(password); err != nil {
        return nil, err
    }

    // Return response
    return &response, nil

}

