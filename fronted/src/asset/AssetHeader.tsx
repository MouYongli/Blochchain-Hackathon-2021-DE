import { Box, Grid } from "@awsui/components-react";
import style from "./AssetHeader.module.scss";

const AssetHeader = () => {
  return (
    <Grid
      gridDefinition={[
        { colspan: 3 },
        { colspan: 3 },
        { colspan: 3 },
        { colspan: 3 },
      ]}
      className={style["custom-asset__header"]}
    >
      <div>
        <Box
          variant="h3"
          textAlign="center"
          padding="m"
          className={style["custom-asset__header-title"]}
        >
          Home
        </Box>
      </div>
      <div>
        <Box variant="h3" textAlign="center" padding="m" >
          AI model
        </Box>
      </div>
      <div>
        <Box variant="h3" textAlign="center" padding="m">
          data asset
        </Box>
      </div>
      <div>
        <Box variant="h3" textAlign="center" padding="m">
          computing resources
        </Box>
      </div>
    </Grid>
  );
};

export default AssetHeader;
