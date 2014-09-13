package main

import (
    "github.com/dshearer/jobber"
    "net"
    "net/rpc"
    "fmt"
    "os"
)

func main() {
    // connect to daemon
    addr, err := net.ResolveUnixAddr("unix", jobber.DaemonSocketAddr)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Couldn't resolve Unix addr: %v\n", err)
        os.Exit(1)
    }
    conn, err := net.DialUnix("unix", nil, addr)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Couldn't connect to daemon: %v\n", err)
        os.Exit(1)
    }
    defer conn.Close()
    rpcClient := rpc.NewClient(conn)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Couldn't make RPC client: %v\n", err)
        os.Exit(1)
    }
    
    // call
    var result string
    
    err = rpcClient.Call("RealIpcServer.ListJobs", 1, &result)
    if err != nil {
        fmt.Fprintf(os.Stderr, "RPC failed: %v\n", err)
        os.Exit(1)
    }
    fmt.Printf("%s\n", result)
    
    err = rpcClient.Call("RealIpcServer.ListHistory", 1, &result)
    if err != nil {
        fmt.Fprintf(os.Stderr, "RPC failed: %v\n", err)
        os.Exit(1)
    }
    fmt.Printf("%s\n", result)
    
/*    err = rpcClient.Call("RealIpcServer.Stop", 1, &result)
    if err != nil {
        fmt.Fprintf(os.Stderr, "RPC failed: %v\n", err)
        os.Exit(1)
    }*/
}
