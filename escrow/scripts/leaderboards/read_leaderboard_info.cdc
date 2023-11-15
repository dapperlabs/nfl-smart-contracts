import Escrow from "../../contracts/AllDay.cdc"

// This script returns the leaderboard info for the given leaderboard name.
pub fun main(leaderboardName: String): Escrow.LeaderboardInfo? {
    let account = getAccount("../../contracts/AllDay.cdc")

    let collectionPublic = account
        .getCapability<&Escrow.Collection{Escrow.ICollectionPublic}>(Escrow.CollectionPublicPath)
        .borrow()
        ?? panic("Could not borrow a reference to the public leaderboard collection")

    return collectionPublic.getLeaderboardInfo(name: leaderboardName)
}
