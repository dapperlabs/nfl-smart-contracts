import NonFungibleToken from "../../../contracts/NonFungibleToken.cdc"
import Genies from "../../../contracts/Genies.cdc"
import GeniesShardedCollection from "../../../contracts/GeniesShardedCollection.cdc"

// This transaction creates and stores an empty NFT collection 
// and creates a public capability for it.
// NFTs are split into a number of buckets
// This makes storage more efficient and performant

// Parameters
//
// numBuckets: The number of buckets to split NFTs into

transaction(numBuckets: UInt64) {

    prepare(acct: AuthAccount) {

        if acct.borrow<&GeniesShardedCollection.ShardedCollection>(from: GeniesShardedCollection.CollectionStoragePath) == nil {

            let collection <- GeniesShardedCollection.createEmptyCollection(numBuckets: numBuckets)
            // Put a new Collection in storage
            acct.save(<-collection, to: GeniesShardedCollection.CollectionStoragePath)

            // create a public capability for the collection
            if acct.link<&{Genies.GeniesNFTCollectionPublic}>(Genies.CollectionPublicPath, target: GeniesShardedCollection.CollectionStoragePath) == nil {
                acct.unlink(Genies.CollectionPublicPath)
            }

            acct.link<&{Genies.GeniesNFTCollectionPublic}>(Genies.CollectionPublicPath, target: GeniesShardedCollection.CollectionStoragePath)
        } else {
            panic("Sharded Collection already exists!")
        }
    }
}