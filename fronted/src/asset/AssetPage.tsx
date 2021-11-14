import { AppLayout, Box } from "@awsui/components-react";
import AssetHeader from "./AssetHeader";

const AssetPage = () => {
  return (
    <AppLayout
      // navigation={}
      navigationHide={true}
      content={
        <>
          <AssetHeader />
          {/* <AssetBody /> */}
        </>
      }
      disableContentPaddings={true}
      maxContentWidth={1380}
    />
  );
};

export default AssetPage;
