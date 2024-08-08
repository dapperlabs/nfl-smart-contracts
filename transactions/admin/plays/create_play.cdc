import AllDay from "AllDay"

transaction(
    name: String,
    metadata: {String: String}
   ) {
    // local variable for the admin reference
    let admin: auth(AllDay.Operate) &AllDay.Admin

    prepare(signer: auth(BorrowValue) &Account) {
        // borrow a reference to the Admin resource
        self.admin = signer.storage.borrow<auth(AllDay.Operate) &AllDay.Admin>(from: AllDay.AdminStoragePath)
            ?? panic("Could not borrow a reference to the AllDay Admin capability")
    }

    execute {
        let id = self.admin.createPlay(
            classification: name,
            metadata: metadata
        )

        log("====================================")
        log("New Play:")
        log("PlayID: ".concat(id.toString()))
        log("====================================")
    }
}

