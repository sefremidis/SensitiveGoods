#!/bin/bash

echo
CHANNEL_NAME=mychannel
DELAY=3
LANGUAGE=golang
VERSION=9.0
TIMEOUT=10
VERBOSE=false
COUNTER=1
MAX_RETRY=10

CC_SRC_PATH="github.com/chaincode/myGo/"
echo "Channel name : "$CHANNEL_NAME

# import utils
. scripts/utils.sh

createChannel() {
	setGlobals 0 1

	if [ -z "$CORE_PEER_TLS_ENABLED" -o "$CORE_PEER_TLS_ENABLED" = "false" ]; then
                set -x
		peer channel create -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx >&log.txt
		res=$?
                set +x
	else
				set -x
		peer channel create -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA >&log.txt
		res=$?
				set +x
	fi
	cat log.txt
	verifyResult $res "Channel creation failed"
	echo "===================== Channel '$CHANNEL_NAME' created ===================== "
	echo
}

joinChannel () {
	for org in 1 2; do
	    for peer in 0 1; do
		joinChannelWithRetry $peer $org
		echo "===================== peer${peer}.org${org} joined channel '$CHANNEL_NAME' ===================== "
		sleep $DELAY
		echo
	    done
	done

	joinChannelWithRetry 0 3
	echo "===================== peer0.org3 joined channel '$CHANNEL_NAME' ===================== "
	sleep $DELAY
	echo

	joinChannelWithRetry 0 4
	echo "===================== peer0.org4 joined channel '$CHANNEL_NAME' ===================== "
	sleep $DELAY
	echo
}

## Create channel
#echo "Creating channel..."
#createChannel
## Join all the peers to the channel
#echo "Having all peers join the channel..."
#joinChannel
## Set the anchor peers for each org in the channel
#echo "Updating anchor peers for org1..."
#updateAnchorPeers 0 1
#echo "Updating anchor peers for org2..."
#updateAnchorPeers 0 2
#updateAnchorPeers 0 3
#updateAnchorPeers 0 4
## Install chaincode on peer0.org1 and peer0.org2
#echo "Installing chaincode on peer0.org1..."
#installChaincode 0 1
#installChaincode 1 1
#echo "Install chaincode on peer0.org2..."
#installChaincode 0 2
#installChaincode 1 2
#installChaincode 0 3
#installChaincode 0 4
# Instantiate chaincode on peer0.org2
#echo "Instantiating chaincode on peer0.org2..."
#instantiateChaincode 0 1
#instantiateChaincode 0 2

#upgradeChaincode 0 1
#upgradeChaincode 0 2
# Invoke chaincode on peer0.org1 and peer0.org2
#echo "Sending invoke transaction on peer0.org1 peer0.org2..."

#chaincodeInvoke 0 1 0 2

# Query on chaincode on peer1.org2, check if the result is 90
#echo "Querying chaincode on peer1.org2..."
chaincodeQuery 0 1

echo
exit 0