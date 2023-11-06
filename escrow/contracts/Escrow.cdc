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
    pub let AdminPublicPath: PublicPath

    // An interface containing the public functions for adding entries to a leaderboard.
    pub resource interface ILeaderboard {
        pub fun addEntry(nft: @NonFungibleToken.NFT, ownerAddress: Address, leaderboardName: String)
    }

    // The resource representing a leaderboard.
    pub resource Leaderboard: ILeaderboard {
        pub var entries: @{UInt64: LeaderboardEntry}
        pub let name: String
        pub let nftType: Type

        // Adds an NFT entry to the leaderboard.
        pub fun addEntry(nft: @NonFungibleToken.NFT, ownerAddress: Address, leaderboardName: String) {
            let nftID = nft.id

            // Check if the entry already exists
            if self.entries[nftID] != nil {
                panic("Entry already exists for this NFT ID in the leaderboard")
            }

            // Create the entry and add it to the entries map
            let entry <- create LeaderboardEntry(
                nftID: nftID,
                ownerAddress: ownerAddress,
                nft: <-nft
            )

            self.entries[nftID] <-! entry

            emit NFTDeposited(leaderboardName: leaderboardName, nftID: nftID, owner: ownerAddress)
        }

        // Destructor for Leaderboard resource.
        destroy() {
            destroy self.entries
        }

        init(name: String, nftType: Type) {
            self.name = name
            self.nftType = nftType
            self.entries <- {} as @{UInt64: LeaderboardEntry}
        }
    }

    // The resource representing an NFT entry in a leaderboard.
    pub resource LeaderboardEntry {
        pub let nftID: UInt64
        pub let ownerAddress: Address
        pub let nft: @NonFungibleToken.NFT

        // Destroys the NFT.
        destroy() {
            destroy self.nft
        }

        init(nftID: UInt64, ownerAddress: Address, nft: @NonFungibleToken.NFT) {
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
        pub fun createLeaderboard(name: String, nftType: Type) {
            // Create a new Leaderboard resource.
            let newLeaderboard <- create Leaderboard(name: name, nftType: nftType)

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

        self.AdminPublicPath = /public/AdminPublic

        self.account.link<&Admin{IAdmin}>(/public/AdminPublic, target: self.AdminStoragePath)
    }

    // Provides access to the Admin resource.
    access(self) fun getAdmin(): &Admin {
        return self.account.borrow<&Admin>(from: self.AdminStoragePath)
        ?? panic("Could not borrow reference to the Admin resource")
    }
}
