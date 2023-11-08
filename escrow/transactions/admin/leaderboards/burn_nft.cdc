import Escrow from "../../../contracts/AllDay.cdc"
import AllDay from "../../../contracts/AllDay.cdc"

// This transaction takes the leaderboardName and nftID and burns the NFT.
transaction(leaderboardName: String, nftID: UInt64) {
    prepare(signer: AuthAccount) {
        // Get a reference to the Collection resource in storage.
        let collectionRef = signer.borrow<&Escrow.Collection>(from: Escrow.CollectionStoragePath)
            ?? panic("Could not borrow reference to the Collection resource")

        // Call withdraw function.
        collectionRef.burn(leaderboardName: leaderboardName, nftID: nftID)
    }

    execute {
        log("Burned NFT from leaderboard")
    }
}
