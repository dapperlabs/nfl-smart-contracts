import AllDay from "AllDay"

transaction(playID: UInt64, optTeamName: String?, optPlayerFirstName: String?, optPlayerLastName: String?,
    optPlayerNumber: String?, optPlayerPosition: String?) {
    // local variable for the admin reference
    let admin: &AllDay.Admin

    prepare(signer: auth(BorrowValue) &Account) {
        // borrow a reference to the Admin resource
        self.admin = signer.storage.borrow<&AllDay.Admin>(from: AllDay.AdminStoragePath)
            ?? panic("Could not borrow a reference to the AllDay Admin capability")
    }


    execute {
        self.admin.updateDynamicMetadata(playID: playID, optTeamName: optTeamName,
            optPlayerFirstName: optPlayerFirstName, optPlayerLastName: optPlayerLastName,
            optPlayerNumber: optPlayerNumber, optPlayerPosition: optPlayerPosition)

        let play = AllDay.getPlayData(id: playID)
        if let teamName = optTeamName {
            assert(play.metadata["teamName"] == teamName, message: "team name update failed")
        }
        if let playerFirstName = optPlayerFirstName {
            assert(play.metadata["playerFirstName"] == playerFirstName, message: "player first name update failed")
        }
        if let playerLastName = optPlayerLastName {
            assert(play.metadata["playerLastName"] == playerLastName, message: "player last name update failed")
        }
        if let playerNumber = optPlayerNumber {
            assert(play.metadata["playerNumber"] == playerNumber, message: "player number update failed")
        }
        if let playerPosition = optPlayerPosition {
            assert(play.metadata["playerPosition"] == playerPosition, message: "player position update failed")
        }
    }
}
