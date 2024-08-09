import NonFungibleToken from "NonFungibleToken"
import AllDay from "AllDay"
import MetadataViews from "MetadataViews"
import PackNFT from "PackNFT"

/// This transaction sets up the signer's account to hold AllDay NFTs and PackNFTs if it hasn't already been configured.
///
transaction {
    prepare(signer: auth(Storage, Capabilities) &Account) {
        // Return early if the account already has a collection
        if signer.storage.borrow<&AllDay.Collection>(from: AllDay.CollectionStoragePath) != nil {
            return
        }

        // Create a new collection and save it to the account storage
        signer.storage.save(<- AllDay.createEmptyCollection(nftType: Type<@AllDay.NFT>()), to: AllDay.CollectionStoragePath)

        // Create a public capability for the collection
        signer.capabilities.unpublish(AllDay.CollectionPublicPath)
        signer.capabilities.publish(
            signer.capabilities.storage.issue<&AllDay.Collection>(AllDay.CollectionStoragePath),
            at: AllDay.CollectionPublicPath
        )

        // Return early if the account already has a collection
        if signer.storage.borrow<&PackNFT.Collection>(from: PackNFT.CollectionStoragePath) != nil {
            return
        }

        // Create a new collection and save it to the account storage
        signer.storage.save(<- PackNFT.createEmptyCollection(nftType: Type<@PackNFT.NFT>()), to: PackNFT.CollectionStoragePath)

        // Create a public capability for the collection
        signer.capabilities.unpublish(PackNFT.CollectionPublicPath)
        signer.capabilities.publish(
            signer.capabilities.storage.issue<&PackNFT.Collection>(PackNFT.CollectionStoragePath),
            at: PackNFT.CollectionPublicPath
        )
    }
}