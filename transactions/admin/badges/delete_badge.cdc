import AllDay from "AllDay"

/// This transaction deletes a badge completely from the system
/// This will remove the badge from all associated plays, editions, and moments
///
/// Parameters:
/// - slug: The unique slug identifier of the badge to delete
///
transaction(slug: String) {

    let admin: auth(AllDay.Operate) &AllDay.Admin

    prepare(signer: auth(BorrowValue) &Account) {
        // borrow a reference to the Admin resource in storage
        self.admin = signer.storage.borrow<auth(AllDay.Operate) &AllDay.Admin>(from: AllDay.AdminStoragePath)
                    ?? panic("Could not borrow admin resource")
    }

    execute {
        // Delete the badge
        self.admin.deleteBadge(slug: slug)
    }
}
