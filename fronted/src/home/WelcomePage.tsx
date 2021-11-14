import { AppLayout, Box } from "@awsui/components-react";
import HomeHeader from "./Header";
import Homebody from "./Homebody";

const WelcomePage = () => {
  return (
    <AppLayout
      // navigation={}
      navigationHide={true}
      content={
        <>
          <HomeHeader />
          <Homebody />
        </>
      }
      disableContentPaddings={true}
      maxContentWidth={1380}
    />
  );
};

export default WelcomePage;
