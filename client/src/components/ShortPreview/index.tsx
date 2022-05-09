import { FC, useState } from "react";
import { PageDocument } from "@/interface";
import { useNavigate } from "react-router-dom";
import { Badge, Box, Tooltip, Image } from "@chakra-ui/react";
import { DownloadIcon, ExternalLinkIcon } from "@chakra-ui/icons";
import { constructXlsxWorksheet } from "@/helpers/common.helpers";

type ShortPreviewProps = {
  document: PageDocument | undefined;
};

declare global {
  interface Window {
    XLSX: any;
  }
}

const ShortPreview: FC<ShortPreviewProps> = ({ document }) => {
  const [isAllLinksShown, setIsAllLinksShown] = useState<boolean>(false);
  const navigate = useNavigate();

  const exportWS = () => {

    var myFile = "data.xlsx";
    var myWorkSheet = window.XLSX.utils.json_to_sheet([constructXlsxWorksheet(document as PageDocument)]);
    var myWorkBook = window.XLSX.utils.book_new();
    window.XLSX.utils.book_append_sheet(myWorkBook, myWorkSheet, "data");
    window.XLSX.writeFile(myWorkBook, myFile);
  };

  if (!document) return null;

  const onSeeMoreClick = () => navigate(`/preview?link=${document.link}`);
  const onExpandLinksClick = () => setIsAllLinksShown(true);
  return (
    <Box borderWidth="1px" borderRadius="lg" overflow="hidden">
      <Box
        display="flex"
        flexDirection="row"
        justifyContent="space-between"
        p="3"
      >
        <Box>
          Title:
          <>
            {document.title && <Badge marginLeft="2">{document.title}</Badge>}
            {!document.title && (
              <Badge marginLeft="2" colorScheme="red">
                Missing
              </Badge>
            )}
          </>
        </Box>
        <Box
          mt="1"
          display="flex"
          flexDirection="row"
          justifyContent="flex-start"
          marginLeft="5"
        >
          <Tooltip label="See more">
            <ExternalLinkIcon
              ml="2"
              style={{ cursor: "pointer" }}
              onClick={onSeeMoreClick}
            />
          </Tooltip>
          <Tooltip label="Download as Excel">
            <DownloadIcon
              ml="2"
              style={{ cursor: "pointer" }}
              onClick={exportWS}
            />
          </Tooltip>
        </Box>
      </Box>
      <Box flexDirection="row" p="3">
        Description:
        <>
          {document.description && <p>{document.description}</p>}
          {!document.description && (
            <Badge marginLeft="2" colorScheme="red">
              Missing
            </Badge>
          )}
        </>
      </Box>
      <Box flexDirection="row" p="3">
        Canonical:
        <>
          {document.canonical && (
            <Badge marginLeft="2">{document.canonical}</Badge>
          )}
          {!document.canonical && (
            <Badge marginLeft="2" colorScheme="red">
              Missing
            </Badge>
          )}
        </>
      </Box>
      <Box display="flex" flexDirection="row" justifyContent="start" p="3">
        Icon:
        <>
          {document.favIcon && (
            <Image ml="4" width="10" src={document.favIcon} alt="Dan Abramov" />
          )}
          {!document.favIcon && (
            <Badge marginLeft="2" colorScheme="red">
              Missing
            </Badge>
          )}
        </>
      </Box>
      <Box flexDirection="row" p="3">
        <Badge m="1" borderRadius="full" px="2" colorScheme="teal">
          Inbound Links: {document.numberOfInboundLinks}
        </Badge>
        <Badge m="1" borderRadius="full" px="2" colorScheme="teal">
          Outbound Links: {document.numberOfOutboundLinks}
        </Badge>
        <Badge m="1" borderRadius="full" px="2" colorScheme="teal">
          Internal CSS: {document.numberOfInternalCss}
        </Badge>
        <Badge m="1" borderRadius="full" px="2" colorScheme="teal">
          External CSS: {document.numberOfInternalCss}
        </Badge>
      </Box>
      {!!document.links && (
        <Box flexDirection="row" p="3">
          Links:
          <Box maxW="sm">
            {isAllLinksShown && (
              <>
                {document.links.map((link, index) => (
                  <Badge marginTop="2" key={`${link}-${index}`}>
                    <a href={link}>{link}</a>
                  </Badge>
                ))}
              </>
            )}
            {!isAllLinksShown && (
              <>
                {document.links.slice(0, 10).map((link, index) => (
                  <Badge marginTop="2" key={`${link}-${index}`}>
                    <a href={link}>{link}</a>
                  </Badge>
                ))}
                <Badge
                  cursor="pointer"
                  marginTop="4"
                  colorScheme="purple"
                  onClick={onExpandLinksClick}
                >
                  See more
                </Badge>
              </>
            )}
          </Box>
        </Box>
      )}
    </Box>
  );
};

export default ShortPreview;
