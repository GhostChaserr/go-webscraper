export interface ScrapeLinkPayload {
  link: string;
}

export interface XLSXFWorksheetRow {
  title: string
  description: string
  visitedUrl: string
  heading: string
}

export interface MetaTags {
  description?: string;
}

export interface OgTags {
  ogType?: string;
  ogTitle?: string;
  ogDescription?: string;
  ogSiteName?: string;
  ogUrl?: string;
  ogImage?: string;
}

export interface TwitterTags {
  title?: string;
  description?: string;
  card?: string;
  creator?: string;
  domain?: string;
  image?: string;
}

export interface Headers {
  [key: string]: string;
}

export interface URL {
  Scheme?: string;
  Opaque?: string;
  User?: string;
  Host: string;
  Path: string;
  RawPath?: string;
  ForceQuery: boolean;
  RawQuery?: string;
  Fragment?: string;
  RawFragment: string;
}

export interface PageDocument {
  link: string
  totalLinksCount: number;
  heading?: string;
  language?: string;
  wordsCount?: { [key: string]: number };
  title?: string;
  links?: string[];
  description?: string;
  url?: URL;
  headers?: Headers;
  texts?: string[];
  metaTags?: MetaTags;
  ogTags?: OgTags;
  twitterTags?: TwitterTags;
  linkTexts?: string[];
  canonical?: string;
  keywords?: string[];
  images?: string[];
  numberOfInternalScripts: number;
  numberOfExternalScripts: number;
  numberOfOutboundLinks: number;
  numberOfInboundLinks: number;
  numberOfExternalCss: number;
  numberOfInternalCss: number;
  favIcon?: string;
  manifest?: string;
  hasViewportHtmlTag: boolean;
  hasGoogleSiteVerificationHtmlTag: boolean;
}
