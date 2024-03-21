import NonFungibleToken from "NonFungibleToken"
import AllDay from "AllDay"
import MetadataViews from "MetadataViews"

access(all) struct NFT {
    access(all) let name: String
    access(all) let description: String
    access(all) let thumbnail: String
    access(all) let owner: Address
    access(all) let type: String
    access(all) let externalURL: String
    access(all) let storagePath: String
    access(all) let publicPath: String
    access(all) let collectionName: String
    access(all) let collectionDescription: String
    access(all) let collectionSquareImage: String
    access(all) let collectionBannerImage: String
    access(all) let royaltyReceiversCount: UInt32
    access(all) let traitsCount: UInt32
    access(all) let videoURL: String

    init(
            name: String,
            description: String,
            thumbnail: String,
            owner: Address,
            type: String,
            externalURL: String,
            storagePath: String,
            publicPath: String,
            privatePath: String,
            collectionName: String,
            collectionDescription: String,
            collectionSquareImage: String,
            collectionBannerImage: String,
            royaltyReceiversCount: UInt32,
            traitsCount: UInt32,
            videoURL: String
    ) {
        self.name = name
        self.description = description
        self.thumbnail = thumbnail
        self.owner = owner
        self.type = type
        self.externalURL = externalURL
        self.storagePath = storagePath
        self.publicPath = publicPath
        self.collectionName = collectionName
        self.collectionDescription = collectionDescription
        self.collectionSquareImage = collectionSquareImage
        self.collectionBannerImage = collectionBannerImage
        self.royaltyReceiversCount = royaltyReceiversCount
        self.traitsCount = traitsCount
        self.videoURL = videoURL
    }
}

access(all) fun main(address: Address, id: UInt64): [AnyStruct] {
    let account = getAccount(address)

    let collectionRef = getAccount(address).capabilities.borrow<&AllDay.Collection>(AllDay.CollectionPublicPath)
            ?? panic("Could not borrow capability from public collection")

    let nft = collectionRef.borrowMomentNFT(id: id)
            ?? panic("Couldn't borrow momentNFT")

    // Get all core views for this NFT
    let displayView = nft.resolveView(Type<MetadataViews.Display>())! as! MetadataViews.Display
    let editionsView = nft.resolveView(Type<MetadataViews.Editions>())! as! MetadataViews.Editions
    let externalURLView = nft.resolveView(Type<MetadataViews.ExternalURL>())! as! MetadataViews.ExternalURL
    let nftCollectionDataView = nft.resolveView(Type<MetadataViews.NFTCollectionData>())! as! MetadataViews.NFTCollectionData
    let nftCollectionDisplayView = nft.resolveView(Type<MetadataViews.NFTCollectionDisplay>())! as! MetadataViews.NFTCollectionDisplay
    let mediasView = nft.resolveView(Type<MetadataViews.Medias>())! as! MetadataViews.Medias
    let royaltiesView = nft.resolveView(Type<MetadataViews.Royalties>())! as! MetadataViews.Royalties
    let serialView = nft.resolveView(Type<MetadataViews.Serial>())! as! MetadataViews.Serial
    let traitsView = nft.resolveView(Type<MetadataViews.Traits>())! as! MetadataViews.Traits

    return [displayView, editionsView, externalURLView, mediasView, nftCollectionDisplayView, royaltiesView, serialView, traitsView]
}
