
#!/bin/bash
# script for purposes of starting network from scratch
RED='\033[0;31m'
NC='\033[0m' # No Color

export IS_DEVELOPMENT=true
echo $IS_DEVELOPMENT
echo -e "${RED}Tearing down HyperledgerFabric network"

./network.sh down

echo -e "${RED}Bringing network up and creating basic channel"

./network.sh up createChannel -ca

echo -e "${RED}Deploying smart contracts"
./network.sh deployCC -ccn basic -ccp ../open-rev-chaincode -ccl go	


# part for transfering crypto material
echo -e "${RED}Transfering cryptogen materials"
./transfer-materials.sh