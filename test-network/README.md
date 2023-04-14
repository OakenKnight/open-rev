# Deploying the Hyperledger Fabric test network
Test network consists of 3 Organizations. Each organization has 3 joined peers.

## Running

startSystem.sh script will tear down already existing HyperledgerFabric network and restart the system, enroll peers and commit chaincode to the ledger. For purposes of development it will also transfer crypto materials to location of Proxy api.

```bash
./startSystem.sh
```

