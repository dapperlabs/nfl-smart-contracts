import NonFungibleToken from "NonFungibleToken"
import AllDay from "AllDay"
import MetadataViews from "MetadataViews"
import PackNFT from "PackNFT"

/// Check if an account has been set up to hold AllDay NFTs.
///
access(all) fun main(address: Address): Bool {
    let account = getAccount(address)
    return account.capabilities.borrow<
        &AllDay.Collection>(AllDay.CollectionPublicPath) != nil &&
        account.capabilities.borrow<
        &PackNFT.Collection>(PackNFT.CollectionPublicPath) != nil
}