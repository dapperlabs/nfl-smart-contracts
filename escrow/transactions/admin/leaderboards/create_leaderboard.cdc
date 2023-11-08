import Escrow from "../../../contracts/AllDay.cdc"
import AllDay from "../../../contracts/AllDay.cdc"

// This transaction takes a name and creates a new leaderboard with that name.
transaction(leaderboardName: String) {
    prepare(signer: AuthAccount) {
        // Get a reference to the Collection resource in storage.
        let collectionRef = signer.borrow<&Escrow.Collection>(from: Escrow.CollectionStoragePath)
            ?? panic("Could not borrow reference to the Collection resource")

        let type = Type<@AllDay.NFT>()

        // Create the leaderboard using the admin resource's method.
        collectionRef.createLeaderboard(name: leaderboardName, nftType: type)
    }
}
