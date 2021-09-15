import NonFungibleToken from "../../../contracts/NonFungibleToken.cdc"
import Genies from "../../../contracts/Genies.cdc"
import GeniesShardedCollection from "../../../contracts/GeniesShardedCollection.cdc"

// This transaction deposits a number of NFTs to a recipient

// Parameters
//
// recipient: the Flow address who will receive the NFTs
// geniesNFTIDs: an array of geniesNFT IDs of NFTs that recipient will receive

transaction(recipient: Address, geniesNFTIDs: [UInt64]) {

    let transferTokens: @NonFungibleToken.Collection
    
    prepare(acct: AuthAccount) {
        
        self.transferTokens <- acct.borrow<&GeniesShardedCollection.ShardedCollection>(
            from: GeniesShardedCollection.CollectionStoragePath
            )!
            .batchWithdraw(ids: geniesNFTIDs)
    }

    execute {

        // get the recipient's public account object
        let recipient = getAccount(recipient)

        // get the Collection reference for the receiver
        let receiverRef = recipient.getCapability(Genies.CollectionPublicPath)
            .borrow<&{Genies.GeniesNFTCollectionPublic}>()!

        // deposit the NFT in the receivers collection
        receiverRef.batchDeposit(tokens: <-self.transferTokens)
    }
}