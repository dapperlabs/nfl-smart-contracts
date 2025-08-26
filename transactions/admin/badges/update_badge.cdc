import AllDay from "AllDay"

/// Updates an existing badge with new attributes
///
/// @param slug: The unique slug identifier of the badge to update
/// @param title: Optional new title for the badge
/// @param description: Optional new description for the badge
/// @param visible: Optional new visibility setting for the badge
/// @param slugV2: Optional new alternative slug identifier
/// @param metadata: Optional new metadata dictionary for the badge
transaction(
    slug: String,
    title: String?,
    description: String?,
    visible: Bool?,
    slugV2: String?,
    metadata: {String: String}?
) {
    
    // Local variable for the admin reference
    let admin: auth(AllDay.Operate) &AllDay.Admin
    
    prepare(signer: auth(BorrowValue) &Account) {
        // Get the admin resource
        self.admin = signer.storage.borrow<auth(AllDay.Operate) &AllDay.Admin>(from: AllDay.AdminStoragePath)
            ?? panic("Could not borrow admin resource")
    }
    
    execute {
        // Update the badge
        self.admin.updateBadge(
            slug: slug,
            title: title,
            description: description,
            visible: visible,
            slugV2: slugV2,
            metadata: metadata
        )
    }
}
