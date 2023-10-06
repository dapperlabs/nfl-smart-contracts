import AllDay from "../../../contracts/AllDay.cdc"

transaction(
    seriesID: UInt64,
    setID: UInt64,
    playID: UInt64,
    tier: String,
    maxMintSize: UInt64?,
    deserialized: Bool?,
    presetSerial: UInt64?,
   ) {
    // local variable for the admin reference
    let admin: &AllDay.Admin

    prepare(signer: AuthAccount) {
        // borrow a reference to the Admin resource
        self.admin = signer.borrow<&AllDay.Admin>(from: AllDay.AdminStoragePath)
            ?? panic("Could not borrow a reference to the AllDay Admin capability")
    }

    execute {
        let id = self.admin.createEdition(
            seriesID: seriesID,
            setID: setID,
            playID: playID,
            maxMintSize: maxMintSize,
            tier: tier,
            deserialized: deserialized,
            presetSerial: presetSerial,
        )

        log("====================================")
        log("New Edition:")
        log("EditionID: ".concat(id.toString()))
        log("====================================")
    }
}

