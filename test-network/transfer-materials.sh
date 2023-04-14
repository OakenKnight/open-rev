sudo rm -rf ../open-rev-gateway-app/hyperledger-crypto-materials/*
cp -R ./organizations/peerOrganizations/* ../open-rev-gateway-app/hyperledger-crypto-materials/
chown -R $USER ../open-rev-gateway-app/hyperledger-crypto-materials
