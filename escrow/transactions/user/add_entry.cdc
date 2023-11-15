import NonFungibleToken from "../../../contracts/NonFungibleToken.cdc"
import AllDay from "../../../contracts/AllDay.cdc"
import Escrow from "../../../contracts/AllDay.cdc"

transaction(leaderboardName: String, nftID: UInt64) {
    let nft: @NonFungibleToken.NFT
    let receiver: Capability<&{NonFungibleToken.CollectionPublic}>
    let collectionPublic: &Escrow.Collection{Escrow.ICollectionPublic}

    prepare(signer: AuthAccount) {
        // Borrow a reference to the user's NFT collection as a Provider
        let collectionRef = signer.borrow<&{NonFungibleToken.Provider}>(
            from: AllDay.CollectionStoragePath
        ) ?? panic("Could not borrow NFT collection reference")

        // Borrow a reference to the user's NFT collection as a Receiver.
        self.receiver = signer.getCapability<&{NonFungibleToken.CollectionPublic}>(AllDay.CollectionPublicPath)!

        // Withdraw the NFT from the user's collection
        self.nft <- collectionRef.withdraw(withdrawID: nftID)

        // Get the public leaderboard collection
        let escrowAccount = getAccount("../../contracts/AllDay.cdc")
        self.collectionPublic = escrowAccount
            .getCapability<&Escrow.Collection{Escrow.ICollectionPublic}>(Escrow.CollectionPublicPath)
            .borrow()
            ?? panic("Could not borrow a reference to the public leaderboard collection")
    }

    execute {
        // Add the NFT entry to the leaderboard
        self.collectionPublic.addEntryToLeaderboard(
            nft: <-self.nft,
            leaderboardName: leaderboardName,
            depositCap: self.receiver
        )
    }
}
