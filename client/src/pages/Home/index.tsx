import { FC, useState } from "react";
import { Button, Heading, Input, Skeleton, Stack } from "@chakra-ui/react";
import styles from "./styles.module.css";
import config from "@/config";
import { PageDocument, ScrapeLinkPayload } from "@/interface";
import axios from "axios";
import { useMutation } from "react-query";
import ShortPreview from "@/components/ShortPreview";
import { useLocalStorageState } from "ahooks";

const Page: FC = () => {
  const [cachedPDocuments, setCachedPDocument] = useLocalStorageState('use-local-storage-state-demo1', {
    defaultValue: '[]',
  })
  const [pageDocument, setPageDocument] = useState<PageDocument>();
  const [link, setLink] = useState<string>();
  const mutation = useMutation((payload: ScrapeLinkPayload) => {
    return axios.post(config.API_URL + "/scrape-link", payload);
  });

  const onScrapeLink = async () => {
    try {
      if (!link) return;
      const data = await mutation.mutateAsync({ link });
      setPageDocument(data.data);
      setCachedPDocument(JSON.stringify([...JSON.parse(cachedPDocuments), data.data]))
    } catch (error) {
      console.log(error);
    }
  };
  return (
    <div className={styles.container}>
      <div className={styles.content_container}>
        <Heading className={styles.header} as="h4" size="xl">
          Enter link to collect data:
        </Heading>
        <div className={styles.input_container}>
          <Input onChange={(e) => setLink(e.target.value)} />
          <Button
            isLoading={mutation.isLoading}
            onClick={onScrapeLink}
            colorScheme="blue"
          >
            Collect data
          </Button>
        </div>
        {mutation.isSuccess && (
          <div className={styles.preview_container}>
            <ShortPreview document={pageDocument} />
          </div>
        )}
        {mutation.isLoading && (
          <div className={styles.preview_container}>
            <Stack>
              <Skeleton height="400px" />
            </Stack>
          </div>
        )}
      </div>
    </div>
  );
};

export default Page;
