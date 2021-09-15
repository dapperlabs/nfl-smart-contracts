import NonFungibleToken from "../../contracts/NonFungibleToken.cdc"
import Genies from "../../contracts/Genies.cdc"

// This transaction configures an account to hold Genies NFTs.

transaction {
    prepare(signer: AuthAccount) {
        // if the account doesn't already have a collection
        if signer.borrow<&Genies.Collection>(from: Genies.CollectionStoragePath) == nil {

            // create a new empty collection
            let collection <- Genies.createEmptyCollection()
            
            // save it to the account
            signer.save(<-collection, to: Genies.CollectionStoragePath)

            // create a public capability for the collection
            signer.link<&Genies.Collection{NonFungibleToken.CollectionPublic, Genies.GeniesNFTCollectionPublic}>(
                Genies.CollectionPublicPath,
                target: Genies.CollectionStoragePath
            )
        }
    }
}
