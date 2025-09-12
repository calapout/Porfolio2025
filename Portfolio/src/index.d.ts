export type PrizeModel = {
    Name: string,
    locale: string
}

export type CompanyModel = {
    Name: string,
    locale: string,
    Website: string,
}

export type BitmapModel = {
    ext: string,
    url: string,
    hash: string,
    mime: string,
    name: string,
    size: number,
    width: number,
    height: number,
    sizeInBytes: number
}

export type ImageModel = {
    name: string,
    alternativeText: string,
    caption: string,
    width: number,
    height: number,
    formats: {
        large: BitmapModel,
        small: BitmapModel,
        medium: BitmapModel,
        thumbnail: BitmapModel,
    },
    hash: string,
    ext: string,
    mime: string,
    size: number,
    url: string,
    previewUrl: string,
    provider: string,
}

export type ProjectModel = {
    Title: string,
    Locale: string,
    Description: string,
    TaskRealized: string,
    Technologies: string[],
    HasPrizes: boolean,
    Slug: string,
    IsMadeInCompany: boolean,
    Thumbnail: ImageModel,
    Medias: ImageModel[],
    Prizes: PrizeModel[],
    Company: CompanyModel,
    IsNewAndTrendy: boolean
    IsFavorite: boolean,
    FavoriteText: string
}