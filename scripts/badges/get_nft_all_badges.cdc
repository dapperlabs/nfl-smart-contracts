import AllDay from "AllDay"
import NonFungibleToken from "NonFungibleToken"

/// Gets all badges associated with a specific NFT including those inherited from its play and edition
///
/// @param account: The account address that owns the NFT
/// @param nftID: The ID of the NFT to get badges for
/// @return: An array of all badges associated with the NFT (moment + edition + play badges)
access(all) fun main(accountAddress: Address, nftID: UInt64): [AllDay.Badge]? {
    
    // Get the account's public collection
    let account = getAccount(accountAddress)
    let collectionRef = account.capabilities.borrow<&{NonFungibleToken.CollectionPublic}>(AllDay.CollectionPublicPath)
        ?? panic("Could not borrow collection public reference")
    
    // Borrow the specific NFT
    let nft = collectionRef.borrowNFT(nftID)
        ?? panic("Could not borrow NFT")
    
    // Cast to AllDay NFT to access getBadges function
    let momentNFT = nft as! &AllDay.NFT
    
    // Get all badges for this NFT (includes moment, edition, and play badges)
    return momentNFT.getBadges()
}
