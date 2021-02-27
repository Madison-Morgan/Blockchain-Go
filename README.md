# Blockchain-Go
Building a blockchain in Go. I decided to do a few youtube tutorials and research into this concept... this a document of my findings.

## The Basics
A blockchain is a data structure that allows storing data in a decentralized manner. In blockchain there is an ordered chain of blocks such that each block contains the following information:
* Hash of the previous block
* List of transactions
* Hash of itself, which is the hash of list of transactions and hash of the previous block

In this way, no block is able to be changed, as it would lead to the hashes of all subsequent blocks becoming invalid, leading to a high level of data security. In other words, this data structure is immutable.

The first block is called the Genesis block 


## Applications of Blockchains
The most well known use of blockchain is in cryptocurrency such as bitcoin. There are also uses in Advertising , Supply Chain, Healthcare - effectively eliminating the middlemen involved in these respective transactions.
