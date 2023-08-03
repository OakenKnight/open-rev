CORE_PEER_LOCALMSPID=$1
CHANNEL=$2
peer channel update -o orderer:7050  -c ${CHANNEL} -f ./channel-artifacts/${CORE_PEER_LOCALMSPID}anchors.tx --tls  --cafile $ORDERER_CA 