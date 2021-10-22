import Showdown from "../../../contracts/Showdown.cdc"

transaction(name: String) {
    // local variable for the admin reference
    let admin: &Showdown.Admin

    prepare(signer: AuthAccount) {
        // borrow a reference to the Admin resource
        self.admin = signer.borrow<&Showdown.Admin>(from: Showdown.AdminStoragePath)
            ?? panic("Could not borrow a reference to the Showdown Admin capability")
    }

    execute {
        let id = self.admin.createSeries(
            name: name,
        )

        log("====================================")
        log("New Series: ".concat(name))
        log("SeriesID: ".concat(id.toString()))
        log("====================================")
    }
}

