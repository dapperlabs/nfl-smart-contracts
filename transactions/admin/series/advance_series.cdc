import Genies from "../../../contracts/Genies.cdc"

// Close down all previous editions and collections are closed,
// then create and advance to new Series.

transaction(
    name: String,
    metadata: {String: String}
   ) {
    // local variable for the admin reference
    let admin: &Genies.Admin

    prepare(signer: AuthAccount) {
        // borrow a reference to the Admin resource
        self.admin = signer.borrow<&Genies.Admin>(from: Genies.AdminStoragePath)
            ?? panic("Could not borrow a reference to the Genies Admin capability")
    }

    execute {
        // Now we know everything in the previous series is closed,
        // we can advance the series.
        // This makes sure that the previous series is deactivated.
        let id = self.admin.advanceSeries(
            nextSeriesName: name,
            nextSeriesMetadata: metadata
        )

        log("====================================")
        log("New Series: ".concat(name))
        log("seriesID")
        log(id)
        log("====================================")
    }
}
