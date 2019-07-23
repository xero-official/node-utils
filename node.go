package main

import (
    "flag"
    "fmt"
    "log"
    "bufio"
    "os"
    "math"
    "math/big"
    "strings"
    "syscall"
    "os/signal"
    "context"

    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/common"
)

const logo = `
|
|____  __
|\   \/  /___________  ____   _____
| \     // __ \_  __ \/  _ \ /     \
| /     \  ___/|  | \(  <_> )  Y Y  \
|/___/\  \___  >__|   \____/|__|_|  /
|      \_/   \/                   \/
|             Node-Protocol Utilities
`

func main() {
    // Flags
    adminFlag := flag.Bool("admin", false, "a bool")
    flag.Parse()

    // Print logo
    fmt.Println(logo)
    // Setup graceful exit
    var gracefulStop = make(chan os.Signal)
    signal.Notify(gracefulStop, syscall.SIGTERM)
    signal.Notify(gracefulStop, syscall.SIGINT)
    go func() {
           <-gracefulStop
           fmt.Printf("\nExiting Program\n")
           os.Exit(0)
    }()

    var selectionFlag = false

    for selectionFlag != true {

        var contractOption int

        fmt.Println("1) Check Address Balance")
        fmt.Println("2) Exit")
        if *adminFlag {
            fmt.Println("3) Nothing Yet")
            fmt.Println("4) Nothing Yet")
        }

        _, _ = fmt.Scan(&contractOption)

        if contractOption == 1 {


            reader := bufio.NewReader(os.Stdin)
            // Get Address
            var address string
            fmt.Println("Enter Address:")
            address, _ = reader.ReadString('\n')
            address = strings.TrimSuffix(address, "\n")

            getBalance(address)

            selectionFlag = true


        } else if contractOption == 2 {

            selectionFlag = true
            fmt.Printf("\nExiting Program\n")
            os.Exit(0)


        } else {

            fmt.Println("\nInvalid Input\n")
            selectionFlag = false

        }

    }
}

func getBalance(address string) {

    client, err := ethclient.Dial(DefaultDataDir() + "/geth.ipc")
    if err != nil {
        log.Fatal(err)
    }

    checkAddress := common.HexToAddress(address)

    balance, err := client.BalanceAt(context.Background(), checkAddress, nil)
    if err != nil {
        log.Fatal(err)
    }

    fbalance := new(big.Float)
    fbalance.SetString(balance.String())
    ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

    fmt.Println("Checking Balance..")
    fmt.Println("Address Balance: " + ethValue.String())
    fmt.Println("\n")
}
/*
// Get user home directory from env
func getHomeDirectory() string {
    usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
    }
    return usr.HomeDir
}

// Retrieve nodekey and calculate enodeid
func getNodeId() []byte {
    b, err := ioutil.ReadFile(getHomeDirectory() + "/.xerom/geth/nodekey")
    if err != nil {
        fmt.Print(err)
        return []byte{}
    }
    enodeId, err := crypto.HexToECDSA(string(b))
    if err != nil {
        fmt.Print(err)
        return []byte{}
    }
    pubkeyBytes := crypto.FromECDSAPub(&enodeId.PublicKey)[1:]
    return pubkeyBytes
}
*/
