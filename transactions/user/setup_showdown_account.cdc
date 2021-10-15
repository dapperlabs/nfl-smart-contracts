import NonFungibleToken from "../../contracts/NonFungibleToken.cdc"
import Showdown from "../../contracts/Showdown.cdc"

// This transaction configures an account to hold Showdown NFTs.

transaction {
    prepare(signer: AuthAccount) {
        // if the account doesn't already have a collection
        if signer.borrow<&Showdown.Collection>(from: Showdown.CollectionStoragePath) == nil {

            // create a new empty collection
            let collection <- Showdown.createEmptyCollection()
            
            // save it to the account
            signer.save(<-collection, to: Showdown.CollectionStoragePath)

            // create a public capability for the collection
            signer.link<&Showdown.Collection{NonFungibleToken.CollectionPublic, Showdown.MomentNFTCollectionPublic}>(
                Showdown.CollectionPublicPath,
                target: Showdown.CollectionStoragePath
            )
        }
    }
}
