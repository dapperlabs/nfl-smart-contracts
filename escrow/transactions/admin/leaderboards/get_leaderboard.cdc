import Escrow from "../../../contracts/AllDay.cdc"

transaction(leaderboardName: String) {

    let adminRef: &Escrow.Admin{Escrow.IAdmin}

    prepare(signer: AuthAccount) {
        // Borrow a reference to the Admin resource in the signer's account
        self.adminRef = signer.borrow<&Escrow.Admin{Escrow.IAdmin}>(from: Escrow.AdminStoragePath)
            ?? panic("Could not borrow reference to the Admin resource")
    }

    execute {
        let leaderboard = self.adminRef.getLeaderboard(name: leaderboardName)

        // Check if the leaderboard exists and log an appropriate message
        if let board = leaderboard {
            log("Leaderboard found with name: ".concat(leaderboardName))
        } else {
            panic("Leaderboard not found with name: ".concat(leaderboardName))
        }
    }
}
