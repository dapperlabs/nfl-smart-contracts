import AllDay from "../../../contracts/AllDay.cdc"

transaction(seriesID: UInt32) {
    // local variable for the admin reference
    let admin: &AllDay.Admin

    prepare(signer: AuthAccount) {
        // borrow a reference to the Admin resource
        self.admin = signer.borrow<&AllDay.Admin>(from: AllDay.AdminStoragePath)
            ?? panic("Could not borrow a reference to the AllDay Admin capability")
    }

    execute {
        let id = self.admin.closeSeries(id: seriesID)

        log("====================================")
        log("Closed Series:")
        log("SeriesID: ".concat(id.toString()))
        log("====================================")
    }
}

