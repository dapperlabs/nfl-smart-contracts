import Showdown from "../../../contracts/Showdown.cdc"

transaction(editionID: UInt32) {
    // local variable for the admin reference
    let admin: &Showdown.Admin

    prepare(signer: AuthAccount) {
        // borrow a reference to the Admin resource
        self.admin = signer.borrow<&Showdown.Admin>(from: Showdown.AdminStoragePath)
            ?? panic("Could not borrow a reference to the Showdown Admin capability")
    }

    execute {
        let id = self.admin.closeEdition(id: editionID)

        log("====================================")
        log("Closed Edition:")
        log("EditionID: ".concat(id.toString()))
        log("====================================")
    }
}

