import { Box, Grid } from "@awsui/components-react";
import Goods from "./Goods";

const Homebody = () => {
  return (
    <Box margin={{ top: "m" }} padding={{ top: "m" }}>
      <Grid
        gridDefinition={[
          {
            colspan: 10,
            offset: { l: 2, xxs: 1 },
          },
        ]}
      >
        <div>
          <Goods />
        </div>
      </Grid>
    </Box>
  );
};

export default Homebody;
