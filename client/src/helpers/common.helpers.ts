import { PageDocument, XLSXFWorksheetRow } from "@/interface";


export const constructXlsxWorksheet = (pageDocument: PageDocument): XLSXFWorksheetRow => {
  // TODO. add mroe fields
  const row: XLSXFWorksheetRow = {
    title:  pageDocument.title || 'Missing',
    description:  pageDocument.description || 'Missing',
    visitedUrl: pageDocument.link,
    heading: pageDocument.heading || 'Missing',
  }
  return row
}