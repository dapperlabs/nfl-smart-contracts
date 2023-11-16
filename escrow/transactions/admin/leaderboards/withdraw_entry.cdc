import Escrow from "../../../contracts/AllDay.cdc"
import AllDay from "../../../contracts/AllDay.cdc"

// This transaction takes the leaderboardName and nftID and returns it to the correct owner.
transaction(leaderboardName: String, nftID: UInt64) {
    prepare(signer: AuthAccount) {
        // Get a reference to the Collection resource in storage.
        let collectionRef = signer.borrow<&Escrow.Collection>(from: Escrow.CollectionStoragePath)
            ?? panic("Could not borrow reference to the Collection resource")

        // Call returnNftToCollection function.
        collectionRef.returnNftToCollection(leaderboardName: leaderboardName, nftID: nftID)
    }

    execute {
        log("Withdrawn NFT from leaderboard")
    }
}
