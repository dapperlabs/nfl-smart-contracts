/*
    Author: Jude Zhu jude.zhu@dapperlabs.com
*/


import NonFungibleToken from "./NonFungibleToken.cdc"

/*
    There are 5 levels of entity:
    1. Editions
    2. Seasonal NFTs
    
    An Edition is created with metadata. Seasonal NFTs are minted out of Editions.
 */

// The AllDay NFTs and metadata contract
//
pub contract AllDaySeasonal: NonFungibleToken {
    //------------------------------------------------------------
    // Events
    //------------------------------------------------------------

    // Contract Events
    //
    pub event ContractInitialized()

    // NFT Collection Events
    //
    pub event Withdraw(id: UInt64, from: Address?)
    pub event Deposit(id: UInt64, to: Address?)
    pub event Burned(id: UInt64)
    pub event Minted(id: UInt64, editionID: UInt64)

    // Edition Events
    //
    // Emitted when a new edition has been created by an admin
    pub event EditionCreated(
        id: UInt64, 
        metadata: {String: String}
    )
    // Emitted when an edition is either closed by an admin, or the max amount of moments have been minted
    pub event EditionClosed(id: UInt64)

    //------------------------------------------------------------
    // Named values
    //------------------------------------------------------------

    // Named Paths
    //
    pub let CollectionStoragePath:  StoragePath
    pub let CollectionPublicPath:   PublicPath
    pub let AdminStoragePath:       StoragePath
    pub let MinterPrivatePath:      PrivatePath

    //------------------------------------------------------------
    // Publicly readable contract state
    //------------------------------------------------------------

    // totalSupply
    // The total number of {{ contractName }} NFTs that have been minted.
    //
    pub var totalSupply:        UInt64

    // totalEditions
    // The total number of {{ contractName }} editions that have been created.
    //
    pub var totalEditions: UInt64

    //------------------------------------------------------------
    // Internal contract state
    //------------------------------------------------------------

    // Metadata Dictionaries
    //
    // This is so we can find Series by their names (via seriesByID)
    access(self) let editionByID:       @{UInt64: Edition}

    //------------------------------------------------------------
    // Edition
    //------------------------------------------------------------

    // A public struct to access Edition data
    //
    pub struct EditionData {
        pub let id: UInt64
        pub let metadata: {String: String}

        // initializer
        //
        init (id: UInt64) {
            if let play = &AllDay.playByID[id] as &AllDay.Play? {
            self.id = id
            self.metadata = play.metadata
            } else {
                panic("play does not exist")
            }
        }
    }

    // A top level Edition with a unique ID
    //
    pub resource Edition {
        pub let id: UInt64
        // Contents writable if borrowed!
        // This is deliberate, as it allows admins to update the data.
        pub let metadata: {String: String}
        pub var active: Bool

        // Close this series
        //
        pub fun close() {
            pre {
                self.active == true: "not active"
            }

            self.active = false

            emit EditionClosed(id: self.id)
        }

        // Mint a Seasonal NFT in this edition, with the given minting mintingDate.
        // Note that this will panic if the max mint size has already been reached.
        //
        pub fun mint(): @AllDaySeasonal.NFT {
            pre {
                self.active: "edition closed, cannot mint"
            }

            // Create the Moment NFT, filled out with our information
            let momentNFT <- create NFT(
                id: AllDaySeasonal.totalSupply + 1,
                editionID: self.id,
                serialNumber: self.numMinted + 1
            )
            AllDaySeasonal.totalSupply = AllDaySeasonal.totalSupply + 1
            // Keep a running total (you'll notice we used this as the serial number)
            self.numMinted = self.numMinted + 1 as UInt64

            return <- momentNFT
        }

        // initializer
        //
        init (metadata: {String: String}) {
            self.id = AllDaySeasonal.nextEditionID
            self.metadata = metadata

            AllDaySeasonal.nextEditionID = self.id + 1 as UInt64
            emit EditionCreated(id: self.id, metadata: self.metadata)
        }
    }

    // Get the publicly available data for a Edition
    //
    pub fun getEditionData(id: UInt64): AllDaySeasonal.EditionData {
        pre {
            AllDaySeasonal.editionByID[id] != nil: "Cannot borrow edition, no such id"
        }

        return AllDaySeasonal.EditionData(id: id)
    }

    //------------------------------------------------------------
    // NFT
    //------------------------------------------------------------

    // A Moment NFT
    //
    pub resource NFT: NonFungibleToken.INFT {
        pub let id: UInt64
        pub let editionID: UInt64
        pub let serialNumber: UInt64
        pub let mintingDate: UFix64

        // Destructor
        //
        destroy() {
            emit MomentNFTBurned(id: self.id)
        }

        // NFT initializer
        //
        init(
            id: UInt64,
            editionID: UInt64,
            serialNumber: UInt64
        ) {
            pre {
                AllDaySeasonal.editionByID[editionID] != nil: "no such editionID"
                EditionData(id: editionID).maxEditionMintSizeReached() != true: "max edition size already reached"
            }

            self.id = id
            self.editionID = editionID
            self.serialNumber = serialNumber
            self.mintingDate = getCurrentBlock().timestamp

            emit MomentNFTMinted(id: self.id, editionID: self.editionID, serialNumber: self.serialNumber)
        }
    }

    //------------------------------------------------------------
    // Collection
    //------------------------------------------------------------

    // A public collection interface that allows Moment NFTs to be borrowed
    //
    pub resource interface MomentNFTCollectionPublic {
        pub fun deposit(token: @NonFungibleToken.NFT)
        pub fun batchDeposit(tokens: @NonFungibleToken.Collection)
        pub fun getIDs(): [UInt64]
        pub fun borrowNFT(id: UInt64): &NonFungibleToken.NFT
        pub fun borrowMomentNFT(id: UInt64): &AllDaySeasonal.NFT? {
            // If the result isn't nil, the id of the returned reference
            // should be the same as the argument to the function
            post {
                (result == nil) || (result?.id == id): 
                    "Cannot borrow Moment NFT reference: The ID of the returned reference is incorrect"
            }
        }
    }

    // An NFT Collection
    //
    pub resource Collection:
        NonFungibleToken.Provider,
        NonFungibleToken.Receiver,
        NonFungibleToken.CollectionPublic,
        MomentNFTCollectionPublic
    {
        // dictionary of NFT conforming tokens
        // NFT is a resource type with an UInt64 ID field
        //
        pub var ownedNFTs: @{UInt64: NonFungibleToken.NFT}

        // withdraw removes an NFT from the collection and moves it to the caller
        //
        pub fun withdraw(withdrawID: UInt64): @NonFungibleToken.NFT {
            let token <- self.ownedNFTs.remove(key: withdrawID) ?? panic("missing NFT")

            emit Withdraw(id: token.id, from: self.owner?.address)

            return <-token
        }

        // deposit takes a NFT and adds it to the collections dictionary
        // and adds the ID to the id array
        //
        pub fun deposit(token: @NonFungibleToken.NFT) {
            let token <- token as! @AllDaySeasonal.NFT
            let id: UInt64 = token.id

            // add the new token to the dictionary which removes the old one
            let oldToken <- self.ownedNFTs[id] <- token

            emit Deposit(id: id, to: self.owner?.address)

            destroy oldToken
        }

        // batchDeposit takes a Collection object as an argument
        // and deposits each contained NFT into this Collection
        //
        pub fun batchDeposit(tokens: @NonFungibleToken.Collection) {
            // Get an array of the IDs to be deposited
            let keys = tokens.getIDs()

            // Iterate through the keys in the collection and deposit each one
            for key in keys {
                self.deposit(token: <-tokens.withdraw(withdrawID: key))
            }

            // Destroy the empty Collection
            destroy tokens
        }

        // getIDs returns an array of the IDs that are in the collection
        //
        pub fun getIDs(): [UInt64] {
            return self.ownedNFTs.keys
        }

        // borrowNFT gets a reference to an NFT in the collection
        //
        pub fun borrowNFT(id: UInt64): &NonFungibleToken.NFT {
            pre {
                self.ownedNFTs[id] != nil: "Cannot borrow NFT, no such id"
            }

            return (&self.ownedNFTs[id] as &NonFungibleToken.NFT?)!
        }

        // borrowMomentNFT gets a reference to an NFT in the collection
        //
        pub fun borrowMomentNFT(id: UInt64): &AllDaySeasonal.NFT? {
            if self.ownedNFTs[id] != nil {
                if let ref = &self.ownedNFTs[id] as auth &NonFungibleToken.NFT? {
                    return ref! as! &AllDaySeasonal.NFT
                }
                return nil
            } else {
                return nil
            }
        }

        // Collection destructor
        //
        destroy() {
            destroy self.ownedNFTs
        }

        // Collection initializer
        //
        init() {
            self.ownedNFTs <- {}
        }
    }

    // public function that anyone can call to create a new empty collection
    //
    pub fun createEmptyCollection(): @NonFungibleToken.Collection {
        return <- create Collection()
    }

    //------------------------------------------------------------
    // Admin
    //------------------------------------------------------------

    // An interface containing the Admin function that allows minting NFTs
    //
    pub resource interface NFTMinter {
        // Mint a single NFT
        // The Edition for the given ID must already exist
        //
        pub fun mintNFT(editionID: UInt64): @AllDaySeasonal.NFT
    }

    // A resource that allows managing metadata and minting NFTs
    //
    pub resource Admin: NFTMinter {

        // Borrow an Edition
        //
        pub fun borrowEdition(id: UInt64): &AllDaySeasonal.Edition {
            pre {
                AllDaySeasonal.editionByID[id] != nil: "Cannot borrow edition, no such id"
            }

            return (&AllDaySeasonal.editionByID[id] as &AllDaySeasonal.Edition?)!
        }

        // Create a Edition 
        //
        pub fun createEdition(metadata: {String: String}): UInt64 {
            // Create and store the new edition
            let edition <- create AllDaySeasonal.Edition(
                metadata: metadata,
            )
            let editionID = edition.id
            AllDaySeasonal.editionByID[edition.id] <-! edition

            // Return the new ID for convenience
            return editionID
        }


        // Close an Edition
        //
        pub fun closeEdition(id: UInt64): UInt64 {
            if let edition = &AllDaySeasonal.editionByID[id] as &AllDaySeasonal.Edition? {
                edition.close()
                return edition.id
            }
            panic("edition does not exist")
        }

        // Mint a single NFT
        // The Edition for the given ID must already exist
        //
        pub fun mintNFT(editionID: UInt64): @AllDaySeasonal.NFT {
            pre {
                // Make sure the edition we are creating this NFT in exists
                AllDaySeasonal.editionByID.containsKey(editionID): "No such EditionID"
            }
            return <- self.borrowEdition(id: editionID).mint()
        }
    }

    //------------------------------------------------------------
    // Contract lifecycle
    //------------------------------------------------------------

    // AllDay contract initializer
    //
    init() {
        // Set the named paths
        self.CollectionStoragePath = /storage/AllDayNFTCollection
        self.CollectionPublicPath = /public/AllDayNFTCollection
        self.AdminStoragePath = /storage/AllDayAdmin
        self.MinterPrivatePath = /private/AllDayMinter

        // Initialize the entity counts        
        self.totalSupply = 0
        self.totalEditions = 0
        self.nextEditionID = 1

        // Initialize the metadata lookup dictionaries
        self.editionByID <- {}

        // Create an Admin resource and save it to storage
        let admin <- create Admin()
        self.account.save(<-admin, to: self.AdminStoragePath)
        // Link capabilites to the admin constrained to the Minter
        // and Metadata interfaces
        self.account.link<&AllDaySeasonal.Admin{AllDaySeasonal.NFTMinter}>(
            self.MinterPrivatePath,
            target: self.AdminStoragePath
        )

        // Let the world know we are here
        emit ContractInitialized()
    }
}
