import NonFungibleToken from "../../contracts/NonFungibleToken.cdc"
import Showdown from "../../contracts/Showdown.cdc"

// Check to see if an account looks like it has been set up to hold Showdown NFTs.

pub fun main(address: Address): Bool {
    let account = getAccount(address)
    return account.getCapability<&{
            NonFungibleToken.CollectionPublic,
            Showdown.MomentNFTCollectionPublic
        }>(Showdown.CollectionPublicPath)
        != nil
}
