import AllDay from "AllDay"

/// Creates a new badge with the specified attributes
///
/// @param slug: The unique slug identifier for the badge
/// @param title: The display title of the badge
/// @param description: A description of what the badge represents
/// @param visible: Whether the badge should be visible to users
/// @param slugV2: An alternative slug identifier
transaction(slug: String, title: String, description: String, visible: Bool, slugV2: String) {
    
    // Local variable for the admin reference
    let admin: auth(AllDay.Operate) &AllDay.Admin
    
    prepare(signer: auth(BorrowValue) &Account) {
        // Get the admin resource
        self.admin = signer.storage.borrow<auth(AllDay.Operate) &AllDay.Admin>(from: AllDay.AdminStoragePath)
            ?? panic("Could not borrow admin resource")
    }
    
    execute {
        // Create the badge
        self.admin.createBadge(
            slug: slug,
            title: title,
            description: description,
            visible: visible,
            slugV2: slugV2
        )
    }
}
