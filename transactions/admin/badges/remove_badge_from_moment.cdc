import AllDay from "AllDay"

/// Removes a badge from a specific moment NFT
///
/// @param badgeSlug: The slug of the badge to remove from the moment
/// @param momentID: The ID of the moment NFT to remove the badge from
transaction(badgeSlug: String, momentID: UInt64) {
    
    // Local variable for the admin reference
    let admin: auth(AllDay.Operate) &AllDay.Admin
    
    prepare(signer: auth(BorrowValue) &Account) {
        // Get the admin resource
        self.admin = signer.storage.borrow<auth(AllDay.Operate) &AllDay.Admin>(from: AllDay.AdminStoragePath)
            ?? panic("Could not borrow admin resource")
    }
    
    execute {
        // Remove the badge from the moment
        self.admin.removeBadgeFromMoment(
            badgeSlug: badgeSlug,
            momentID: momentID
        )
    }
}
