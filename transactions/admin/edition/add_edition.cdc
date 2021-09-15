import Genies from "../../../contracts/Genies.cdc"

transaction(
    name: String,
    collectionID: UInt32,
    metadata: {String: String}
   ) {
    // local variable for storing the collection reference
    let collection: &Genies.GeniesCollection

    prepare(signer: AuthAccount) {
        // borrow a reference to the Admin resource
        let admin = signer.borrow<&Genies.Admin>(from: Genies.AdminStoragePath)
            ?? panic("Could not borrow a reference to the Genies Admin capability")

        self.collection = admin.borrowGeniesCollection(id: collectionID)
    }

    execute {
        let id = self.collection.addEdition(
            editionName: name,
            editionMetadata: metadata
        )

        log("====================================")
        log("New Edition: ".concat(name))
        log("editionID")
        log(id)
        log("collectionID")
        log(collectionID)
        log("====================================")
    }
}