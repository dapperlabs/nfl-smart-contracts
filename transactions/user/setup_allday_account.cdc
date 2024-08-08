import NonFungibleToken from "NonFungibleToken"
import AllDay from "AllDay"

// This transaction configures an account to hold AllDay NFTs.

transaction {
    prepare(signer: auth(Storage, Capabilities) &Account) {
        // if the account doesn't already have a collection
        if signer.storage.borrow<&AllDay.Collection>(from: AllDay.CollectionStoragePath) == nil {

            // create a new empty collection
            let collection <- AllDay.createEmptyCollection(nftType: Type<@AllDay.NFT>())
            
            // save it to the account
            signer.storage.save(<-collection, to: AllDay.CollectionStoragePath)

            // create a public capability for the collection
            signer.capabilities.unpublish(AllDay.CollectionPublicPath)
            signer.capabilities.publish(
                signer.capabilities.storage.issue<&AllDay.Collection>(AllDay.CollectionStoragePath),
                at: AllDay.CollectionPublicPath
            )
        }
    }
}
