# jdchain deploy scripts

## deploy
1. build jdchain, link to [jdchain](https://github.com/blockchain-jd-com/jdchain)
2. put targets, `jdchain-gateway-${release_version}.RELEASE.zip` and `jdchain-peer-${release_version}.RELEASE.zip`, under this directory
3. change permission for script `testnet.sh` and run
4. after that, it will build 4 peers and a gateway, and initialize ledger, and `start.sh` script for starting jdchain and `shutdown.sh` script for shuting down jdchain
5. change permission for scripts, `start.sh` and `shutdown.sh`
6. run script to start jdchain
