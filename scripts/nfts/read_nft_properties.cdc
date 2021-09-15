import NonFungibleToken from "../../contracts/NonFungibleToken.cdc"
import Genies from "../../contracts/Genies.cdc"

// This script returns the size of an account's Genies collection.

pub fun main(address: Address, id: UInt64): [AnyStruct] {
    let account = getAccount(address)

    let collectionRef = account.getCapability(Genies.CollectionPublicPath)
        .borrow<&{Genies.GeniesNFTCollectionPublic}>()
        ?? panic("Could not borrow capability from public collection")
    
    let nft = collectionRef.borrowGeniesNFT(id: id)
        ?? panic("Couldn't borrow geniesNFT")

    return [nft.id, nft.editionID, nft.serialNumber, nft.mintingDate]
}
