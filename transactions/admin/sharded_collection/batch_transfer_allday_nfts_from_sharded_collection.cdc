import NonFungibleToken from "../../../contracts/NonFungibleToken.cdc"
import AllDay from "../../../contracts/AllDay.cdc"
import AllDayShardedCollection from "../../../contracts/AllDayShardedCollection.cdc"

// This transaction deposits a number of NFTs to a recipient

// Parameters
//
// recipient: the Flow address who will receive the NFTs
// momentNFTIDs: an array of momentNFT IDs of NFTs that recipient will receive

transaction(recipient: Address, momentNFTIDs: [UInt64]) {

    let transferTokens: @NonFungibleToken.Collection
    
    prepare(acct: AuthAccount) {
        
        self.transferTokens <- acct.borrow<&AllDayShardedCollection.ShardedCollection>(
            from: AllDayShardedCollection.CollectionStoragePath
            )!
            .batchWithdraw(ids: momentNFTIDs)
    }

    execute {

        // get the recipient's public account object
        let recipient = getAccount(recipient)

        // get the Collection reference for the receiver
        let receiverRef = recipient.getCapability(AllDay.CollectionPublicPath)
            .borrow<&{AllDay.MomentNFTCollectionPublic}>()!

        // deposit the NFT in the receivers collection
        receiverRef.batchDeposit(tokens: <-self.transferTokens)
    }
}

