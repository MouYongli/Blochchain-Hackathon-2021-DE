import { AppLayout, Box } from "@awsui/components-react";
import React from "react";
import HomeHeader from "./Header";

const WelcomePage = () => {
  return (
    <AppLayout
      // navigation={}
      content={
        <Box margin={{ bottom: "l" }} padding="xs">
          <HomeHeader />
        </Box>
      }
      disableContentPaddings={true}
      maxContentWidth={1200}
    />
  );
};

export default WelcomePage;
