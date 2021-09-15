import Genies from "../../../contracts/Genies.cdc"

transaction(editionID: UInt32) {
    // local variable for storing the collection reference
    let collection: &Genies.GeniesCollection

    prepare(signer: AuthAccount) {
        // borrow a reference to the Admin resource
        let admin = signer.borrow<&Genies.Admin>(from: Genies.AdminStoragePath)
            ?? panic("Could not borrow a reference to the Genies Admin capability")

        let collectionID = admin.borrowEdition(id: editionID).collectionID
        self.collection = admin.borrowGeniesCollection(id: collectionID)
    }

    execute {
        let id = self.collection.retireEdition(editionID: editionID)
    }
}

