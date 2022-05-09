import { PageDocument } from "@/interface";
import { FC } from "react";
import styles from "./styles.module.css";

import FieldCard from "@/components/FieldCard";

type FullPreviewProps = {
  document: PageDocument;
};

const FullPreview: FC<FullPreviewProps> = ({ document }) => {
  return (
    <div className={styles.container}>
      <div className={styles.report_container}>
        <FieldCard
          available={!!document.title}
          field={"title"}
          content={document.title}
        />
        <FieldCard
          available={!!document.title}
          field={"title"}
          content={document.title}
        />
        <FieldCard
          available={!!document.title}
          field={"title"}
          content={document.title}
        />
        <FieldCard
          available={!!document.title}
          field={"title"}
          content={document.title}
        />
        <FieldCard
          available={!!document.title}
          field={"title"}
          content={document.title}
        />
        <FieldCard
          available={!!document.title}
          field={"title"}
          content={document.title}
        />
        <FieldCard
          available={!!document.title}
          field={"title"}
          content={document.title}
        />
      </div>
    </div>
  );
};

export default FullPreview;
