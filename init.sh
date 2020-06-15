#!/bin/bash
rm -r ~/.greenCLI
rm -r ~/.greenD

greenD init mynode --chain-id green-1001

greenCLI config keyring-backend test

greenCLI keys add me
greenCLI keys add you

greenD add-genesis-account $(greenCLI keys show me -a) 1000foo,100000000stake
greenD add-genesis-account $(greenCLI keys show you -a) 1foo

greenCLI config chain-id green-1001
greenCLI config output json
greenCLI config indent true
greenCLI config trust-node true

greenD gentx --name me --keyring-backend test
greenD collect-gentxs

