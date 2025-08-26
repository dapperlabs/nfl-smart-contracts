import AllDay from "AllDay"

/// Gets a badge by its slug identifier
///
/// @param slug: The unique slug identifier of the badge to retrieve
/// @return: The badge data or nil if not found
access(all) fun main(slug: String): AllDay.Badge? {
    return AllDay.getBadge(slug)
}
