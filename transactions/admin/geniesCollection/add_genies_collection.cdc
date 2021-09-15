import Genies from "../../../contracts/Genies.cdc"

transaction(
    name: String,
    seriesID: UInt32,
    metadata: {String: String}
   ) {
    // local variable for storing the series reference
    let series: &Genies.Series

    prepare(signer: AuthAccount) {
        // borrow a reference to the Admin resource
        let admin = signer.borrow<&Genies.Admin>(from: Genies.AdminStoragePath)
            ?? panic("Could not borrow a reference to the Genies Admin capability")

        self.series = admin.borrowSeries(id: seriesID)
    }

    execute {
        let id = self.series.addCollection(
            collectionName: name,
            collectionMetadata: metadata
        )

        log("====================================")
        log("New collection: ".concat(name))
        log("collectionID")
        log(id)
        log("seriesID")
        log(seriesID)
        log("====================================")
    }
}
