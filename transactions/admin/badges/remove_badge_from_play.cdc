import AllDay from "AllDay"

/// Removes a badge from a specific play
///
/// @param badgeSlug: The slug of the badge to remove from the play
/// @param playID: The ID of the play to remove the badge from
transaction(badgeSlug: String, playID: UInt64) {
    
    // Local variable for the admin reference
    let admin: auth(AllDay.Operate) &AllDay.Admin
    
    prepare(signer: auth(BorrowValue) &Account) {
        // Get the admin resource
        self.admin = signer.storage.borrow<auth(AllDay.Operate) &AllDay.Admin>(from: AllDay.AdminStoragePath)
            ?? panic("Could not borrow admin resource")
    }
    
    execute {
        // Remove the badge from the play
        self.admin.removeBadgeFromPlay(
            badgeSlug: badgeSlug,
            playID: playID
        )
    }
}
