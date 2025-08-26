import AllDay from "AllDay"

/// Adds a badge to a specific moment NFT
///
/// @param badgeSlug: The slug of the badge to add to the moment
/// @param momentID: The ID of the moment NFT to add the badge to
/// @param metadata: Additional metadata for this badge-moment association
transaction(badgeSlug: String, momentID: UInt64, metadata: {String: String}) {
    
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

        // Add the badge to the moment
        self.admin.addBadgeToMoment(
            badgeSlug: badgeSlug,
            momentID: momentID,
            metadata: metadata
        )
    }
}
