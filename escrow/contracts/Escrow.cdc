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
    pub event LeaderboardCreated(name: String)

    // Event emitted when an NFT is deposited to a leaderboard.
    pub event NFTDeposited(leaderboardName: String, nftID: UInt64, owner: Address)

    // Event emitted when an NFT is withdrawn from a leaderboard.
    pub event NFTWithdrawn(leaderboardName: String, nftID: UInt64, owner: Address)

    // Event emitted when an NFT is burned from a leaderboard.
    pub event NFTBurned(leaderboardName: String, nftID: UInt64)

    // Named Paths
    pub let AdminStoragePath: StoragePath
    pub let AdminPublicPath: PublicPath

    // An interface containing the public functions for adding entries to a leaderboard.
    pub resource interface ILeaderboard {
        pub fun addEntry(nft: @NonFungibleToken.NFT, leaderboardName: String, depositCap: Capability<&{NonFungibleToken.CollectionPublic}>)
    }

    // The resource representing a leaderboard.
    pub resource Leaderboard: ILeaderboard {
        pub var entries: @{UInt64: LeaderboardEntry}
        pub let name: String
        pub let nftType: Type

        // Adds an NFT entry to the leaderboard.
        pub fun addEntry(nft: @NonFungibleToken.NFT, leaderboardName: String, depositCap: Capability<&{NonFungibleToken.CollectionPublic}>) {
            pre {
                 nft.isInstance(self.nftType): "NFT cannot be used for this leaderboard!"
            }

            let nftID = nft.id

            // Check if the entry already exists
            if self.entries[nftID] != nil {
                panic("Entry already exists for this NFT ID in the leaderboard")
            }

            // Create the entry and add it to the entries map
            let entry <- create LeaderboardEntry(
                nftID: nftID,
                ownerAddress: depositCap.address,
                nft: <-nft,
                depositCap: depositCap
            )

            self.entries[nftID] <-! entry

            emit NFTDeposited(leaderboardName: leaderboardName, nftID: nftID, owner: depositCap.address)
        }

        pub fun getEntriesLength(): Int {
            return self.entries.keys.length
        }

        access(contract) fun withdraw(nftID: UInt64) {
            let entry <- self.entries.remove(key: nftID)!
            entry.withdraw()
            emit NFTWithdrawn(leaderboardName: self.name, nftID: nftID, owner: entry.ownerAddress)
            destroy entry
        }

        access(contract) fun burn(nftID: UInt64) {
            let entry <- self.entries.remove(key: nftID)!
            emit NFTBurned(leaderboardName: self.name, nftID: nftID)
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
        }
    }

    // The resource representing an NFT entry in a leaderboard.
    pub resource LeaderboardEntry {
        pub let nftID: UInt64
        pub let ownerAddress: Address
        pub let nft: @{UInt64: NonFungibleToken.NFT}
        pub let depositCapability: Capability<&{NonFungibleToken.CollectionPublic}>

        pub fun withdraw() {
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
            if self.leaderboards[name] != nil {
                panic("Leaderboard already exists with this name")
            }

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

        // Calls withdraw.
        pub fun withdraw(leaderboardName: String, nftID: UInt64) {
            let leaderboard <- self.leaderboards.remove(key: leaderboardName)!
            leaderboard.withdraw(nftID: nftID)
            self.leaderboards[leaderboardName] <-! leaderboard
        }

        // Calls burn.
        pub fun burn(leaderboardName: String, nftID: UInt64) {
            let leaderboard <- self.leaderboards.remove(key: leaderboardName)!
            leaderboard.burn(nftID: nftID)
            self.leaderboards[leaderboardName] <-! leaderboard
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

        self.account.link<&Admin{IAdmin}>(self.AdminPublicPath, target: self.AdminStoragePath)
    }
}
