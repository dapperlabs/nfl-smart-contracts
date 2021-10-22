import NonFungibleToken from "../../../contracts/NonFungibleToken.cdc"
import AllDay from "../../../contracts/AllDay.cdc"
import AllDayShardedCollection from "../../../contracts/AllDayShardedCollection.cdc"

// This transaction deposits an NFT to a recipient

// Parameters
//
// recipient: the Flow address who will receive the NFT
// momentID: moment ID of NFT that recipient will receive

transaction(recipient: Address, momentID: UInt64) {

    let transferToken: @NonFungibleToken.NFT
    
    prepare(acct: AuthAccount) {

        self.transferToken <- acct.borrow<&AllDayShardedCollection.ShardedCollection>(
            from: AllDayShardedCollection.CollectionStoragePath
            )!
            .withdraw(withdrawID: momentID)
    }

    execute {
        
        // get the recipient's public account object
        let recipient = getAccount(recipient)

        // get the Collection reference for the receiver
        let receiverRef = recipient.getCapability(AllDay.CollectionPublicPath)
            .borrow<&{AllDay.MomentNFTCollectionPublic}>()!

        // deposit the NFT in the receivers collection
        receiverRef.deposit(token: <-self.transferToken)
    }
}

