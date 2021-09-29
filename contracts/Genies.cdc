import NonFungibleToken from "./NonFungibleToken.cdc"

/*
    Showdown is structured similarly to Genies and TopShot.
    Unlike TopShot, we use resources for all entities and manage access to their data
    by copying it to structs (this simplifies access control, in particular write access).
    We also encapsulate resource creation for the admin in member functions on the parent type.
    
    There are 5 levels of entity:
    1. Series
    2. Sets
    3. Plays
    4. Editions
    4. Moment NFT (an NFT)
    
    An Edition is created with a combination of a Series, Set, and Play
    Moment NFTs are minted out of Editions.

    Note that we cache some information (Series names/ids, counts of closed entities) rather
    than calculate it each time.
    This is enabled by encapsulation and saves gas for entity lifecycle operations.
 */

// The Showdown NFTs and metadata contract
//
pub contract Showdown: NonFungibleToken {
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

    // Series Events
    //
    // Emitted when a new series has been created by an admin
    pub event SeriesCreated(id: UInt32, name: String, metadata: {String: String})
    // Emitted when a series is closed by an admin
    pub event SeriesClosed(id: UInt32)

    // Set Events
    //
    // Emitted when a new set has been created by an admin
    pub event SetCreated(id: UInt32, name: String, metadata: {String: String})

    // Play Events
    //
    // Emitted when a new play has been created by an admin
    pub event PlayCreated(id: UInt32, classification: String, metadata: {String: String})

    // Edition Events
    //
    // Emitted when a new edition has been created by an admin
    pub event EditionCreated(
        id: UInt32, 
        seriesID: UInt32, 
        setID: UInt32, 
        playID: UInt32, 
        maxMintSize: UInt32?,
        tier: String, 
        metadata: {String: String},
    )
    // Emitted when an edition is either closed by an admin, or the max amount of moments have been minted
    pub event EditionClosed(id: UInt32)

    // NFT Events
    //
    pub event NFTMinted(id: UInt64, editionID: UInt32, serialNumber: UInt32)
    pub event NFTBurned(id: UInt64)

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

    // Entity Counts
    //
    pub var totalSupply:        UInt64
    pub var nextSeriesID:       UInt32
    pub var nextSetID:          UInt32
    pub var nextPlayID:         UInt32
    pub var nextEditionID:      UInt32

    //------------------------------------------------------------
    // Internal contract state
    //------------------------------------------------------------

    // Metadata Dictionaries
    //
    // This is so we can find Series by their names (via seriesByID)
    access(self) let seriesIDByName:    {String: UInt32}
    access(self) let seriesByID:        @{UInt32: Series}
    access(self) let setByID:           @{UInt32: Set}
    access(self) let playByID:          @{UInt32: Play}
    access(self) let editionByID:       @{UInt32: Edition}

    //------------------------------------------------------------
    // Series
    //------------------------------------------------------------

    // A public struct to access Series data
    //
    pub struct SeriesData {
        pub let id: UInt32
        pub let name: String
        pub let metadata: {String: String}
        pub let active: Bool

        // initializer
        //
        init (id: UInt32) {
            let series = &Showdown.seriesByID[id] as! &Showdown.Series
            self.id = series.id
            self.name = series.name
            self.metadata = series.metadata
            self.active = series.active
        }
    }

    // A top-level Series with a unique ID and name
    //
    pub resource Series {
        pub let id: UInt32
        pub let name: String
        // Contents writable if borrowed!
        // This is deliberate, as it allows admins to update the data.
        pub let metadata: {String: String}
        // We manage this list, but need to access it to fill out the struct,
        // so it is access(contract)
        pub var active: Bool

        // Close this series
        //
        pub fun close() {
            pre {
                self.active == true: "not active"
            }

            self.active = false

            emit SeriesClosed(id: self.id)
        }

        // initializer
        // We pass in ID as the logic for it is more complex than it should be,
        // and we don't want to spread it out.
        //
        init (id: UInt32, name: String, metadata: {String: String}) {
            pre {
                !Showdown.seriesIDByName.containsKey(name): "A Series with that name already exists"
            }
            self.id = Showdown.nextSeriesID
            self.name = name
            self.metadata = metadata
            self.active = true   

            Showdown.nextSeriesID = Showdown.nextSeriesID + 1 as UInt32

            emit SeriesCreated(
                id: self.id,
                name: self.name,
                metadata: self.metadata
            )
        }
    }

    // Get the publicly available data for a Series by id
    //
    pub fun getSeriesData(id: UInt32): Showdown.SeriesData {
        pre {
            Showdown.seriesByID[id] != nil: "Cannot borrow series, no such id"
        }

        return Showdown.SeriesData(id: id)
    }

    // Get the publicly available data for a Series by name
    //
    pub fun getSeriesDataByName(name: String): Showdown.SeriesData {
        pre {
            Showdown.seriesIDByName[name] != nil: "Cannot borrow series, no such name"
        }

        let id = Showdown.seriesIDByName[name]!

        return Showdown.SeriesData(id: id)
    }

    // Get all series names (this will be *long*)
    //
    pub fun getAllSeriesNames(): [String] {
        return Showdown.seriesIDByName.keys
    }

    // Get series id for name
    //
    pub fun getSeriesIDByName(name: String): UInt32? {
        return Showdown.seriesIDByName[name]
    }

    //------------------------------------------------------------
    // Set
    //------------------------------------------------------------

    // A public struct to access Set data
    //
    pub struct SetData {
        pub let id: UInt32
        pub let name: String
        pub let metadata: {String: String}

        // initializer
        //
        init (id: UInt32) {
            let set = &Showdown.setByID[id] as! &Showdown.Set
            self.id = id
            self.name = set.name
            self.metadata = set.metadata
        }
    }

    // A top level Set with a unique ID and a name
    //
    pub resource Set {
        pub let id: UInt32
        pub let name: String
        // Contents writable if borrowed!
        // This is deliberate, as it allows admins to update the data.
        pub let metadata: {String: String}

        // initializer
        //
        init (name: String, metadata: {String: String}) {
            self.id = Showdown.nextSetID
            self.name = name
            self.metadata = metadata

            Showdown.nextSetID = Showdown.nextSetID + 1 as UInt32

            emit SetCreated(id: self.id, name: self.name, metadata: self.metadata)
        }
    }

    // Get the publicly available data for a Set
    //
    pub fun getSetData(id: UInt32): Showdown.SetData {
        pre {
            Showdown.setByID[id] != nil: "Cannot borrow set, no such id"
        }

        return SetData(id: id)
    }


    //------------------------------------------------------------
    // Play
    //------------------------------------------------------------

    // A public struct to access Play data
    //
    pub struct PlayData {
        pub let id: UInt32
        pub let classification: String
        pub let metadata: {String: String}

        // initializer
        //
        init (id: UInt32) {
            let play = &Showdown.playByID[id] as! &Showdown.Play
            self.id = id
            self.classification = play.classification
            self.metadata = play.metadata
        }
    }

    // A top level Play with a unique ID and a classification
    //
    pub resource Play {
        pub let id: UInt32
        pub let classification: String
        // Contents writable if borrowed!
        // This is deliberate, as it allows admins to update the data.
        pub let metadata: {String: String}

        // initializer
        //
        init (classification: String, metadata: {String: String}) {
            self.id = Showdown.nextSetID
            self.classification = classification
            self.metadata = metadata

            Showdown.nextPlayID = Showdown.nextPlayID + 1 as UInt32

            emit PlayCreated(id: self.id, classification: self.classification, metadata: self.metadata)
        }
    }

    // Get the publicly available data for a Play
    //
    pub fun getPlayData(id: UInt32): Showdown.PlayData {
        pre {
            Showdown.playByID[id] != nil: "Cannot borrow play, no such id"
        }

        return PlayData(id: id)
    }



    //------------------------------------------------------------
    // Edition
    //------------------------------------------------------------

    // A public struct to access Edition data
    //
    pub struct EditionData {
        pub let id: UInt32
        pub let seriesID: UInt32
        pub let setID: UInt32
        pub let playID: UInt32
        // null means there is no max size, minting is unlimited
        pub let maxMintSize: UInt32?
        pub let numMinted: UInt32
        pub let tier: String
        pub let metadata: {String: String}

       // member function to check if max edition size has been reached
       pub fun maxEditionMintSizeReached(): Bool {
            if self.numMinted == self.maxMintSize {
                return true
            }
            return false
        }

        // initializer
        //
        init (id: UInt32) {
            let edition = &Showdown.editionByID[id] as! &Showdown.Edition
            self.id = id
            self.seriesID = edition.seriesID
            self.playID = edition.playID
            self.setID = edition.setID
            self.maxMintSize = edition.maxMintSize
            self.numMinted = edition.numMinted
            self.tier = edition.tier
            self.metadata = edition.metadata
        }
    }

    // A top level Edition that contains a Series, Set, and Play
    //
    pub resource Edition {
        pub let id: UInt32
        pub let seriesID: UInt32
        pub let setID: UInt32
        pub let playID: UInt32
        pub var maxMintSize: UInt32?
        pub var numMinted: UInt32
        pub var tier: String
        // Contents writable if borrowed!
        // This is deliberate, as it allows admins to update the data.
        pub let metadata: {String: String}

        // Retire this edition so that no more Moment NFTs can be minted in it
        //
        access(contract) fun retire() {
            pre {
                self.numMinted == self.maxMintSize: "max number of minted moments has already been reached"
            }

            self.maxMintSize = self.numMinted

            emit EditionClosed(id: self.id)
        }

        // Mint a Moment NFT in this edition, with the given minting mintingDate.
        // Note that this will panic if the max mint size has already been reached.
        //
        pub fun mint(): @Showdown.NFT {
            pre {
                self.numMinted == self.maxMintSize: "max number of minted moments has been reached"
            }

            // Create the Moment NFT, filled out with our information
            let momentNFT <- create NFT(
                id: Showdown.totalSupply,
                editionID: self.id,
                serialNumber: self.numMinted
            )
            Showdown.totalSupply = Showdown.totalSupply + 1
            // Keep a running total (you'll notice we used this as the serial number)
            self.numMinted = self.numMinted + 1 as UInt32

            return <- momentNFT
        }

        // initializer
        //
        init (
            seriesID: UInt32,
            setID: UInt32,
            playID: UInt32,
            maxSize: UInt32?,
            tier: String,
            metadata: {String: String}
        ) {
            pre {
                Showdown.seriesByID.containsKey(seriesID): "seriesID does not exist"
                Showdown.setByID.containsKey(setID): "setID does not exist"
                Showdown.playByID.containsKey(playID): "playID does not exist"
                SetData(id: setID).setPlayExistsInEdition() != true: "set play combination already exists in an edition"
            }

            self.id = Showdown.nextEditionID
            self.seriesID = seriesID
            self.setID = setID
            self.playID = playID

            // If an edition size is not set, it has unlimited minting potential
            if maxSize == 0 {
                self.maxMintSize = nil
            } else {
                self.maxMintSize = maxSize
            }

            self.numMinted = 0 as UInt32
            self.tier = tier
            self.metadata = metadata

            Showdown.nextEditionID = Showdown.nextEditionID + 1 as UInt32

            emit EditionCreated(
                id: self.id,
                seriesID: self.seriesID,
                setID: self.setID,
                playID: self.playID,
                maxMintSize: self.maxMintSize,
                tier: self.tier,
                metadata: self.metadata
            )
        }
    }

    // Get the publicly available data for an Edition
    //
    pub fun getEditionData(id: UInt32): EditionData {
        pre {
            Showdown.editionByID[id] != nil: "Cannot borrow edition, no such id"
        }

        return EditionData(id: id)
    }

    //------------------------------------------------------------
    // NFT
    //------------------------------------------------------------

    // A Moment NFT
    //
    pub resource NFT: NonFungibleToken.INFT {
        pub let id: UInt64
        pub let editionID: UInt32
        pub let serialNumber: UInt32
        pub let mintingDate: UFix64

        // Destructor
        //
        destroy() {
            emit NFTBurned(id: self.id)
        }

        // NFT initializer
        //
        init(
            id: UInt64,
            editionID: UInt32,
            serialNumber: UInt32
        ) {
            pre {
                Showdown.editionByID[editionID] != nil: "no such editionID"
                EditionData(id: editionID).maxEditionMintSizeReached() != true: "max edition size already reached"
            }

            self.id = id
            self.editionID = editionID
            self.serialNumber = serialNumber
            self.mintingDate = getCurrentBlock().timestamp

            emit NFTMinted(id: self.id, editionID: self.editionID, serialNumber: self.serialNumber)
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
        pub fun borrowMomentNFT(id: UInt64): &Showdown.NFT? {
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
            let token <- token as! @Showdown.NFT
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
            return &self.ownedNFTs[id] as &NonFungibleToken.NFT
        }

        // borrowMomentNFT gets a reference to an NFT in the collection
        //
        pub fun borrowMomentNFT(id: UInt64): &Showdown.NFT? {
            if self.ownedNFTs[id] != nil {
                let ref = &self.ownedNFTs[id] as auth &NonFungibleToken.NFT
                return ref as! &Showdown.NFT
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
        pub fun mintNFT(editionID: UInt32): @Showdown.NFT
    }

    // A resource that allows managing metadata and minting NFTs
    //
    pub resource Admin: NFTMinter {
        // Borrow a Series
        //
        pub fun borrowSeries(id: UInt32): &Showdown.Series {
            pre {
                Showdown.seriesByID[id] != nil: "Cannot borrow series, no such id"
            }

            return &Showdown.seriesByID[id] as &Showdown.Series
        }

        // Borrow a Set
        //
        pub fun borrowSet(id: UInt32): &Showdown.Set {
            pre {
                Showdown.setByID[id] != nil: "Cannot borrow Set, no such id"
            }

            return &Showdown.setByID[id] as &Showdown.Set
        }

        // Borrow a Play
        //
        pub fun borrowPlay(id: UInt32): &Showdown.Play {
            pre {
                Showdown.playByID[id] != nil: "Cannot borrow Play, no such id"
            }

            return &Showdown.playByID[id] as &Showdown.Play
        }

        // Borrow an Edition
        //
        pub fun borrowEdition(id: UInt32): &Showdown.Edition {
            pre {
                Showdown.editionByID[id] != nil: "Cannot borrow edition, no such id"
            }

            return &Showdown.editionByID[id] as &Showdown.Edition
        }

        // Mint a single NFT
        // The Edition for the given ID must already exist
        //
        pub fun mintNFT(editionID: UInt32): @Showdown.NFT {
            pre {
                // Make sure the edition we are creating this NFT in exists
                Showdown.editionByID.containsKey(editionID): "No such EditionID"
            }

            return <- self.borrowEdition(id: editionID).mint()
        }
    }

    //------------------------------------------------------------
    // Contract lifecycle
    //------------------------------------------------------------

    // Showdown contract initializer
    //
    init() {
        // Set the named paths
        self.CollectionStoragePath = /storage/ShowdownNFTCollection
        self.CollectionPublicPath = /public/ShowdownNFTCollection
        self.AdminStoragePath = /storage/ShowdownAdmin
        self.MinterPrivatePath = /private/ShowdownMinter

        // Initialize the entity counts
        self.totalSupply = 0
        self.nextSeriesID = 0
        self.nextSetID = 0
        self.nextPlayID = 0
        self.nextEditionID = 0

        // Initialize the metadata lookup dictionaries
        self.seriesByID <- {}
        self.seriesIDByName = {}
        self.setByID <- {}
        self.playByID <- {}
        self.editionByID <- {}

        // Create an Admin resource and save it to storage
        let admin <- create Admin()
        self.account.save(<-admin, to: self.AdminStoragePath)
        // Link capabilites to the admin constrained to the Minter
        // and Metadata interfaces
        self.account.link<&Showdown.Admin{Showdown.NFTMinter}>(
            self.MinterPrivatePath,
            target: self.AdminStoragePath
        )

        // Let the world know we are here
        emit ContractInitialized()
    }
}
 