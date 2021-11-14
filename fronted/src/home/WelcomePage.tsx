import { AppLayout, Box } from "@awsui/components-react";
import HomeHeader from "./Header";
import Homebody from "./Homebody";

const WelcomePage = () => {
  return (
    <AppLayout
      // navigation={}
      navigationHide={true}
      content={
        <Box margin={{ bottom: "l" }} padding="xs">
          <HomeHeader />
          <Homebody />
        </Box>
      }
      disableContentPaddings={true}
      maxContentWidth={1380}
    />
  );
};

export default WelcomePage;
