import AllDay from "AllDay"

/// Removes a badge from a specific edition
///
/// @param badgeSlug: The slug of the badge to remove from the edition
/// @param editionID: The ID of the edition to remove the badge from
transaction(badgeSlug: String, editionID: UInt64) {
    
    // Local variable for the admin reference
    let admin: auth(AllDay.Operate) &AllDay.Admin
    
    prepare(signer: auth(BorrowValue) &Account) {
        // Get the admin resource
        self.admin = signer.storage.borrow<auth(AllDay.Operate) &AllDay.Admin>(from: AllDay.AdminStoragePath)
            ?? panic("Could not borrow admin resource")
    }
    
    execute {
        // Remove the badge from the edition
        self.admin.removeBadgeFromEdition(
            badgeSlug: badgeSlug,
            editionID: editionID
        )
    }
}
