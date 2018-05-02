package nodes



//var runFile = `
//cd /
//
//curl -X POST -H 'Content-Type: application/json' -d "%dropletname%:About to clone - %repoBranch% %repoURL%" %callback%/api/node/v1/log
//
//git clone -b %repoBranch% %repoURL%
//
//curl -X POST -H 'Content-Type: application/json' -d "%dropletname%:Clone complete - %repoBranch% %repoURL%" %callback%/api/node/v1/log
//
//
//cd /navcoin-core
//
//curl -X POST -H 'Content-Type: application/json' -d "%dropletname%:cd to navcoin-core" %callback%/api/node/v1/log
//
//
//./autogen.sh
//
//curl -X POST -H 'Content-Type: application/json' -d "%dropletname%:Ran autogen" %callback%/api/node/v1/log
//
//curl -X POST -H 'Content-Type: application/json' -d "%dropletname%:Configure navcoin " %callback%/api/node/v1/log
//./configure LDFLAGS="-L/usr/local/berkeley-db-4.8/lib/" \
//                CPPFLAGS="-I /usr/local/berkeley-db-4.8/include/" \
//                --enable-hardening \
//                --without-gui \
//                --enable-upnp-default
//
//curl -X POST -H 'Content-Type: application/json' -d "%dropletname%:Seed done" %callback%/api/node/v1/log
//`




var runFile = `
cd /

curl -X POST -H 'Content-Type: application/json' -d "%dropletname%:About to clone - %repoBranch% %repoURL%" %callback%/api/node/v1/log

git clone -b %repoBranch% %repoURL%

curl -X POST -H 'Content-Type: application/json' -d "%dropletname%:Clone complete - %repoBranch% %repoURL%" %callback%/api/node/v1/log


cd /navcoin-core

curl -X POST -H 'Content-Type: application/json' -d "%dropletname%:cd to navcoin-core" %callback%/api/node/v1/log


./autogen.sh

curl -X POST -H 'Content-Type: application/json' -d "%dropletname%:Ran autogen" %callback%/api/node/v1/log

curl -X POST -H 'Content-Type: application/json' -d "%dropletname%:Configure navcoin " %callback%/api/node/v1/log
./configure LDFLAGS="-L/usr/local/berkeley-db-4.8/lib/" \
                CPPFLAGS="-I /usr/local/berkeley-db-4.8/include/" \
                --enable-hardening \
                --without-gui \
                --enable-upnp-default

curl -X POST -H 'Content-Type: application/json' -d "%dropletname%:Making navcoin" %callback%/api/node/v1/log

ls
make

curl -X POST -H 'Content-Type: application/json' -d "%dropletname%:Make Complete" %callback%/api/node/v1/log
make install
curl -X POST -H 'Content-Type: application/json' -d "%dropletname%:Install Complete" %callback%/api/node/v1/log

cd ..

rm -fr /navcoin-core/*
rm -r /navcoin-core/

curl -X POST -H 'Content-Type: application/json' -d "%dropletname%:rm navcoin-src Complete" %callback%/api/node/v1/log


navcoind -devnet -rpcuser=hi -rpcpassword=pass &
curl -X POST -H 'Content-Type: application/json' -d "%dropletname%:Start navcoin core Complete" %callback%/api/node/v1/log

sleep 10
navcoin-cli -devnet -rpcuser=hi -rpcpassword=pass addnode "159.65.46.57" "add"

curl -X POST -H 'Content-Type: application/json' -d "%dropletname%:add nodes complete" %callback%/api/node/v1/log


OUTPUT="$(navcoin-cli -devnet -staking -rpcuser=hi -rpcpassword=pass listreceivedbyaddress 0 true)"
echo "${OUTPUT}"

curl -X POST -H 'Content-Type: application/json' -d "%dropletname%::${OUTPUT}" %callback%/api/node/v1/log/address

curl -X POST -H 'Content-Type: application/json' -d "%dropletname%:Finished" %callback%/api/node/v1/log


sleep 9999999999`
