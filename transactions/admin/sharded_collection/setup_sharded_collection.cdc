import NonFungibleToken from "../../../contracts/NonFungibleToken.cdc"
import AllDay from "../../../contracts/AllDay.cdc"
import AllDayShardedCollection from "../../../contracts/AllDayShardedCollection.cdc"

// This transaction creates and stores an empty NFT collection 
// and creates a public capability for it.
// NFTs are split into a number of buckets
// This makes storage more efficient and performant

// Parameters
//
// numBuckets: The number of buckets to split NFTs into

transaction(numBuckets: UInt64) {

    prepare(acct: AuthAccount) {

        if acct.borrow<&AllDayShardedCollection.ShardedCollection>(from: AllDayShardedCollection.CollectionStoragePath) == nil {

            let collection <- AllDayShardedCollection.createEmptyCollection(numBuckets: numBuckets)
            // Put a new Collection in storage
            acct.save(<-collection, to: AllDayShardedCollection.CollectionStoragePath)

            // create a public capability for the collection
            if acct.link<&{AllDay.MomentNFTCollectionPublic}>(AllDay.CollectionPublicPath, target: AllDayShardedCollection.CollectionStoragePath) == nil {
                acct.unlink(AllDay.CollectionPublicPath)
            }

            acct.link<&{AllDay.MomentNFTCollectionPublic}>(AllDay.CollectionPublicPath, target: AllDayShardedCollection.CollectionStoragePath)
        } else {
            panic("Sharded Collection already exists!")
        }
    }
}

