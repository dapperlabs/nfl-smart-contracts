import NonFungibleToken from "NonFungibleToken"
import AllDay from "AllDay"

// This script returns the size of an account's AllDay collection.

access(all) fun main(address: Address, id: UInt64): [AnyStruct] {
    let account = getAccount(address)

    let collectionRef = getAccount(address).capabilities.borrow<&AllDay.Collection>(AllDay.CollectionPublicPath)
        ?? panic("Could not borrow capability from public collection")
    
    let nft = collectionRef.borrowMomentNFT(id: id)
        ?? panic("Couldn't borrow momentNFT")

    return [nft.id, nft.editionID, nft.serialNumber, nft.mintingDate]
}

