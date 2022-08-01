import NonFungibleToken from "../../contracts/NonFungibleToken.cdc"
import AllDaySeasonal from "../../contracts/AllDaySeasonal.cdc"

// This transaction configures an account to hold AllDay NFTs.

transaction {
    prepare(signer: AuthAccount) {
        // if the account doesn't already have a collection
        if signer.borrow<&AllDaySeasonal.Collection>(from: AllDaySeasonal.CollectionStoragePath) == nil {

            // create a new empty collection
            let collection <- AllDaySeasonal.createEmptyCollection()
            
            // save it to the account
            signer.save(<-collection, to: AllDaySeasonal.CollectionStoragePath)

            // create a public capability for the collection
            signer.link<&AllDaySeasonal.Collection{NonFungibleToken.CollectionPublic, AllDaySeasonal.MomentNFTCollectionPublic}>(
                AllDaySeasonal.CollectionPublicPath,
                target: AllDaySeasonal.CollectionStoragePath
            )
        }
    }
}
