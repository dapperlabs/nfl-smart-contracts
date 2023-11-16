/*
    Escrow Contract for managing NFTs in a Leaderboard Context.
    Holds NFTs in Escrow account awaiting transfer or burn.

    Authors:
        Corey Humeston: corey.humeston@dapperlabs.com
        Deewai Abdullahi: innocent.abdullahi@dapperlabs.com
*/

import NonFungibleToken from "./NonFungibleToken.cdc"

pub contract Escrow {
    // Event emitted when a new leaderboard is created.
    pub event LeaderboardCreated(name: String, nftType: Type)

    // Event emitted when an NFT is deposited to a leaderboard.
    pub event EntryDeposited(leaderboardName: String, nftID: UInt64, owner: Address)

    // Event emitted when an NFT is returned to the original collection from a leaderboard.
    pub event EntryReturnedToCollection(leaderboardName: String, nftID: UInt64, owner: Address)

    // Event emitted when an NFT is burned from a leaderboard.
    pub event EntryBurned(leaderboardName: String, nftID: UInt64)

    // Named Paths
    pub let CollectionStoragePath: StoragePath
    pub let CollectionPublicPath: PublicPath
    pub let CollectionPrivatePath: PrivatePath

    pub struct LeaderboardInfo {
        pub let name: String
        pub let nftType: Type
        pub let entriesLength: Int

        init(name: String, nftType: Type, entriesLength: Int) {
            self.name = name
            self.nftType = nftType
            self.entriesLength = entriesLength
        }
    }

    // The resource representing a leaderboard.
    pub resource Leaderboard {
        pub var entries: @{UInt64: LeaderboardEntry}
        pub let name: String
        pub let nftType: Type
        pub var entriesLength: Int
        pub var metadata: {String: AnyStruct}

        // Adds an NFT entry to the leaderboard.
        pub fun addEntryToLeaderboard(nft: @NonFungibleToken.NFT, leaderboardName: String, depositCap: Capability<&{NonFungibleToken.CollectionPublic}>) {
            pre {
                 nft.isInstance(self.nftType): "This NFT cannot be used for leaderboard. NFT is not of the correct type."
            }

            let nftID = nft.id

            // Check if the entry already exists
            if self.entries[nftID] != nil {
                panic("Entry already exists for this NFT in the leaderboard")
            }

            // Create the entry and add it to the entries map
            let entry <- create LeaderboardEntry(
                nftID: nftID,
                ownerAddress: depositCap.address,
                nft: <-nft,
                depositCap: depositCap
            )

            self.entries[nftID] <-! entry

            // Increment entries length.
            self.entriesLength = self.entriesLength + 1

            emit EntryDeposited(leaderboardName: leaderboardName, nftID: nftID, owner: depositCap.address)
        }

        // Withdraws an NFT entry from the leaderboard.
        access(contract) fun returnNftToCollection(nftID: UInt64) {
            let entry <- self.entries.remove(key: nftID)!
            entry.returnNftToCollection()
            emit EntryReturnedToCollection(leaderboardName: self.name, nftID: nftID, owner: entry.ownerAddress)

            // Decrement entries length.
            self.entriesLength = self.entriesLength - 1

            destroy entry
        }

        // Burns an NFT entry from the leaderboard.
        access(contract) fun burn(nftID: UInt64) {
            let entry <- self.entries.remove(key: nftID)!
            emit EntryBurned(leaderboardName: self.name, nftID: nftID)

            // Decrement entries length.
            self.entriesLength = self.entriesLength - 1

            destroy entry
        }

        // Destructor for Leaderboard resource.
        destroy() {
            destroy self.entries
        }

        init(name: String, nftType: Type) {
            self.name = name
            self.nftType = nftType
            self.entries <- {}
            self.entriesLength = 0
            self.metadata = {}
        }
    }

    // The resource representing an NFT entry in a leaderboard.
    pub resource LeaderboardEntry {
        pub let nftID: UInt64
        pub let ownerAddress: Address
        pub let nft: @{UInt64: NonFungibleToken.NFT}
        pub let depositCapability: Capability<&{NonFungibleToken.CollectionPublic}>
        pub var metadata: {String: AnyStruct}

        pub fun returnNftToCollection() {
            if self.depositCapability.check() {
                let receiver = self.depositCapability.borrow()
                    as &{NonFungibleToken.CollectionPublic}?
                    ?? panic("Could not borrow the NFT receiver from the capability")

                let nft <- self.nft.remove(key: self.nftID)!
                receiver!.deposit(token: <- nft)
            } else {
                panic("Deposit capability is not valid")
            }
        }

        // Destroys the NFT.
        destroy() {
            destroy self.nft
        }

        init(nftID: UInt64, ownerAddress: Address, nft: @NonFungibleToken.NFT, depositCap: Capability<&{NonFungibleToken.CollectionPublic}>) {
            self.nftID = nftID
            self.ownerAddress = ownerAddress
            self.nft <- {nftID: <-nft}
            self.depositCapability = depositCap
            self.metadata = {}
        }
    }

    // An interface containing the Collection function that gets leaderboards by name.
    pub resource interface ICollectionPublic {
        pub fun getLeaderboardInfo(name: String): LeaderboardInfo?
        pub fun addEntryToLeaderboard(nft: @NonFungibleToken.NFT, leaderboardName: String, depositCap: Capability<&{NonFungibleToken.CollectionPublic}>)
    }

    pub resource interface ICollectionPrivate {
        pub fun createLeaderboard(name: String, nftType: Type)
        pub fun returnNftToCollection(leaderboardName: String, nftID: UInt64)
        pub fun burn(leaderboardName: String, nftID: UInt64)
    }

    // The resource representing a collection.
    pub resource Collection: ICollectionPublic, ICollectionPrivate {
        // A dictionary holding leaderboards.
        access(self) var leaderboards: @{String: Leaderboard}

        // Creates a new leaderboard and stores it.
        pub fun createLeaderboard(name: String, nftType: Type) {
            if self.leaderboards[name] != nil {
                panic("Leaderboard already exists with this name")
            }

            // Create a new leaderboard resource.
            let newLeaderboard <- create Leaderboard(name: name, nftType: nftType)

            // Store the leaderboard for future access.
            self.leaderboards[name] <-! newLeaderboard

            // Emit the event.
            emit LeaderboardCreated(name: name, nftType: nftType)
        }

        // Returns leaderboard info with the given name.
        pub fun getLeaderboardInfo(name: String): LeaderboardInfo? {
            let leaderboard = &self.leaderboards[name] as &Leaderboard?
            if leaderboard == nil {
                return nil
            }

            return LeaderboardInfo(
                name: leaderboard!.name,
                nftType: leaderboard!.nftType,
                entriesLength: leaderboard!.entriesLength
            )
        }

        // Call addEntry.
        pub fun addEntryToLeaderboard(nft: @NonFungibleToken.NFT, leaderboardName: String, depositCap: Capability<&{NonFungibleToken.CollectionPublic}>) {
            let leaderboard = &self.leaderboards[leaderboardName] as &Leaderboard?
            if leaderboard == nil {
                panic("Leaderboard does not exist with this name")
            }

            leaderboard!.addEntryToLeaderboard(nft: <-nft, leaderboardName: leaderboardName, depositCap: depositCap)
        }

        // Calls returnNftToCollection.
        pub fun returnNftToCollection(leaderboardName: String, nftID: UInt64) {
            let leaderboard = &self.leaderboards[leaderboardName] as &Leaderboard?
            if leaderboard == nil {
                panic("Leaderboard does not exist with this name")
            }

            leaderboard!.returnNftToCollection(nftID: nftID)
        }

        // Calls burn.
        pub fun burn(leaderboardName: String, nftID: UInt64) {
            let leaderboard = &self.leaderboards[leaderboardName] as &Leaderboard?
            if leaderboard == nil {
                panic("Leaderboard does not exist with this name")
            }

            leaderboard!.burn(nftID: nftID)
        }

        // Destructor for Collection resource.
        destroy() {
            destroy self.leaderboards
        }

        init() {
            self.leaderboards <- {}
        }
    }

    init() {
        self.CollectionStoragePath = /storage/EscrowLeaderboardCollection
        self.CollectionPrivatePath = /private/EscrowLeaderboardCollectionAccess
        self.CollectionPublicPath = /public/EscrowLeaderboardCollectionInfo

        let collection <- create Collection()
        self.account.save(<-collection, to: self.CollectionStoragePath)
        self.account.link<&Collection{ICollectionPrivate}>(self.CollectionPrivatePath, target: self.CollectionStoragePath)
        self.account.link<&Collection{ICollectionPublic}>(self.CollectionPublicPath, target: self.CollectionStoragePath)
    }
}
