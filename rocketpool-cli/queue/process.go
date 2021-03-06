package queue

import (
    "fmt"

    "github.com/urfave/cli"

    "github.com/rocket-pool/smartnode/shared/services/rocketpool"
)


func processQueue(c *cli.Context) error {

    // Get RP client
    rp, err := rocketpool.NewClientFromCtx(c)
    if err != nil { return err }
    defer rp.Close()

    // Check deposit queue can be processed
    canProcess, err := rp.CanProcessQueue()
    if err != nil {
        return err
    }
    if !canProcess.CanProcess {
        fmt.Println("The deposit queue cannot be processed:")
        if canProcess.AssignDepositsDisabled {
            fmt.Println("Deposit assignments are currently disabled.")
        }
        if canProcess.NoMinipoolsAvailable {
            fmt.Println("No minipools are available for assignment.")
        }
        if canProcess.InsufficientDepositBalance {
            fmt.Println("The deposit pool has an insufficient balance for assignment.")
        }
        return nil
    }

    // Process deposit queue
    if _, err := rp.ProcessQueue(); err != nil {
        return err
    }

    // Log & return
    fmt.Println("The deposit queue was successfully processed.")
    return nil

}

