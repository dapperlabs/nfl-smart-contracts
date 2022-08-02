import NonFungibleToken from "../../contracts/NonFungibleToken.cdc"
import AllDaySeasonal from "../../contracts/AllDaySeasonal.cdc"

// This script returns the size of an account's AllDay collection.

pub fun main(address: Address, id: UInt64): [AnyStruct] {
    let account = getAccount(address)

    let collectionRef = account.getCapability(AllDaySeasonal.CollectionPublicPath)
        .borrow<&{AllDaySeasonal.SeasonalNFTCollectionPublic}>()
        ?? panic("Could not borrow capability from public collection")
    
    let nft = collectionRef.borrowSeasonalNFT(id: id)
        ?? panic("Couldn't borrow momentNFT")

    return [nft.id, nft.editionID]
}

