import AllDay from "AllDay"

transaction(playID: UInt64, description: String) {
    // local variable for the admin reference
    let admin: auth(AllDay.Operate) &AllDay.Admin

    prepare(signer: auth(BorrowValue) &Account) {
        // borrow a reference to the Admin resource
        self.admin = signer.storage.borrow<auth(AllDay.Operate) &AllDay.Admin>(from: AllDay.AdminStoragePath)
            ?? panic("Could not borrow a reference to the AllDay Admin capability")
    }


    execute {
        let id = self.admin.updatePlayDescription(
            playID: playID,
            description: description
        )
    }

    post {
        AllDay.getPlayData(id: playID).metadata["description"] == description :
            "play description update failed"
    }
}
