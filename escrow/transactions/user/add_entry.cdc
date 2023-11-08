import NonFungibleToken from "../../../contracts/NonFungibleToken.cdc"
import AllDay from "../../../contracts/AllDay.cdc"
import Escrow from "../../../contracts/AllDay.cdc"

transaction(leaderboardName: String, nftID: UInt64) {
    let nft: @NonFungibleToken.NFT
    let escrowRef: &Escrow.Admin
    let receiver: Capability<&{NonFungibleToken.CollectionPublic}>

    prepare(signer: AuthAccount, admin: AuthAccount) {
        // Borrow a reference to the user's NFT collection as a Provider
        let collectionRef = signer.borrow<&{NonFungibleToken.Provider}>(
            from: AllDay.CollectionStoragePath
        ) ?? panic("Could not borrow NFT collection reference")

        // Borrow a reference to the user's NFT collection as a Receiver.
        self.receiver = signer.getCapability<&{NonFungibleToken.CollectionPublic}>(AllDay.CollectionPublicPath)!

        // Withdraw the NFT from the user's collection
        self.nft <- collectionRef.withdraw(withdrawID: nftID)

        // Borrow a reference to the Escrow Admin resource from the admin account
        self.escrowRef = admin.borrow<&Escrow.Admin>(from: Escrow.AdminStoragePath)
            ?? panic("Could not borrow Escrow Admin reference")
    }

    execute {
        // Get a reference to the desired leaderboard
        let leaderboard = self.escrowRef.getLeaderboard(name: leaderboardName)
            ?? panic("Leaderboard not found: ".concat(leaderboardName))

        // Add the NFT entry to the leaderboard
        leaderboard.addEntry(
            nft: <-self.nft,
            leaderboardName: leaderboardName,
            depositCap: self.receiver
        )
    }
}
