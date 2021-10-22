import Showdown from "../../../contracts/Showdown.cdc"

transaction(
    name: String,
    metadata: {String: String}
   ) {
    // local variable for the admin reference
    let admin: &Showdown.Admin

    prepare(signer: AuthAccount) {
        // borrow a reference to the Admin resource
        self.admin = signer.borrow<&Showdown.Admin>(from: Showdown.AdminStoragePath)
            ?? panic("Could not borrow a reference to the Showdown Admin capability")
    }

    execute {
        let id = self.admin.createPlay(
            classification: name,
            metadata: metadata
        )

        log("====================================")
        log("New Play:")
        log("PlayID: ".concat(id.toString()))
        log("====================================")
    }
}
