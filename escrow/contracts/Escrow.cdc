/*
    Escrow Contract for managing NFTs in a Leaderboard Context.
    Holds NFTs in Escrow account awaiting transfer or burn.

    Authors: Corey Humeston corey.humeston@dapperlabs.com
*/

import NonFungibleToken from "./NonFungibleToken.cdc"

pub contract Escrow {
    // Event emitted when a new leaderboard is created
    pub event LeaderboardCreated(name: String)

    // Event emitted when an NFT is deposited to a leaderboard
    pub event NFTDeposited(leaderboardName: String, nftID: UInt64, owner: Address)

    // Named Paths
    pub let AdminStoragePath: StoragePath

    // An interface containing the public functions for adding entries to a leaderboard.
    pub resource interface ILeaderboard {
        pub fun addEntry(nft: @NonFungibleToken.NFT, ownerAddress: Address, leaderboardName: String, typeName: String)
    }

    // The resource representing a leaderboard.
    pub resource Leaderboard: ILeaderboard {
        pub var entries: @{String: {String: LeaderboardEntry}}
        pub let name: String

        // Adds an NFT entry to the leaderboard.
        pub fun addEntry(nft: @NonFungibleToken.NFT, ownerAddress: Address, leaderboardName: String, typeName: String) {
            let nftID = nft.id

            let entry <- create LeaderboardEntry(typeName: typeName, nftID: nftID, ownerAddress: ownerAddress, nft: <-nft)

            if self.entries[self.name] == nil {
                self.entries[self.name] <-! {}
            }
            let leaderboardDict <- self.entries.remove(key: self.name) ?? panic("Expected leaderboard dictionary")

            let typeAndId = typeName.concat("-").concat(nftID.toString())
            if leaderboardDict[typeAndId] != nil {
                panic("Entry already exists for this NFT ID in the leaderboard")
            }

            leaderboardDict[typeAndId] <-! entry

            self.entries[self.name] <-! leaderboardDict

            emit NFTDeposited(leaderboardName: leaderboardName, nftID: nftID, owner: ownerAddress)
        }

        // Destructor for Leaderboard resource.
        destroy() {
            destroy self.entries
        }

        init(name: String) {
            self.name = name
            self.entries <- {}
            self.entries[self.name] <-! {}
        }
    }

    // The resource representing an NFT entry in a leaderboard.
    pub resource LeaderboardEntry {
        pub let typeName: String
        pub let nftID: UInt64
        pub let ownerAddress: Address
        pub let nft: @NonFungibleToken.NFT

        // Destroys the NFT.
        destroy() {
            destroy self.nft
        }

        init(typeName: String, nftID: UInt64, ownerAddress: Address, nft: @NonFungibleToken.NFT) {
            self.typeName = typeName
            self.nftID = nftID
            self.ownerAddress = ownerAddress
            self.nft <- nft
        }
    }

    // An interface containing the Admin function that gets leaderboards by name.
    pub resource interface IAdmin {
        pub fun getLeaderboard(name: String): &Leaderboard?
    }

    // The resource representing an admin.
    pub resource Admin: IAdmin {
        // A dictionary holding leaderboards.
        pub var leaderboards: @{String: Leaderboard}

        // Creates a new leaderboard and stores it.
        pub fun createLeaderboard(name: String) {
            // Create a new Leaderboard resource.
            let newLeaderboard <- create Leaderboard(name: name)

            // Store the leaderboard for future access.
            self.leaderboards[name] <-! newLeaderboard

            // Emit the event.
            emit LeaderboardCreated(name: name)
        }

        // Returns a reference to the leaderboard with the given name.
        pub fun getLeaderboard(name: String): &Leaderboard? {
            return &self.leaderboards[name] as &Leaderboard?
        }

        // Destructor for Admin resource.
        destroy() {
            destroy self.leaderboards
        }

        init() {
            self.leaderboards <- {}
        }
    }

    init() {
        self.AdminStoragePath = /storage/EscrowAdmin
        let admin <- create Admin()
        self.account.save(<-admin, to: self.AdminStoragePath)

        self.account.link<&Admin{IAdmin}>(/public/AdminPublic, target: self.AdminStoragePath)
    }

    // Provides access to the Admin resource.
    access(self) fun getAdmin(): &Admin {
        return self.account.borrow<&Admin>(from: self.AdminStoragePath)
        ?? panic("Could not borrow reference to the Admin resource")
    }
}
