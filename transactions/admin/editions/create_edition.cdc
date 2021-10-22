import Showdown from "../../../contracts/Showdown.cdc"

transaction(
    seriesID: UInt32,
    setID: UInt32,
    playID: UInt32,
    tier: String,
    maxMintSize: UInt32?,
   ) {
    // local variable for the admin reference
    let admin: &Showdown.Admin

    prepare(signer: AuthAccount) {
        // borrow a reference to the Admin resource
        self.admin = signer.borrow<&Showdown.Admin>(from: Showdown.AdminStoragePath)
            ?? panic("Could not borrow a reference to the Showdown Admin capability")
    }

    execute {
        let id = self.admin.createEdition(
            seriesID: seriesID,
            setID: setID,
            playID: playID,
            maxMintSize: maxMintSize,
            tier: tier,
        )

        log("====================================")
        log("New Edition:")
        log("EditionID: ".concat(id.toString()))
        log("====================================")
    }
}