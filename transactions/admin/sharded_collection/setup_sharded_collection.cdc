import NonFungibleToken from "../../../contracts/NonFungibleToken.cdc"
import Showdown from "../../../contracts/Showdown.cdc"
import ShowdownShardedCollection from "../../../contracts/ShowdownShardedCollection.cdc"

// This transaction creates and stores an empty NFT collection 
// and creates a public capability for it.
// NFTs are split into a number of buckets
// This makes storage more efficient and performant

// Parameters
//
// numBuckets: The number of buckets to split NFTs into

transaction(numBuckets: UInt64) {

    prepare(acct: AuthAccount) {

        if acct.borrow<&ShowdownShardedCollection.ShardedCollection>(from: ShowdownShardedCollection.CollectionStoragePath) == nil {

            let collection <- ShowdownShardedCollection.createEmptyCollection(numBuckets: numBuckets)
            // Put a new Collection in storage
            acct.save(<-collection, to: ShowdownShardedCollection.CollectionStoragePath)

            // create a public capability for the collection
            if acct.link<&{Showdown.MomentNFTCollectionPublic}>(Showdown.CollectionPublicPath, target: ShowdownShardedCollection.CollectionStoragePath) == nil {
                acct.unlink(Showdown.CollectionPublicPath)
            }

            acct.link<&{Showdown.MomentNFTCollectionPublic}>(Showdown.CollectionPublicPath, target: ShowdownShardedCollection.CollectionStoragePath)
        } else {
            panic("Sharded Collection already exists!")
        }
    }
}