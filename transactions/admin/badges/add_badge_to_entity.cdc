import AllDay from "AllDay"

/// Adds a badge to a specific entity (play, edition, or moment)
///
/// @param badgeSlug: The slug of the badge to add to the entity
/// @param entityType: The type of entity ("play", "edition", or "moment")
/// @param entityID: The ID of the entity to add the badge to
/// @param metadata: Additional metadata for this badge-entity association
transaction(badgeSlug: String, entityType: String, entityID: UInt64, metadata: {String: String}) {
    
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

        // Convert string to BadgeEntityType enum and add the badge
        switch entityType {
            case "play":
                self.admin.addBadgeToEntity(
                    badgeSlug: badgeSlug,
                    entityType: AllDay.BadgeEntityType.play,
                    entityID: entityID,
                    metadata: metadata
                )
            case "edition":
                self.admin.addBadgeToEntity(
                    badgeSlug: badgeSlug,
                    entityType: AllDay.BadgeEntityType.edition,
                    entityID: entityID,
                    metadata: metadata
                )
            case "moment":
                self.admin.addBadgeToEntity(
                    badgeSlug: badgeSlug,
                    entityType: AllDay.BadgeEntityType.moment,
                    entityID: entityID,
                    metadata: metadata
                )
            default:
                panic("Invalid entity type: ".concat(entityType))
        }
    }
}
