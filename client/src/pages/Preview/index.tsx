import { FC } from "react";
import styles from "./styles.module.css";
import { useExtractQueryStringParams } from "@/hooks";
import { useQuery } from "react-query";
import config from "@/config";
import axios from "axios";
import FullPreview from "@/components/FullPreview";
import { PageDocument } from "@/interface";

const Preview: FC = () => {
  const queryStringParams = useExtractQueryStringParams();
  const link = queryStringParams.get("link") || "";

  const { isLoading, isError, data } = useQuery(`loading-link-${link}`, () => {
    if (!link) {
      console.log('not running')
      return
    }
    console.log('running')
    return axios.get(config.API_URL + `/scrape-link?link=${link}`)
  })

  if (!link) return <div>Missing Query string</div>;
  if (isLoading) return <div>Loading...</div>;
  if (isError) return <div>Error...</div>;

  return (
    <div className={styles.container}>
      <FullPreview document={data?.data as PageDocument} />
    </div>
  );
};

export default Preview;
