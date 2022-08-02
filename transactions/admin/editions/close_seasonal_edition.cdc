import AllDaySeasonal from "../../../contracts/AllDaySeasonal.cdc"

transaction(editionID: UInt64) {
    // local variable for the admin reference
    let admin: &AllDaySeasonal.Admin

    prepare(signer: AuthAccount) {
        // borrow a reference to the Admin resource
        self.admin = signer.borrow<&AllDaySeasonal.Admin>(from: AllDaySeasonal.AdminStoragePath)
            ?? panic("Could not borrow a reference to the AllDay Admin capability")
    }

    execute {
        let id = self.admin.closeEdition(id: editionID)

        log("====================================")
        log("Closed Edition:")
        log("EditionID: ".concat(id.toString()))
        log("====================================")
    }
}

