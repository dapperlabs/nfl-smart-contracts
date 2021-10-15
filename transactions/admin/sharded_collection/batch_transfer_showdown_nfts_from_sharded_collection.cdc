import NonFungibleToken from "../../../contracts/NonFungibleToken.cdc"
import Showdown from "../../../contracts/Showdown.cdc"
import ShowdownShardedCollection from "../../../contracts/ShowdownShardedCollection.cdc"

// This transaction deposits a number of NFTs to a recipient

// Parameters
//
// recipient: the Flow address who will receive the NFTs
// momentNFTIDs: an array of momentNFT IDs of NFTs that recipient will receive

transaction(recipient: Address, momentNFTIDs: [UInt64]) {

    let transferTokens: @NonFungibleToken.Collection
    
    prepare(acct: AuthAccount) {
        
        self.transferTokens <- acct.borrow<&ShowdownShardedCollection.ShardedCollection>(
            from: ShowdownShardedCollection.CollectionStoragePath
            )!
            .batchWithdraw(ids: momentNFTIDs)
    }

    execute {

        // get the recipient's public account object
        let recipient = getAccount(recipient)

        // get the Collection reference for the receiver
        let receiverRef = recipient.getCapability(Showdown.CollectionPublicPath)
            .borrow<&{Showdown.MomentNFTCollectionPublic}>()!

        // deposit the NFT in the receivers collection
        receiverRef.batchDeposit(tokens: <-self.transferTokens)
    }
}