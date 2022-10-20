import NonFungibleToken from "../../contracts/NonFungibleToken.cdc"
import AllDay from "../../contracts/AllDay.cdc"
import MetadataViews from "../../contracts/MetadataViews.cdc"

pub struct NFT {
    pub let name: String
    pub let description: String
    pub let thumbnail: String
    pub let owner: Address
    pub let type: String
    pub let externalURL: String
    pub let storagePath: String
    pub let publicPath: String
    pub let privatePath: String
    pub let collectionName: String
    pub let collectionDescription: String
    pub let collectionSquareImage: String
    pub let collectionBannerImage: String
    pub let royaltyReceiversCount: UInt32
    pub let traitsCount: UInt32
    pub let videoURL: String

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
        self.privatePath = privatePath
        self.collectionName = collectionName
        self.collectionDescription = collectionDescription
        self.collectionSquareImage = collectionSquareImage
        self.collectionBannerImage = collectionBannerImage
        self.royaltyReceiversCount = royaltyReceiversCount
        self.traitsCount = traitsCount
        self.videoURL = videoURL
    }
}

pub fun main(address: Address, id: UInt64): [AnyStruct] {
    let account = getAccount(address)

    let collectionRef = account.getCapability(AllDay.CollectionPublicPath)
            .borrow<&{AllDay.MomentNFTCollectionPublic}>()
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
