# Design Document

## Artiqox Data/Information Marketplace
The Artiqox marketplace application is the central zone within the Artiqox ecosystem.
It is the decentralized data/information exchange to sell and buy structured information
provided by a range of *data providers*. A *data provider* is a seperat DApp/AppBC which
generates structured information.

### Artiqox Application Blockchain
In the following the Artiqox AppBC modules are discussed in detail.

1) (IPFS) storage/warehouse module  
The storage/warehouse module manages the structured information reports submitted by 
the specific application blockchains (ABC). Such ABCs generate structured information 
packages and store the into the Artiqox Cloud (IPFS?). The marketplace will then track
new data stored on the cloud and make it available for purchase to a paying customer.

2) Marketplace module  
The marketplace module provides the *API* to purchase structured information blobs.
It should accept different filter/ranges to select the blobs of interest. Since all 
information blobs will be indexed it should be possible to apply index filter for purchasing
specific information such as names, addresses, bandwith...

3) Governance  
The governance module manages the supply of **Artiqox Energy**

## Roadmap
This roadmap is the first approach to write down a simple development roadmap for building
the Artiqox ecosystem
1) Building the Artiqox Application Blockchain (aka Marketplace)  
  Develop the baseapp with a simple banking and auth module  
  Add the storage module
  Add the marketplace module
  Implement the wallet for customers  
  Deploy testnet
2) Build the Artiqox PegZone to transfer AIQ into the new ecosystem
  Develop the AIQc (c for cosmos) application blockchain  
  (split this in more subtasks once more experience is gained)
3) Implement the Notel ABC as the first data provider

## Module Description
**Storage** Module  
The storage module handles access to stored reports. It finds the right file(s) and checks for 
access permissions etc.
