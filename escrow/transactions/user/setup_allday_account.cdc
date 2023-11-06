import NonFungibleToken from "../../contracts/NonFungibleToken.cdc"
import AllDay from "../../contracts/AllDay.cdc"

// This transaction configures an account to hold AllDay NFTs.

transaction {
    prepare(signer: AuthAccount) {
        // if the account doesn't already have a collection
        if signer.borrow<&AllDay.Collection>(from: AllDay.CollectionStoragePath) == nil {

            // create a new empty collection
            let collection <- AllDay.createEmptyCollection()
            
            // save it to the account
            signer.save(<-collection, to: AllDay.CollectionStoragePath)

            // create a public capability for the collection
            signer.link<&AllDay.Collection{NonFungibleToken.CollectionPublic, AllDay.MomentNFTCollectionPublic}>(
                AllDay.CollectionPublicPath,
                target: AllDay.CollectionStoragePath
            )
        }
    }
}
