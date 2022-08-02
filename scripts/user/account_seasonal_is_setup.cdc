import NonFungibleToken from "../../contracts/NonFungibleToken.cdc"
import AllDaySeasonal from "../../contracts/AllDaySeasonal.cdc"

// Check to see if an account looks like it has been set up to hold AllDay NFTs.

pub fun main(address: Address): Bool {
    let account = getAccount(address)
    return account.getCapability<&{
            NonFungibleToken.CollectionPublic,
            AllDaySeasonal.SeasonalNFTCollectionPublic
        }>(AllDaySeasonal.CollectionPublicPath)
        != nil
}

