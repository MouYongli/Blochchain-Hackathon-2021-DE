import React from "react";
import { Box, Button, Container, Grid } from "@awsui/components-react";
import style from "./Goods.module.scss";
const Goods = () => {
  const goods_list = [
    {
      title: "AI model 1",
      body: "description",
      url: "/asset",
      Image: "../images/header.png",
    },
    {
      title: "AI model 1",
      body: "description",
      url: "/path/to/the/product",
      Image: "../images/header.png",
    },
    {
      title: "AI model 1",
      body: "description",
      url: "/src/asset/AssetPage.tsx",
      Image: "../images/header.png",
    },
    {
      title: "AI model 1",
      body: "description",
      url: "/path/to/the/product",
      Image: "../images/header.png",
    },
  ];
  return (
    <>
      <Box variant="h1" margin={{ top: "xxl" }}>
        Available products
      </Box>
      <Grid
        gridDefinition={[
          { colspan: { xs: 3, xxs: 6 } },
          { colspan: { xs: 3, xxs: 6 } },
          { colspan: { xs: 3, xxs: 6 } },
          { colspan: { xs: 3, xxs: 6 } },
        ]}
      >
        {goods_list.map((good, index) => (
          <Container key={index} className={style["quick-box"]}>
            <Box variant="h3" padding={{ top: "n" }} fontWeight="bold">
              {good.title}
            </Box>
            <img src={good.Image} alt="BigCo Inc. logo" />
            <Box variant="p">{good.body}</Box>
            <Button href={good.url} variant="normal">
              learn more
            </Button>
          </Container>
        ))}
      </Grid>
    </>
  );
};

export default Goods;
