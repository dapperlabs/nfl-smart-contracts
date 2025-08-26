import AllDay from "AllDay"

/// Adds a badge to a specific play
///
/// @param badgeSlug: The slug of the badge to add to the play
/// @param playID: The ID of the play to add the badge to
/// @param metadata: Additional metadata for this badge-play association
transaction(badgeSlug: String, playID: UInt64, metadata: {String: String}) {
    
    // Local variable for the admin reference
    let admin: auth(AllDay.Operate) &AllDay.Admin
    
    prepare(signer: auth(BorrowValue) &Account) {
        // Get the admin resource
        self.admin = signer.storage.borrow<auth(AllDay.Operate) &AllDay.Admin>(from: AllDay.AdminStoragePath)
            ?? panic("Could not borrow admin resource")
    }
    
    execute {
        if AllDay.getBadge(badgeSlug) == nil{
                    panic("Badge with specified slug does not exist")
                }

        // Add the badge to the play
        self.admin.addBadgeToPlay(
            badgeSlug: badgeSlug,
            playID: playID,
            metadata: metadata
        )
    }
}
