import Escrow from "../../../contracts/AllDay.cdc"
import AllDay from "../../../contracts/AllDay.cdc"

// This transaction takes a name and creates a new leaderboard with that name.
transaction(leaderboardName: String) {
    prepare(signer: AuthAccount) {
        // Get a reference to the Admin resource in storage.
        let adminRef = signer.borrow<&Escrow.Admin>(from: Escrow.AdminStoragePath)
            ?? panic("Could not borrow reference to the Admin resource")

        let type = Type<@AllDay.NFT>()

        // Create the leaderboard using the admin resource's method.
        adminRef.createLeaderboard(name: leaderboardName, nftType: type)
    }
}
