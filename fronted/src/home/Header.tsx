import React from "react";
import { Box, Grid } from "@awsui/components-react";
import style from "./Header.module.scss";

const HomeHeader = () => {
  return (
    <Grid
      gridDefinition={[{ colspan: { xxs: 12 } }]}
      className={style["custom-home__header"]}
    >
      <Box padding={{ vertical: "xxl" }}>
        <Grid
          gridDefinition={[
            { offset: { l: 2, xxs: 1 }, colspan: { l: 8, xxs: 10 } },
            {
              colspan: { xl: 6, l: 5, s: 6, xxs: 10 },
              offset: { l: 2, xxs: 1 },
            },
            {
              colspan: { xl: 2, l: 3, s: 4, xxs: 10 },
              offset: { s: 0, xxs: 1 },
            },
          ]}
        >
          <div className={style["custom-home__header-title"]}>
            <Box
              variant="h1"
              fontWeight="bold"
              padding="n"
              fontSize="display-l"
              color="inherit"
            >
              AI MarketPlace
            </Box>
            {/* <Box
              fontWeight="light"
              padding={{ bottom: "s" }}
              fontSize="body-m"
              color="inherit"
            >
              This is a discentralized
            </Box> */}
            <Box variant="p" fontWeight="light">
              <span className={style["custom-home__header-sub-title"]}>
                hhhhhhhhhhh
              </span>
            </Box>
          </div>
          <div></div>
        </Grid>
      </Box>
    </Grid>
  );
};

export default HomeHeader;
