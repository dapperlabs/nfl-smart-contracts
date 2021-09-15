import NonFungibleToken from "../../contracts/NonFungibleToken.cdc"
import Genies from "../../contracts/Genies.cdc"

// Check to see if an account looks like it has been set up to hold Genies NFTs.

pub fun main(address: Address): Bool {
    let account = getAccount(address)
    return account.getCapability<&{
            NonFungibleToken.CollectionPublic,
            Genies.GeniesNFTCollectionPublic
        }>(Genies.CollectionPublicPath)
        != nil
}
