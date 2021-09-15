import NonFungibleToken from "../../contracts/NonFungibleToken.cdc"
import Genies from "../../contracts/Genies.cdc"

// This script returns an array of all the NFT IDs in an account's collection.

pub fun main(address: Address): [UInt64] {
    let account = getAccount(address)

    let collectionRef = account.getCapability(Genies.CollectionPublicPath).borrow<&{NonFungibleToken.CollectionPublic}>()
        ?? panic("Could not borrow capability from public collection")
    
    return collectionRef.getIDs()
}
