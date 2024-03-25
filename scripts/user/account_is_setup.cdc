import NonFungibleToken from "NonFungibleToken"
import AllDay from "AllDay"

// Check to see if an account looks like it has been set up to hold AllDay NFTs.

access(all) fun main(address: Address): Bool {
    return getAccount(address).capabilities.borrow<
        &AllDay.Collection>(AllDay.CollectionPublicPath) != nil
}

