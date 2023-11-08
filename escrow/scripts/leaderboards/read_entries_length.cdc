import Escrow from "../../contracts/AllDay.cdc"

// This script returns the number of entries in a specific leaderboard.
pub fun main(leaderboardName: String, address: Address): Int {
    let account = getAccount(address)

    let adminPublic = account
        .getCapability<&Escrow.Admin{Escrow.IAdmin}>(Escrow.AdminPublicPath)
        .borrow()
        ?? panic("Could not borrow reference to the Admin resource")

    // Use the Admin public reference to get the leaderboard
    let leaderboard = adminPublic.getLeaderboard(name: leaderboardName)
        ?? panic("Leaderboard not found")

    // Call the getEntriesLength function on the leaderboard to get the number of entries
    return leaderboard.getEntriesLength()
}
