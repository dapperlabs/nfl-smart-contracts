import NonFungibleToken from "../../../contracts/NonFungibleToken.cdc"
import AllDay from "../../../contracts/AllDay.cdc"
import Escrow from "../../../contracts/AllDay.cdc"

transaction(leaderboardName: String, nftID: UInt64, nftType: String) {
    let nftVault: @NonFungibleToken.NFT
    let escrowRef: &Escrow.Admin
    let withdrawAddress: Address

    prepare(signer: AuthAccount, admin: AuthAccount) {
        // Borrow a reference to the user's NFT collection as a Provider
        let collectionRef = signer.borrow<&{NonFungibleToken.Provider}>(
            from: AllDay.CollectionStoragePath
        ) ?? panic("Could not borrow NFT collection reference")

        // Withdraw the NFT from the user's collection
        self.nftVault <- collectionRef.withdraw(withdrawID: nftID)

        // Borrow a reference to the Escrow Admin resource from the admin account
        self.escrowRef = admin.borrow<&Escrow.Admin>(from: Escrow.AdminStoragePath)
            ?? panic("Could not borrow Escrow Admin reference")

        self.withdrawAddress = signer.address
    }

    execute {
        // Get a reference to the desired leaderboard
        let leaderboard = self.escrowRef.getLeaderboard(name: leaderboardName)
            ?? panic("Leaderboard not found: ".concat(leaderboardName))

        // Add the NFT entry to the leaderboard
        leaderboard.addEntry(
            nft: <-self.nftVault,
            ownerAddress: self.withdrawAddress,
            leaderboardName: leaderboardName,
            typeName: nftType
        )
    }
}
