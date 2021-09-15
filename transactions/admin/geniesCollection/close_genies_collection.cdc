import Genies from "../../../contracts/Genies.cdc"

transaction(collectionID: UInt32) {
    // local variable for storing the series reference
    let series: &Genies.Series

    prepare(signer: AuthAccount) {
        // borrow a reference to the Admin resource
        let admin = signer.borrow<&Genies.Admin>(from: Genies.AdminStoragePath)
            ?? panic("Could not borrow a reference to the Genies Admin capability")

        let seriesID = admin.borrowGeniesCollection(id: collectionID).seriesID
        self.series = admin.borrowSeries(id: seriesID)
    }

    execute {
        let id = self.series.closeGeniesCollection(collectionID: collectionID)
    }
}

