import AllDay from "AllDay"

/// Removes a badge from a specific entity (play, edition, or moment)
///
/// @param badgeSlug: The slug of the badge to remove from the entity
/// @param entityType: The type of entity ("play", "edition", or "moment")
/// @param entityID: The ID of the entity to remove the badge from
transaction(badgeSlug: String, entityType: String, entityID: UInt64) {
    
    // Local variable for the admin reference
    let admin: auth(AllDay.Operate) &AllDay.Admin
    
    prepare(signer: auth(BorrowValue) &Account) {
        // Get the admin resource
        self.admin = signer.storage.borrow<auth(AllDay.Operate) &AllDay.Admin>(from: AllDay.AdminStoragePath)
            ?? panic("Could not borrow admin resource")
    }
    
    execute {
        // Convert string to BadgeEntityType enum and remove the badge
        switch entityType {
            case "play":
                self.admin.removeBadgeFromEntity(
                    badgeSlug: badgeSlug,
                    entityType: AllDay.BadgeEntityType.play,
                    entityID: entityID
                )
            case "edition":
                self.admin.removeBadgeFromEntity(
                    badgeSlug: badgeSlug,
                    entityType: AllDay.BadgeEntityType.edition,
                    entityID: entityID
                )
            case "moment":
                self.admin.removeBadgeFromEntity(
                    badgeSlug: badgeSlug,
                    entityType: AllDay.BadgeEntityType.moment,
                    entityID: entityID
                )
            default:
                panic("Invalid entity type: ".concat(entityType))
        }
    }
}
