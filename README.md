<html>
<center><img src="./godefi.png"></img></center>
</br>
</html>

A golang sdk for working with DeFi protocols and general utilities for working with ethereum-compatible blockchains.

# packages

* cli
* config
* database
* sushiswap
* uniswap
* testenv
* utils

## cli

cli package

## config

configuration management package

## database

database management packlage

## sushiswap

Wrapper around go-ethereum's `ethclient` package for using sushiswap v2.

## uniswap

Wrapper around go-ethereum's `ethclient` package for using uniswap v2. 

## testenv

Provides a wrapper around the SimulatedBackend allowing for an in-memory blockchain. It is particularly useful for local smart contract development, or developing backend dApps.

## utils

Provides utility functions including all [goethereum book utils](https://goethereumbook.org/en/util-go/), a helper to make `bind.TransactOpts` safe for concurrent use, as well as a general `Blockchain` interface that satisfies all functions provided by simulated backend, as well as `ethclient` which is useful for drop-in replacement of actual RPC clients, and test environments.
