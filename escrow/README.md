# Escrow Smart Contract
The Escrow Smart Contract is designed for managing Non-Fungible Tokens (NFTs) within the context of a leaderboard system. It provides functionality to hold NFTs in escrow until certain conditions are met, either awaiting transfer to the rightful owner based on leaderboard standings or for burning (permanent removal) of the NFTs from circulation.

## NFL Contract Addresses
| Network   | Address     |              |
| ----------|:-----------:| -------------|
| Testnet   |  4dfd62c88d1b6462   | [Flow View Source](https://flow-view-source.com/mainnet/account/0x4dfd62c88d1b6462) |
| Mainnet   |  e4cf4bdc1751c65d   | [Flow View Source](https://flow-view-source.com/testnet/account/0xe4cf4bdc1751c65d) |

# Transactions
### AddEntry
Add an NFT to a specific leaderboard.
- **Parameters:**
    - `leaderboardName`: The name of the leaderboard to add the NFT to.
    - `nftId`: The unique identifier of the NFT to add.
- **Emits:**
    - `EntryDeposited` event on successful addition.

### WithdrawEntry
Withdraw an NFT from a specific leaderboard.
- **Parameters:**
    - `leaderboardName`: The name of the leaderboard to withdraw the NFT from.
    - `nftId`: The unique identifier of the NFT to withdraw. 
- **Emits:**
    - `EntryWithdrawn` event on successful withdrawal.

### BurnEntry
Burn an NFT from a specific leaderboard.
- **Parameters:**
    - `leaderboardName`: The name of the leaderboard to burn the NFT from.
    - `nftId`: The unique identifier of the NFT to burn.
- **Emits:**
    - `EntryBurned` event on successful burn.

### CreateLeaderboard
Create a new leaderboard for NFTs.
- **Parameters:**
    - `leaderboardName`: The name for the new leaderboard.
- **Emits:**
    - `LeaderboardCreated` event on successful creation.

# Entities

### LeaderboardInfo
A structure that holds information about a leaderboard, including its name, the type of NFT it accepts, and the number of entries it contains.

### Leaderboard
A resource that represents a leaderboard, allowing for the addition of NFT entries, withdrawal, and burning of NFTs. It maintains a mapping of the NFT entries and emits relevant events for each action taken.

### LeaderboardEntry
A resource that represents an individual NFT entry within a leaderboard. It holds the NFT, the owner's address, and a capability for depositing the NFT back to the owner if needed.

### Collection
A resource that implements ICollectionPublic and ICollectionPrivate, managing the collection of leaderboards. It handles the creation, addition of entries, and provides access to leaderboards and their information.

# Interfaces

### ICollectionPublic
A resource interface that outlines public functions available for interacting with leaderboards, such as getting leaderboard information.

### ICollectionPrivate
A resource interface that provides private functions for leaderboard management, including the creation of leaderboards and the ability to withdraw or burn NFTs.

# Contract Events
The contract emits events to signal actions taken within the system, providing transparency and a way for users to track changes. These events include:

### LeaderboardCreated: 
Emitted when a new leaderboard is created, including the leaderboard name and the type of NFT it accepts.
### EntryDeposited: 
Emitted when an NFT is deposited into a leaderboard, indicating the leaderboard's name, the NFT's ID, and the owner's address.
### EntryWithdrawn: 
Emitted when an NFT is withdrawn from a leaderboard, with details of the leaderboard name, NFT ID, and owner's address.
### EntryBurned: 
Emitted when an NFT associated with a leaderboard is burned.

## NFT Metadata Standard
The contract conforms to the Flow NFT Metadata standard and implements the Core NFT Views. See
[Flow NFT Catalog](https://www.flow-nft-catalog.com/) for details.
