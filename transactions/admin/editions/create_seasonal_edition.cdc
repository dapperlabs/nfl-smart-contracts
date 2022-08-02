import AllDaySeasonal from "../../../contracts/AllDaySeasonal.cdc"

transaction(
    metadata: {String: String}
   ) {
    // local variable for the admin reference
    let admin: &AllDaySeasonal.Admin

    prepare(signer: AuthAccount) {
        // borrow a reference to the Admin resource
        self.admin = signer.borrow<&AllDaySeasonal.Admin>(from: AllDaySeasonal.AdminStoragePath)
            ?? panic("Could not borrow a reference to the AllDay Admin capability")
    }

    execute {
        let id = self.admin.createEdition(
            metadata: metadata
        )

        log("====================================")
        log("New Edition:")
        log("EdiionID: ".concat(id.toString()))
        log("====================================")
    }
}

