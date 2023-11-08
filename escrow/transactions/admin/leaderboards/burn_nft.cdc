import Escrow from "../../../contracts/AllDay.cdc"
import AllDay from "../../../contracts/AllDay.cdc"

// This transaction takes the leaderboardName and nftID and burns the NFT.
transaction(leaderboardName: String, nftID: UInt64) {
    prepare(signer: AuthAccount) {
        // Get a reference to the Admin resource in storage.
        let adminRef = signer.borrow<&Escrow.Admin>(from: Escrow.AdminStoragePath)
            ?? panic("Could not borrow reference to the Admin resource")

        // Call withdraw function.
        adminRef.burn(leaderboardName: leaderboardName, nftID: nftID)
    }

    execute {
        log("Burned NFT from leaderboard")
    }
}
