# iceforg
some test or study about blockchain

```
yum install libtool-ltdl-devel
```
### use
deploy a network
```
    make deploy
    
```
clean environment
```
    make clean
    
```

### mark
[CK](https://www.jianshu.com/p/191d1e21f7ed)
```
    ./chaincode_example02
    peer chaincode install -p chaincodedev/chaincode/chaincode_example02 -n mycc -v 0
    peer chaincode instantiate -n mycc -v 0 -c '{"Args":["init","a","100","b","200"]}' -C myc
    peer chaincode invoke -n mycc -c '{"Args":["invoke","a","b","10"]}' -C myc
    query -n mycc -c '{"Args":["query","a"]}' -C myc
```

###contract
```
    dir: /opt/data/contracts
```